package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thiaudiott/personalities-api/controllers"
	"github.com/thiaudiott/personalities-api/middlewares"
)

const PORT = ":8080"

func HandleRequests() {

	log.Println("Creating router...")
	r := mux.NewRouter()

	// defining middleware
	log.Println("Defining middleware...")
	middlewares.ContentTypeJSON(r)

	log.Println("Creating routes...")
	r.HandleFunc("/", controllers.HelloWorld).Methods("GET") // test route
	RegisterPersonalityRoutes(r)

	log.Println("All done. Listening on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
