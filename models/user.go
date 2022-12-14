package models

type User struct {
	ID             int
	FullName       string `gorm:"size:255"`
	Email          string `gorm:"size:255"`
	Password       string `gorm:"size:255"`
	Phone          string `gorm:"size:255"`
	City           string `gorm:"size:255"`
	Ward           string `gorm:"size:255"`
	District       string `gorm:"size:255"`
	Province       string `gorm:"size:255"`
	Address        string `gorm:"size:255"`
	ProfilePicture string `gorm:"size:255"`
	Orders         []Order
}
