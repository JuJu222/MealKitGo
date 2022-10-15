package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"main/controllers"
	"main/database"
	"main/models"
	"net/http"
	"os"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config :=
		database.Config{
			ServerName: os.Getenv("SERVER_NAME"),
			User:       os.Getenv("USER"),
			Password:   os.Getenv("PASSWORD"),
			DB:         os.Getenv("DB"),
		}

	connectionString := database.GetConnectionString(config)
	err = database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&models.Person{})
}
