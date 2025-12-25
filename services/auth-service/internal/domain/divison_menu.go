package domain

type DivisionMenu struct {
	DivisionID uint `gorm:"primaryKey"`
	MenuID     uint `gorm:"primaryKey"`
}
