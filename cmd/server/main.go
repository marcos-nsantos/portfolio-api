package main

import "log"

func run() error {
	log.Println("Starting server...")
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
