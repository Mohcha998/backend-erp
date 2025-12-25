package domain

type Permission struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique"`
	Name string
}
