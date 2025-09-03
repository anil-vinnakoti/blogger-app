package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/anil-vinnakoti/blogger-app/database"
	"github.com/anil-vinnakoti/blogger-app/internal/config"
	"github.com/anil-vinnakoti/blogger-app/internal/router"
)

func main() {
	// Load .env file first
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, falling back to system environment variables")
	}

	// Load env/config
	cfg := config.Load()

	// Connect DB
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}

	// Setup router with routes & middleware
	r := router.Setup(db)

	// Run server
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
