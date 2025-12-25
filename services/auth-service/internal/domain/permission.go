package domain

type Permission struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Code        string `gorm:"unique;not null"`
	Description string
}
