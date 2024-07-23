package routes

import (
	"log"
	"net/http"

	"github.com/thiaudiott/personalities-api/controllers"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
