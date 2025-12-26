package seeders

import (
	"log"

	"auth-service/internal/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seedSuperAdmin(db *gorm.DB) error {
	// ===============================
	// 1. Pastikan Division ada
	// ===============================
	var division domain.Division
	err := db.FirstOrCreate(
		&division,
		domain.Division{Name: "SYSTEM"},
	).Error
	if err != nil {
		return err
	}

	// ===============================
	// 2. Pastikan Role superadmin ada
	// ===============================
	var role domain.Role
	err = db.FirstOrCreate(
		&role,
		domain.Role{Name: "superadmin"},
	).Error
	if err != nil {
		return err
	}

	// ===============================
	// 3. Cek apakah user sudah ada
	// ===============================
	var existing domain.User
	if err := db.Where("email = ?", "admin@system.local").First(&existing).Error; err == nil {
		log.Println("✅ Superadmin already exists")
		return nil
	}

	// ===============================
	// 4. Buat user
	// ===============================
	password, _ := bcrypt.GenerateFromPassword(
		[]byte("admin123"),
		bcrypt.DefaultCost,
	)

	user := domain.User{
		Name:       "Super Admin",
		Email:      "admin@system.local",
		Password:   string(password),
		DivisionID: division.ID, // ❗ INI YANG PENTING
		IsActive:   true,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	// ===============================
	// 5. Assign role ke user
	// ===============================
	userRole := domain.UserRole{
		UserID: user.ID,
		RoleID: role.ID,
	}

	if err := db.Create(&userRole).Error; err != nil {
		return err
	}

	log.Println("✅ Super Admin seeded successfully")
	return nil
}
