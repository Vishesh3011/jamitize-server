package controller

import (
	"example/errors"
	"example/internal/core/application"
	"net/http"
)

func HealthCheck(app application.Application, _ http.ResponseWriter, request *http.Request) (*string, errors.AppError) {
	status := "I am alive!!"
	return &status, nil
}
