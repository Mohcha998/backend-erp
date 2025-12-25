package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type AuditLogRepository interface {
	Create(log *domain.AuditLog) error
}

type auditRepo struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditRepo{db}
}

func (r *auditRepo) Create(log *domain.AuditLog) error {
	return r.db.Create(log).Error
}
