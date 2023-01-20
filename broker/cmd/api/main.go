package main

import (
	"broker/cmd/routes"
	"log"

	"github.com/gin-contrib/cors"
)

const PORT = "80"

type Config struct {}

func main() {
	log.Println("Starting broker service on port: ", PORT)

	r := routes.BrokerRouter()
	r.Use(cors.Default())
	_ = r.Run(":" + PORT)
	log.Println("Authentication service started on port: ", PORT)
}
