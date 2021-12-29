package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gabrielnascimento2048/api-rest-go/database"
	"github.com/gabrielnascimento2048/api-rest-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page with Golang")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func FindbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

	/*for _, personality := range models.Personalidades {
		if strconv.Itoa(personality.Id) == id {
			json.NewEncoder(w).Encode(personality)
		}
	}*/
}

func CreateNewPersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personalidade
	json.NewDecoder(r.Body).Decode(&newPersonality)
	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)

}

func DeleteOnePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
