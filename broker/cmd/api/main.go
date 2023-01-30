package main

import (
	"broker/cmd/routes"
	"log"

	"github.com/gin-contrib/cors"
)

const PORT = "80"

type Config struct {}

func main() {
	log.Println("Starting Broker service on port: ", PORT)

	r := routes.BrokerRouter()
	r.Use(cors.Default())
	r.Run(":" + PORT)
	log.Println("Broker service started on port: ", PORT)
}
