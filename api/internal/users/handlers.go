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
		session := auth.Session{
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
