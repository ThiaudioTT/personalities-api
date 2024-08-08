package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to the database
// and initializer DB variable
func Connect() {
	if DB != nil {
		return
	}

	// Connect to database
	// todo: use env variables
	const dsn = "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"

	log.Println("Initializing the database connection...")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database")
	}

	log.Println("Connected to database")
}

// Connect to the database
// and return the connection
// func Connect() (*gorm.DB, error) {
// 	if DB != nil {
// 		return DB, nil
// 	}
// 	// Connect to database
// 	const dsn = "user=root password=root dbname=root sslmode=disable"

// 	var err error
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Panic("Failed to connect to database")
// 		return nil, err
// 	}

// 	log.Println("Connected to database")
// 	return DB, nil
// }
