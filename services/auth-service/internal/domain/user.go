package domain

import "gorm.io/gorm"

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	IsActive   bool   `gorm:"default:true"`
	DivisionID uint
	Division   Division
	Roles      []Role `gorm:"many2many:user_roles"`
	CreatedAt  int64
	UpdatedAt  int64
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
