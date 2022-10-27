package models

import "time"

type MealIngredient struct {
	ID           int
	MealID       int
	IngredientID int
	Amount       float32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Meal         Meal       `gorm:"foreignKey:MealID"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID"`
}
