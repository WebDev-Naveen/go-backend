package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"

	controllers "github.com/WebDev-Naveen/go-backend/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User: "postgres",
		Password: "123456",
		Addr: "localhost:5432",
		Database: "gotest",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	// controllers.CreateTable(db)
	controllers.InitiateDB(db)
	return db
}
