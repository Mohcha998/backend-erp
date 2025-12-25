package domain

type Menu struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:50;not null"`

	Divisions []Division `gorm:"many2many:division_menus"`
	Roles     []Role     `gorm:"many2many:role_menus"`
}
