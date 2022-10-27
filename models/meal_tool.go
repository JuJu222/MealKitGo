package models

import "time"

type MealTool struct {
	ID        int
	MealID    int
	ToolID    int
	Amount    float32
	CreatedAt time.Time
	UpdatedAt time.Time
	Meal      Meal `gorm:"foreignKey:MealID"`
	Tool      Tool `gorm:"foreignKey:ToolID"`
}
