package domain

import "time"

type AuditLog struct {
	ID        uint
	UserID    uint
	Action    string
	Entity    string
	EntityID  uint
	IPAddress string
	CreatedAt time.Time
}
