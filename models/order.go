package models

import "time"

type Order struct {
	ID        int `json:"id"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
