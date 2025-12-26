package seeders

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

func seedRolePermissions(db *gorm.DB) error {
	var role domain.Role
	if err := db.Where("name = ?", "superadmin").First(&role).Error; err != nil {
		return err
	}

	var permissions []domain.Permission
	if err := db.Find(&permissions).Error; err != nil {
		return err
	}

	for _, perm := range permissions {
		if err := db.FirstOrCreate(&domain.RolePermission{
			RoleID:       role.ID,
			PermissionID: perm.ID,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}
