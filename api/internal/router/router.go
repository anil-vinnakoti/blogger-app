package router

import (
	"net/http"

	"github.com/anil-vinnakoti/blogger-app/internal/auth"
	"github.com/anil-vinnakoti/blogger-app/internal/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Home route
	r.GET("/", handleHome)

	// Public routes (no middleware)
	r.POST("/register", users.RegisterHandler(db))
	r.POST("/login", users.LoginHandler(db))

	// Protected routes (session required)
	api := r.Group("/api")
	api.Use(auth.SessionMiddleware(db)) // ✅ only here
	{
		// api.GET("/posts", posts.ListHandler(db))
		// api.POST("/posts", posts.CreateHandler(db))
		// add update, delete later
	}

	return r
}

func handleHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to Blogger App")
}
