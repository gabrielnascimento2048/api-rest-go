package routes

import (
	"log"
	"net/http"

	"github.com/gabrielnascimento2048/api-rest-go/controllers"
	middlaware "github.com/gabrielnascimento2048/api-rest-go/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlaware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.FindbyId).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreateNewPersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeleteOnePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
