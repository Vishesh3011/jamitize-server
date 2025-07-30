package controller

import (
	"example/errors"
	"example/internal/core/application"
	"net/http"
)

func HealthCheck(application.Application, *http.ResponseWriter, *http.Request) (*string, errors.AppError) {
	status := "I am alive!!"
	return &status, nil
}
