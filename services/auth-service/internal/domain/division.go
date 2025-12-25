package domain

import "github.com/google/uuid"

type Division struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string
}
