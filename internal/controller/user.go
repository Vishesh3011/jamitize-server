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

	if err := validate.Struct(userRequest); err != nil {
		logger.Error("Invalid request body", "error", err)
		return nil, errors.NewAppError(err, types.BadRequest, nil, types.Controller)
	}

	userService := service.NewService(ctx, app.Logger(), app.DB()).User()
	user, appError := userService.CreateUser(userRequest)
	if appError != nil {
		logger.Error("Failed to create user", "error", appError)
		return nil, appError
	}

	tok, err := utils.GenerateJWT(user.ID, user.Email, app.Config().JwtConfig().Secret())
	if err != nil || len(tok) == 0 {
		logger.Error("Failed to generate JWT", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, nil, types.Controller)
	}

	refreshToken, err := utils.GenerateJWTRefresh(user.ID.String(), user.Email, app.Config().JwtConfig().Secret())
	if err != nil || len(refreshToken) == 0 {
		logger.Error("Failed to generate refresh token", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, nil, types.Controller)
	}

	user, err = userService.RefreshUserToken(user.Email, refreshToken)
	if err != nil {
		logger.Error("Failed to refresh user token", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, nil, types.Controller)
	}

	return user.ToUserResponse(tok), nil
}

func LoginUserController(app application.Application, _ http.ResponseWriter, request *http.Request) (*models.UserResponse, errors.AppError) {
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

	tok, err := utils.GenerateJWT(user.ID, user.Email, app.Config().JwtConfig().Secret())
	if err != nil || len(tok) == 0 {
		logger.Error("Failed to generate JWT", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, nil, types.Controller)
	}
	return user.ToUserResponse(tok), nil
}

func RefreshUserTokenController(app application.Application, _ http.ResponseWriter, request *http.Request) (*models.UserResponse, errors.AppError) {
	ctx := request.Context()
	logger := app.Logger()
	token := request.Header.Get("token")
	email, userId, err := utils.ParseJWT(token, app.Config().JwtConfig().Secret())
	if err != nil {
		logger.Error("Failed to parse JWT", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, nil, types.Controller)
	}

	refreshToken, err := utils.GenerateJWTRefresh(*userId, *email, app.Config().JwtConfig().Secret())
	if err != nil || len(refreshToken) == 0 {
		logger.Error("Failed to generate refresh token", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, nil, types.Controller)
	}

	user, appError := service.NewService(ctx, app.Logger(), app.DB()).User().RefreshUserToken(*email, refreshToken)
	if appError != nil {
		logger.Error("Failed to refresh user token", "error", appError)
		return nil, appError
	}

	tok, err := utils.GenerateJWT(user.ID, user.Email, app.Config().JwtConfig().Secret())
	if err != nil || len(tok) == 0 {
		logger.Error("Failed to generate JWT", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, nil, types.Controller)
	}

	return user.ToUserResponse(tok), nil
}
