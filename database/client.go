package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"main/models"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

func Migrate() {
	Connector.DropTable(
		&models.User{},
		&models.Meal{},
	)
	Connector.AutoMigrate(
		&models.User{},
		&models.Meal{},
	)
	log.Println("Tables migrated")
}
