package database

import (
	"log"

	"auth-service/internal/domain"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	log.Println("🚀 running database migration")

	err := db.AutoMigrate(
		// master
		&domain.Division{},
		&domain.Menu{},
		&domain.Role{},
		&domain.Permission{},
		&domain.User{},
		&domain.DivisionMenu{},
		&domain.RoleMenu{},
		&domain.RolePermission{},
		&domain.UserRole{},
	)
	if err != nil {
		log.Fatal("❌ migration failed:", err)
	}

	log.Println("✅ migration completed")
}
