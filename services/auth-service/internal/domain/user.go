package domain

type User struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Name       string
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	DivisionID uint   `gorm:"not null"`
}
