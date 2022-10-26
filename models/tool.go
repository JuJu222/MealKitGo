package models

type Tool struct {
	ID   int
	Name string `gorm:"size:255"`
}
