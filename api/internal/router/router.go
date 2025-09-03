package router

import (
	"github.com/anil-vinnakoti/blogger-app/internal/auth"
	"github.com/anil-vinnakoti/blogger-app/internal/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Global middleware
	r.Use(auth.SessionMiddleware(db))

	// Public routes
	r.POST("/register", users.RegisterHandler(db))
	r.POST("/login", users.LoginHandler(db))

	// // Protected routes
	// api := r.Group("/api")
	// {
	// 	api.GET("/posts", posts.ListHandler(db))
	// 	api.POST("/posts", posts.CreateHandler(db))
	// 	// add update, delete
	// }

	return r
}
