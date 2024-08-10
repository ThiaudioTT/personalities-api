package routes

import (
	"github.com/gorilla/mux"
	"github.com/thiaudiott/personalities-api/controllers"
)

func RegisterPersonalityRoutes(r *mux.Router) {
	r.HandleFunc("/personalities", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonality).Methods("GET")
}
