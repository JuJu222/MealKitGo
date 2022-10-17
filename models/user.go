package models

type User struct {
	ID        int
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	Phone     string `gorm:"size:255"`
	City      string `gorm:"size:255"`
	Address   string `gorm:"size:255"`
	Orders    []Order
}
