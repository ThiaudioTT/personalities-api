package main

import (
	"fmt"

	"github.com/thiaudiott/personalities-api/routes"
)

func main() {
	fmt.Println("Initializing the application...")
	routes.HandleRequests()
}
