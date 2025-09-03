package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Session struct {
	ID        string `gorm:"primaryKey"`
	UserID    uint
	ExpiresAt time.Time
}

// Middleware to check session cookie
func SessionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil || sessionID == "" {
			// block protected routes
			if c.FullPath() != "/login" && c.FullPath() != "/register" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return
			}
			c.Next()
			return
		}

		var session Session
		if err := db.First(&session, "id = ? AND expires_at > ?", sessionID, time.Now()).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired session"})
			return
		}

		c.Set("userID", session.UserID)
		c.Next()
	}
}
