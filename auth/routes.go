package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Config) routes() http.Handler {
	router := gin.Default()

	router.Use(cors.New((cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods: []string {"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string {"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	})))

	return router;
}
