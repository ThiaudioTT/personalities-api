package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/thiaudiott/personalities-api/database"
	"github.com/thiaudiott/personalities-api/routes"
)

func runMigrations() {
	// prob this piece of code should be in a separate package like database

	log.Println("Creating migrate instance...")
	// todo: use env variables
	m, err := migrate.New("file://migrations/", "postgres://root:root@localhost:5433/root?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Applying migrations...")
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("No migrations to apply.")
			return
		}
		log.Fatal(err)
	}
}

func main() {
	log.Println("Initializing the application...")

	// Connect to the database
	database.Connect()

	log.Println("Running migrations...")
	runMigrations()

	log.Println("Starting the server...")
	routes.HandleRequests()
}
