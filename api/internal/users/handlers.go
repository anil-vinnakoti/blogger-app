package users

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/anil-vinnakoti/blogger-app/internal/auth"
	"github.com/anil-vinnakoti/blogger-app/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterHandler - create new user
func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		hashed, err := auth.HashPassword(body.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
			return
		}

		user := models.User{
			Username:     body.Username,
			PasswordHash: hashed,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}

		// Create new session
		session := models.Session{
			ID:        generateSessionID(),
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}
		db.Create(&session)

		// Set session_id cookie
		c.SetCookie("session_id", session.ID, 3600*24, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "user registered"})
	}
}

// LoginHandler - verify user and create session
func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		var user models.User
		if err := db.First(&user, "username = ?", body.Username).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		if !auth.CheckPassword(user.PasswordHash, body.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		// Create new session
		session := models.Session{
			ID:        generateSessionID(),
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}
		db.Create(&session)

		// Set session_id cookie
		c.SetCookie("session_id", session.ID, 3600*24, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "logged in"})
	}
}

func LogoutHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessonId, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no active session"})
			return
		}

		if err := db.Delete(&models.Session{}, "id = ?", sessonId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout"})
			return
		}

		//clear cookie in client
		c.SetCookie("session_id", "", -1, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "logged out"})
	}
}

// MeHandler - get current logged-in user from session
func MeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get session_id cookie
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
			return
		}

		// 2. Find session in DB
		var session models.Session
		if err := db.First(&session, "id = ?", sessionID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			return
		}

		// 3. Check expiration
		if time.Now().After(session.ExpiresAt) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
			return
		}

		// 4. Fetch the user
		var user models.User
		if err := db.First(&user, "id = ?", session.UserID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		// 5. Return user (omit password hash)
		c.JSON(http.StatusOK, gin.H{
			"id":       user.ID,
			"username": user.Username,
		})
	}
}

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users models.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
			return
		}
		c.IndentedJSON(http.StatusOK, users)
	}
}

// Utility to generate random session IDs
func generateSessionID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
