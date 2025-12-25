package domain

import "time"

type ActivityLog struct {
	ID        uint
	UserID    uint
	Activity  string
	IPAddress string
	CreatedAt time.Time
}
