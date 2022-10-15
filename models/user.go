package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
	Password  string
	Meals     []Meal `gorm:"many2many:user_meals;"`
}
