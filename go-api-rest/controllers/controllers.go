package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tupizz/go-api-rest/database"
	"github.com/tupizz/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
}

func GetAllPersonas(w http.ResponseWriter, r *http.Request) {
	var sliceOfPersonas []models.Persona
	database.DB.Find(&sliceOfPersonas)
	json.NewEncoder(w).Encode(sliceOfPersonas)
}

func GetOnePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var persona models.Persona
	database.DB.First(&persona, vars["personaId"])
	json.NewEncoder(w).Encode(persona)
}

func CreatePersona(w http.ResponseWriter, r *http.Request) {
	var newPersona models.Persona
	json.NewDecoder(r.Body).Decode(&newPersona)
	database.DB.Create(&newPersona)
	json.NewEncoder(w).Encode(newPersona)
}

func DeletePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["personaId"]
	var persona models.Persona
	database.DB.Delete(&persona, id)
	json.NewEncoder(w).Encode(persona)
}

func EditPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["personaId"]
	var persona models.Persona
	database.DB.First(&persona, id)
	json.NewDecoder(r.Body).Decode(&persona)
	database.DB.Save(&persona)
	json.NewEncoder(w).Encode(persona)
}
