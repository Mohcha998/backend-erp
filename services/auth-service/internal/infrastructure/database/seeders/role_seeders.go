package seeders

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

func seedRoles(db *gorm.DB) error {
	roles := []domain.Role{
		{Name: "superadmin"},
		{Name: "admin"},
		{Name: "staff"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, domain.Role{Name: role.Name}).Error; err != nil {
			return err
		}
	}

	return nil
}
