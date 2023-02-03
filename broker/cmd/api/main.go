package main

import (
	"broker/cmd/routes"
	"log"
)

const PORT = "80"

func main() {
	log.Println("Starting Broker service on port: ", PORT)
	r := routes.BrokerRouter()
	r.Run(":" + PORT)
	log.Println("Broker service started on port: ", PORT)
}
