package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func (app *Config) readJSON(ctx *gin.Context , data any) error {
	maxBytes := 1048576 // 1MB

	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, int64(maxBytes))

	decoder := json.NewDecoder(ctx.Request.Body)

	err := decoder.Decode(data)

	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("request body must only contain a single JSON object")
	}

	return nil
}

func (app *Config) writeJson(ctx *gin.Context, status int, data any) error {
	out, err := json.Marshal(data)

	if err != nil {
		return err
	}

	ctx.Writer.Write(out)

	return nil
}

func (app *Config) errJson(ctx *gin.Context, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := jsonResponse{
		Error: true,
		Message: err.Error(),
	}

	return app.writeJson(ctx, statusCode, payload)
}
