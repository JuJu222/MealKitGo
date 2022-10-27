package models

type Step struct {
	ID          int
	Title       string `gorm:"size:255"`
	Description string
	MealID      uint
}
