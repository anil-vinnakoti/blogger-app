package router

import (
	"net/http"
	"time"

	"github.com/anil-vinnakoti/blogger-app/internal/auth"
	"github.com/anil-vinnakoti/blogger-app/internal/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// ✅ Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // needed for cookies
		MaxAge:           12 * time.Hour,
	}))

	// Home route
	r.GET("/", handleHome)

	api := r.Group("/api")
	// Public routes
	api.POST("/register", users.RegisterHandler(db))
	api.POST("/login", users.LoginHandler(db))

	// Protected routes
	api.Use(auth.SessionMiddleware(db))

	api.GET("/me", users.MeHandler(db))
	api.GET("/users", users.GetUsers(db))

	return r
}

func handleHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to Blogger App")
}
