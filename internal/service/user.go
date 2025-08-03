package service

import (
	"context"
	"example/errors"
	"example/internal/models"
	"example/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type UserService interface {
	CreateUser(*models.CreateUserRequest) (*models.User, errors.AppError)
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
	u.logger.Info("Creating user", "request", request)
	user, appError := repository.NewRepository(u.ctx, u.logger, u.db).User().CreateUser(request.ToUser())
	if appError != nil {
		u.logger.Error("Failed to create user", "error", appError)
		return nil, appError
	}

	return user, nil
}
