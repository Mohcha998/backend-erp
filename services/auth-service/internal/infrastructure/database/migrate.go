package database

import (
	"log"

	"auth-service/internal/domain"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	log.Println("🚀 running database migration")

	err := db.AutoMigrate(
		&domain.Division{},
		&domain.Role{},
		&domain.Menu{},
		&domain.Permission{},
		&domain.User{},
		&domain.UserRole{},
	)
	if err != nil {
		log.Fatal("❌ migration failed:", err)
	}

	log.Println("✅ migration completed")
}
