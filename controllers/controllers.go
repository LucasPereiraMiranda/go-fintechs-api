package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-fintechs-api/database"
	"go-fintechs-api/entities"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Start Page")
}

func GetFintechs(w http.ResponseWriter, r *http.Request) {
	var fintechs []entities.Fintech
	database.DB.Find(&fintechs)
	json.NewEncoder(w).Encode(fintechs)
}

func GetOneFintech(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var fintech entities.Fintech

	database.DB.First(&fintech, id)
	json.NewEncoder(w).Encode(fintech)
}

func CreateFintech(w http.ResponseWriter, r *http.Request) {
	var fintech entities.Fintech
	json.NewDecoder(r.Body).Decode(&fintech)
	database.DB.Create(&fintech)
	fmt.Println(r.Body)
	json.NewEncoder(w).Encode(fintech)
}

func DeleteFintech(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var fintech entities.Fintech
	database.DB.Delete(&fintech, id)
	json.NewEncoder(w).Encode(fintech)
}

func EditFintech(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var fintech entities.Fintech
	database.DB.First(&fintech, id)
	json.NewDecoder(r.Body).Decode(&fintech)
	database.DB.Save(&fintech)
	json.NewEncoder(w).Encode(fintech)
}
