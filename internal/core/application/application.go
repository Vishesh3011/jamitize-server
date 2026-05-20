package application

import (
	"example/errors"
	"example/internal/clients"
	"example/internal/core/config"
	"example/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type Application interface {
	Logger() *slog.Logger
	DB() *mongo.Database
	Config() config.AppConfig
}

type application struct {
	logger *slog.Logger
	db     *mongo.Database
	config config.AppConfig
}

func NewApplication(appConfig config.AppConfig) (Application, errors.AppError) {
	logger, err := config.NewLogger(appConfig.LoggerConfig().URL())
	if err != nil {
		return nil, errors.ToAppError(err, types.InternalServerError, types.Application)
	}

	client, err := clients.NewClient(appConfig)
	if err != nil {
		return nil, errors.ToAppError(err, types.InternalServerError, types.Application)
	}
	db := client.DBClient().Database(appConfig.DBConfig().Database())

	app := application{
		logger: logger,
		db:     db,
		config: appConfig,
	}

	return app, nil
}

func (a application) Logger() *slog.Logger {
	return a.logger
}

func (a application) DB() *mongo.Database {
	return a.db
}

func (a application) Config() config.AppConfig {
	return a.config
}
