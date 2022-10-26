package models

import "time"

type Review struct {
	ID        int
	MealID    int
	UserID    int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Meal      Meal `gorm:"foreignKey:MealID"`
	User      User `gorm:"foreignKey:UserID"`
}
