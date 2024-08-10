package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thiaudiott/personalities-api/database"
	"github.com/thiaudiott/personalities-api/models"
)

// GET - returns all personalities
func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {

	var personalities []models.Personality

	database.DB.Find(&personalities)

	if len(personalities) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// todo: limit the number of personalities returned
	json.NewEncoder(w).Encode(personalities)
}

// GET - returns a single personality
func GetPersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	database.DB.First(&personality, id)

	if personality.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(personality)
}
