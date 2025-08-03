package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type Repository interface {
	User() UserRepository
}

type repository struct {
	userRepo UserRepository
}

func (r repository) User() UserRepository {
	return r.userRepo
}

func NewRepository(ctx context.Context, logger *slog.Logger, db *mongo.Database) Repository {
	return &repository{
		userRepo: NewUserRepository(ctx, logger, db),
	}
}
