package domain

import "gorm.io/gorm"

type Division struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Roles     []Role `gorm:"many2many:division_roles"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
