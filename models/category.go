package models

type Category struct {
	ID   int
	Name string `gorm:"size:255"`
}
