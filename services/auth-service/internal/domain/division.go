package domain

type Division struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:50;unique;not null"`

	Menus []Menu `gorm:"many2many:division_menus"`
}
