package models

import "time"

type OrderDetail struct {
	ID        int
	MealID    int
	OrderID   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Meal      Meal  `gorm:"foreignKey:MealID"`
	Order     Order `gorm:"foreignKey:OrderID"`
}
