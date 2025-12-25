package domain

import "github.com/google/uuid"

type Menu struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string
	Path string
}
