package service

import (
	"context"
	"example/errors"
	"example/internal/models"
	"example/internal/repository"
	"example/internal/types"
	"example/internal/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type UserService interface {
	CreateUser(*models.CreateUserRequest) (*models.User, errors.AppError)
	LoginUser(*models.UserLoginRequest) (*models.User, errors.AppError)
}

type userService struct {
	ctx    context.Context
	logger *slog.Logger
	db     *mongo.Database
}

func NewUserService(ctx context.Context, logger *slog.Logger, db *mongo.Database) UserService {
	return &userService{
		ctx:    ctx,
		logger: logger,
		db:     db,
	}
}

func (u userService) CreateUser(request *models.CreateUserRequest) (*models.User, errors.AppError) {
	user, appError := repository.NewRepository(u.ctx, u.logger, u.db).User().CreateUser(request.ToUser())
	if appError != nil {
		u.logger.Error("Failed to create user", "error", appError)
		return nil, appError
	}

	return user, nil
}

func (u userService) LoginUser(request *models.UserLoginRequest) (*models.User, errors.AppError) {
	user, appError := repository.NewRepository(u.ctx, u.logger, u.db).User().FindUserByEmail(request.Email)
	if appError != nil {
		u.logger.Error("Failed to find user by email", "error", appError)
		return nil, appError
	}

	if match := utils.CheckPasswordHash(request.Password, user.Password); !match {
		u.logger.Error("Password mismatch", "error", match)
		return nil, errors.NewAppError(fmt.Errorf("password does not match."), types.Unauthorized, types.Unauthorized, types.Service)
	}

	return user, nil
}
