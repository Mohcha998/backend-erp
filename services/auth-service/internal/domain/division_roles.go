package domain

type DivisionRole struct {
	DivisionID uint `gorm:"primaryKey"`
	RoleID     uint `gorm:"primaryKey"`
}
