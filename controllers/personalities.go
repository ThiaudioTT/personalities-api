package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thiaudiott/personalities-api/database"
	"github.com/thiaudiott/personalities-api/models"
)

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {

	var personalities []models.Personality

	database.DB.Find(&personalities)

	if len(personalities) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(personalities)
}
