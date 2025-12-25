package domain

import "github.com/google/uuid"

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string
}
