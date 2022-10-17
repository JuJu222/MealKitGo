package models

import "time"

type MealCategory struct {
	MealID     int `gorm:"primaryKey"`
	CategoryID int `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
