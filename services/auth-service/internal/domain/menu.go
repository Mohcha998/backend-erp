package domain

import "gorm.io/gorm"

type Menu struct {
	ID        uint   `gorm:"primaryKey"`
	Code      string `gorm:"unique;not null"`
	Name      string `gorm:"not null"`
	Path      string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
