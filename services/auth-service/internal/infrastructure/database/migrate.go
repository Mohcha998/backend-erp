package database

import (
	"log"

	"auth-service/internal/domain"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	log.Println("🚀 running database migration")

	err := db.AutoMigrate(
		// ===== MASTER =====
		&domain.Division{},
		&domain.Role{},
		&domain.Menu{},
		&domain.Permission{},
		&domain.User{},

		// ===== RELATION =====
		&domain.UserRole{},
		&domain.DivisionRole{},
		&domain.RoleMenu{},
		&domain.RolePermission{},

		// ===== LOGGING =====
		&domain.ActivityLog{},
		&domain.AuditLog{},
		&domain.RefreshToken{},
	)

	if err != nil {
		log.Fatal("❌ migration failed:", err)
	}

	log.Println("✅ migration completed")
}
