package seeders

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

func seedMenus(db *gorm.DB) error {
	menus := []domain.Menu{
		{Code: "dashboard", Name: "Dashboard", Path: "/dashboard"},
		{Code: "users", Name: "Users", Path: "/users"},
		{Code: "roles", Name: "Roles", Path: "/roles"},
		{Code: "menus", Name: "Menus", Path: "/menus"},
		{Code: "divisions", Name: "Divisions", Path: "/divisions"},
	}

	for _, m := range menus {
		if err := db.FirstOrCreate(&m, domain.Menu{Code: m.Code}).Error; err != nil {
			return err
		}
	}

	return nil
}
