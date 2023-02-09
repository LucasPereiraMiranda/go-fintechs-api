package routes

import (
	"go-fintechs-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/fintechs", controllers.GetFintechs)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
