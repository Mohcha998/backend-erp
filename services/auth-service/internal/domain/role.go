package domain

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:50;unique;not null"`

	Menus       []Menu       `gorm:"many2many:role_menus"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}
