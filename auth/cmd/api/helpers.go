package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func readJSON(ctx *gin.Context, data any) error {
	maxBytes := 1048576
	if ctx.Request.ContentLength > int64(maxBytes) {
		return errors.New("Request body too large")
	}

	err := ctx.ShouldBindJSON(data)
	if err != nil {
		return err
	}

	return nil
}

func writeJson(ctx *gin.Context) {
	var payload jsonResponse
	// Data to be written as JSON
	data := gin.H{
		"message": payload.Message,
		"data": payload.Data,
	}

	// Write JSON data to response
	ctx.JSON(http.StatusOK, data)
}

func errJson(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	
	if len(status) > 0 {
		statusCode = status[0]
	}

	c.JSON(statusCode, gin.H{"error": true, "message": err.Error()})
}
