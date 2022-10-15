package models

import "time"

type UserMeal struct {
	UserID    int `gorm:"primaryKey"`
	MealID    int `gorm:"primaryKey"`
	CreatedAt time.Time
}
