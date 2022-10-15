package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main/models"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

func Migrate() {
	err := Connector.Migrator().DropTable(
		&models.User{},
		&models.Meal{},
		&models.UserMeal{},
	)
	if err != nil {
		panic(err.Error())
	}
	err = Connector.SetupJoinTable(&models.User{}, "Meals", &models.UserMeal{})
	if err != nil {
		panic(err.Error())
	}
	err = Connector.AutoMigrate(
		&models.User{},
		&models.Meal{},
		&models.UserMeal{},
	)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Tables Migrated")
}