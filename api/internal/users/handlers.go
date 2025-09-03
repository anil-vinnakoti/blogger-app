package users

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/anil-vinnakoti/blogger-app/internal/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string // ⚠️ should be hashed in real app
}

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

		user := User{
			Username: body.Username,
			Password: body.Password, // ⚠️ hash this using bcrypt in production
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

		var user User
		if err := db.First(&user, "username = ?", body.Username).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		// ⚠️ In real app, compare hashed password with bcrypt
		if user.Password != body.Password {
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

// Utility to generate random session IDs
func generateSessionID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
