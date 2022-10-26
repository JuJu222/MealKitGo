package controllers

import (
	"encoding/json"
	"io/ioutil"
	"main/database"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllMeal(w http.ResponseWriter, r *http.Request) {
	var meals []models.User
	database.Connector.Find(&meals)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meals)
}

func GetMealByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var meal models.User
	database.Connector.First(&meal, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meal)
}

func CreateMeal(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var meal models.User
	json.Unmarshal(requestBody, &meal)

	database.Connector.Create(meal)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(meal)
}

func UpdateMealByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var meal models.User
	json.Unmarshal(requestBody, &meal)
	database.Connector.Save(&meal)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meal)
}

func DeletMealByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var meal models.User
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&meal)
	w.WriteHeader(http.StatusNoContent)
}
