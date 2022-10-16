package models

import "time"

type MealCategory struct {
	UserID    int `gorm:"primaryKey"`
	MealID    int `gorm:"primaryKey"`
	CreatedAt time.Time
}
