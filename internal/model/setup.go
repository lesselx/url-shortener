package data

import (
	"log"

	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&Link{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated")
}
