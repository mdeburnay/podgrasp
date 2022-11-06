package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/podgrasp/services"
)

func SetupRouter() *gin.Engine {

	env := os.Getenv("ENV")

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}

	r := gin.Default()

	r.GET("/", services.EllorM8)
	r.GET("/podcast-notes", services.GetPodcastNotes)
	r.POST("/send-email", services.SendEmail)

	return r
}
