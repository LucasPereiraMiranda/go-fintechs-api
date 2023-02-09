package controllers

import (
	"encoding/json"
	"fmt"
	"go-fintechs-api/entities"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetFintechs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(entities.Fintechs)
}
