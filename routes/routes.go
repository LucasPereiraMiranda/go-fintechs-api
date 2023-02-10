package routes

import (
	"github.com/gorilla/mux"
	"go-fintechs-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/fintechs", controllers.GetFintechs).Methods("Get")
	router.HandleFunc("/api/fintechs/{id}", controllers.GetOneFintech).Methods("Get")
	router.HandleFunc("/api/fintechs", controllers.CreateFintech).Methods("Post")
	router.HandleFunc("/api/fintechs/{id}", controllers.DeleteFintech).Methods("Delete")
	router.HandleFunc("/api/fintechs/{id}", controllers.EditFintech).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", router))
}
