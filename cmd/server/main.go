package main

import (
	"log"
	"net/http"

	"github.com/marcos-nsantos/portfolio-api/internal/database"
	"github.com/marcos-nsantos/portfolio-api/internal/httpserver"
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

	s := httpserver.CreateNewServer(db.Client)
	s.MountHandlers()
	if err := http.ListenAndServe(":8080", s.Router); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
