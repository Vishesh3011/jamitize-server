package application

import (
	"example/internal/clients"
	"example/internal/core/config"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type Application interface {
	Logger() *slog.Logger
	DB() *mongo.Database
}

type application struct {
	logger *slog.Logger
	db     *mongo.Database
}

func NewApplication(appConfig config.AppConfig) (Application, error) {
	logger, err := config.NewLogger(appConfig.LoggerConfig().URL())
	if err != nil {
		return nil, err
	}

	client, err := clients.NewClient()
	if err != nil {
		return nil, err
	}
	db := client.DBClient().Database(appConfig.DBConfig().Database())

	app := application{
		logger: logger,
		db:     db,
	}

	return app, nil
}

func (a application) Logger() *slog.Logger {
	return a.logger
}

func (a application) DB() *mongo.Database {
	return a.db
}
