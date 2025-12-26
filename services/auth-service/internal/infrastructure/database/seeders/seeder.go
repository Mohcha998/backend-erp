package seeders

import "gorm.io/gorm"

func SeedAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {

		if err := seedRoles(tx); err != nil {
			return err
		}

		if err := seedPermissions(tx); err != nil {
			return err
		}

		if err := seedMenus(tx); err != nil {
			return err
		}

		if err := seedRolePermissions(tx); err != nil {
			return err
		}

		if err := seedSuperAdmin(tx); err != nil {
			return err
		}

		return nil
	})
}
