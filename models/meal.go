package models

import "time"

type Meal struct {
	ID          int
	Name        string `gorm:"size:255"`
	Description string
	Price       int64
	Steps       []Step
	Reviews     []Review
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
