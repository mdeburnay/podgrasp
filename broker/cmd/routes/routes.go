package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BrokerRouter() *gin.Engine {
	env := os.Getenv("ENV")

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Broker!")
	})

	return r
}
