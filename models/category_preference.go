package models

import "time"

type CategoryPreference struct {
	ID         int
	UserID     int
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User     User     `gorm:"foreignKey:UserID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}
