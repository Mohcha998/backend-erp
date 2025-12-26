package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(token *domain.RefreshToken) error
	FindValid(tokenHash string) (*domain.RefreshToken, error)
	Revoke(tokenHash string) error
}

type refreshTokenRepo struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepo{db}
}

func (r *refreshTokenRepo) Create(token *domain.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *refreshTokenRepo) FindValid(tokenHash string) (*domain.RefreshToken, error) {
	var rt domain.RefreshToken
	err := r.db.
		Where("token_hash = ? AND revoked = false AND expires_at > NOW()", tokenHash).
		First(&rt).Error
	return &rt, err
}

func (r *refreshTokenRepo) Revoke(tokenHash string) error {
	return r.db.
		Model(&domain.RefreshToken{}).
		Where("token_hash = ?", tokenHash).
		Update("revoked", true).Error
}
