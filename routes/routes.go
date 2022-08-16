package routes

import (
	"log"
	"net/http"

	"github.com/danitrod/go-rest-api/controllers"
	"github.com/danitrod/go-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.Json)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
