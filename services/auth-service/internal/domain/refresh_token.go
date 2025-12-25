package domain

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;"`
	TokenHash string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Revoked   bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
