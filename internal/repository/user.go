package repository

import (
	"context"
	"example/errors"
	"example/internal/models"
	"example/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type UserRepository interface {
	CreateUser(*models.User) (*models.User, errors.AppError)
	FindUserByEmail(string) (*models.User, errors.AppError)
}

type userRepository struct {
	ctx    context.Context
	db     *mongo.Database
	logger *slog.Logger
}

func NewUserRepository(ctx context.Context, logger *slog.Logger, db *mongo.Database) UserRepository {
	return &userRepository{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, errors.AppError) {
	r.logger.Info("Creating user in repository")
	collection := r.db.Collection("users")
	_, err := collection.InsertOne(r.ctx, user)
	if err != nil {
		r.logger.Error("Failed to create user", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, types.InternalServerError, types.Repository)
	}
	user, appError := r.FindUserByID(user.ID)
	if appError != nil {
		return nil, appError
	}
	return user, nil
}

func (r *userRepository) FindUserByID(userId primitive.ObjectID) (*models.User, errors.AppError) {
	collection := r.db.Collection("users")
	var user models.User
	err := collection.FindOne(r.ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewAppError(err, types.NotFound, types.NotFound, types.Repository)
		}
		r.logger.Error("Failed to find user by ID", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, types.InternalServerError, types.Repository)
	}
	return &user, nil
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, errors.AppError) {
	collection := r.db.Collection("users")
	var user models.User
	err := collection.FindOne(r.ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewAppError(err, types.NotFound, types.NotFound, types.Repository)
		}
		r.logger.Error("Failed to find user by email", "error", err)
		return nil, errors.NewAppError(err, types.InternalServerError, types.InternalServerError, types.Repository)
	}
	return &user, nil
}
