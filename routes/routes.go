package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thiaudiott/personalities-api/controllers"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", r))
}
