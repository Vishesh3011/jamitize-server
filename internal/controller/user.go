package controller

import (
	"encoding/json"
	"example/errors"
	"example/internal/core/application"
	"example/internal/models"
	"example/internal/service"
	"example/internal/types"
	"net/http"
)

func CreateUserController(app application.Application, _ http.ResponseWriter, request *http.Request) (*models.UserResponse, errors.AppError) {
	ctx := request.Context()
	logger := app.Logger()
	var userRequest *models.CreateUserRequest
	if err := json.NewDecoder(request.Body).Decode(&userRequest); err != nil {
		logger.Error("Invalid request body", "error", err)
		return nil, errors.NewAppError(err, types.BadRequest, nil, types.Controller)
	}
	user, appError := service.NewService(ctx, app.Logger(), app.DB()).User().CreateUser(userRequest)
	if appError != nil {
		logger.Error("Failed to create user", "error", appError)
		return nil, appError
	}
	return user.ToUserResponse(), nil
}
