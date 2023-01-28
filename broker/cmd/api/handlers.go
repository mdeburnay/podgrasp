package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func (app *Config) Broker(ctx *gin.Context) {
	payload := jsonResponse{
		Error: false,
		Message: "Hit the broker",
	}

	_ = app.writeJson(ctx, http.StatusOK, payload)
}
