package domain

type RoleMenuPermission struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	RoleID     uint `gorm:"not null"`
	DivisionID uint `gorm:"not null"`
	MenuID     uint `gorm:"not null"`
	CanAccess  bool `gorm:"default:false"`
}
