package models

type Ingredient struct {
	ID   int
	Name string `gorm:"size:255"`
	Unit string `gorm:"size:255"`
}
