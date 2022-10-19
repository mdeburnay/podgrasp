package main

import (
	"github.com/gin-contrib/cors"
	"github.com/podgrasp/routes"
)

// PORT trigger
const PORT string = ":9090"

func main() {
		r := routes.SetupRouter()
		r.Use(cors.Default())
		_ = r.Run(PORT)
	}
