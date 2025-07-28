package application

import (
	"context"
	"example/internal/clients"
	"example/internal/core/config"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type Application interface {
	Context() context.Context
	Logger() *slog.Logger
	DB() *mongo.Database
}

type application struct {
	context context.Context
	logger  *slog.Logger
	db      *mongo.Database
}

func NewApplication(appConfig config.AppConfig) (Application, error) {
	ctx := context.Background()

	logger, err := config.NewLogger(appConfig.LoggerConfig().URL())
	if err != nil {
		return nil, err
	}

	client, err := clients.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	db := client.DBClient().Database(appConfig.DBConfig().Database())

	app := application{
		context: ctx,
		logger:  logger,
		db:      db,
	}

	return app, nil
}

func (a application) Context() context.Context {
	return a.context
}

func (a application) Logger() *slog.Logger {
	return a.logger
}

func (a application) DB() *mongo.Database {
	return a.db
}
