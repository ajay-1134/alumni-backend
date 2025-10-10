package db

import (
	"log"

	"github.com/ajay-1134/alumni-backend/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnnectDB(dsn string) *gorm.DB {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	database.AutoMigrate(&domain.User{})
	database.AutoMigrate(&domain.Post{})

	return database
}
