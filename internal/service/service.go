package service

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type Service interface {
	User() UserService
}

type service struct {
	userSvc UserService
}

func (s service) User() UserService {
	return s.userSvc
}

func NewService(ctx context.Context, logger *slog.Logger, db *mongo.Database) Service {
	return &service{
		userSvc: NewUserService(ctx, logger, db),
	}
}
