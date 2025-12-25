package domain

type Permission struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"size:50;unique;not null"`

	Roles []Role `gorm:"many2many:role_permissions"`
}
