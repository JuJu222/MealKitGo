package models

import "time"

type MealCategory struct {
	ID         int
	MealID     int
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Meal       Meal     `gorm:"foreignKey:MealID"`
	Category   Category `gorm:"foreignKey:CategoryID"`
}
