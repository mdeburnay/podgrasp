package main

import (
	"auth/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const port = "80"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service on port: ", port)

	// TODO: Connect to database

	// Set up Config

	app := Config{}

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Panic("Error starting server: ", err)
	}
}
