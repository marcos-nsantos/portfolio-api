package main

import (
	"log"

	"github.com/marcos-nsantos/portfolio-api/internal/database"
)

func run() error {
	log.Println("Starting server...")
	db, err := database.New()
	if err != nil {
		return err
	}
	if err := db.CreateTables(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
