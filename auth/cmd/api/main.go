package main

import (
	"auth/cmd/routes"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
)

const PORT = "80"

var counts int64

func main() {
	log.Println("Starting authentication service on port: ", PORT)

	// TODO: Connect to database
	conn := connectToDB()
	if conn == nil {
		log.Panicln("Could not connect to database")
	}

	r := routes.AuthRouter()
	r.Use(cors.Default())
	_ = r.Run(":" + PORT)
	log.Println("Authentication service started on port: ", PORT)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
		if err != nil {
			return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
		 log.Println("Error connecting to database: ", err)
		 counts++
		} else {
			log.Println("Connected to database after ", counts, " attempts")
			return connection
		}

		if counts > 10 {
			log.Panicln(err)
			return nil
		}
		
		log.Panicln("Backing off...")
		time.Sleep(2 * time.Second)
		continue
	}
}
