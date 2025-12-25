package seeders

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

func seedPermissions(db *gorm.DB) error {
	perms := []domain.Permission{
		{Code: "user.read"},
		{Code: "user.create"},
		{Code: "user.update"},
		{Code: "user.delete"},

		{Code: "role.read"},
		{Code: "role.create"},
		{Code: "role.update"},
		{Code: "role.delete"},

		{Code: "menu.read"},
		{Code: "menu.create"},
		{Code: "menu.update"},
		{Code: "menu.delete"},

		{Code: "division.read"},
		{Code: "division.create"},
		{Code: "division.update"},
		{Code: "division.delete"},
	}

	for _, p := range perms {
		if err := db.FirstOrCreate(&p, domain.Permission{Code: p.Code}).Error; err != nil {
			return err
		}
	}

	return nil
}
