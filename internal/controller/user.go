package controller

import (
	"encoding/json"
	"example/errors"
	"example/internal/core/application"
	"example/internal/models"
	"example/internal/service"
	"example/internal/types"
	"example/internal/utils"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

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

func LoginUserController(app application.Application, _ http.ResponseWriter, request *http.Request) (*models.UserLoginResponse, errors.AppError) {
	ctx := request.Context()
	logger := app.Logger()
	var userLoginReq *models.UserLoginRequest
	if err := json.NewDecoder(request.Body).Decode(&userLoginReq); err != nil {
		logger.Error("Invalid request body", "error", err)
		return nil, errors.NewAppError(err, types.BadRequest, nil, types.Controller)
	}

	if err := validate.Struct(userLoginReq); err != nil {
		return nil, errors.NewAppError(err, types.BadRequest, nil, types.Controller)
	}

	user, appError := service.NewService(ctx, app.Logger(), app.DB()).User().LoginUser(userLoginReq)
	if appError != nil {
		logger.Error("Failed to login user", "error", appError)
		return nil, appError
	}

	tok, err := utils.GenerateJWT(user.ID, app.Config().JwtConfig().Secret())
	if err != nil {
		logger.Error("Failed to generate JWT", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, nil, types.Controller)
	}
	return models.NewUserLoginResponse(user.ID, user.Email, tok), nil
}
