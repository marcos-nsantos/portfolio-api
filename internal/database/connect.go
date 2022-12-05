package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Client *gorm.DB
}

func New() (*Connection, error) {
	dbURL := os.Getenv("DATABASE_URL")
	var counts uint8
	for {
		connection, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			return &Connection{Client: connection}, nil
		}

		if counts > 10 {
			log.Println(err)
			return nil, err
		}

		log.Println("Backing off for five seconds....")
		time.Sleep(5 * time.Second)
		continue
	}
}
