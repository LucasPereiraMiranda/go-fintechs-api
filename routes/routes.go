package routes

import (
	"go-fintechs-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/fintechs", controllers.GetFintechs).Methods("Get")
	router.HandleFunc("/api/fintechs/{id}", controllers.GetOneFintech).Methods("Get")
	router.HandleFunc("/api/fintechs", controllers.CreateFintech).Methods("Post")
	router.HandleFunc("/api/fintechs/{id}", controllers.DeleteFintech).Methods("Delete")
	router.HandleFunc("/api/fintechs/{id}", controllers.EditFintech).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
