package database

import (
	"log"

	"github.com/anil-vinnakoti/blogger-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate schema
	if err := db.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		return nil, err
	}

	log.Println("✅ Connected to PostgreSQL with GORM")
	return db, nil
}
