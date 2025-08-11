package clients

import (
	"context"
	"example/errors"
	"example/internal/core/config"
	"example/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoHostPrefix = "mongodb://"
)

func getMongoClient(ctx context.Context, appConfig config.AppConfig) (*mongo.Client, errors.AppError) {
	mongoUri := mongoHostPrefix + appConfig.DBConfig().Username() + ":" + appConfig.DBConfig().Password() + "@" + appConfig.DBConfig().Host() + ":" + appConfig.DBConfig().Port() + "/" + appConfig.DBConfig().Database() + "?authSource=" + appConfig.DBConfig().Database()
	clientOptions := options.Client().ApplyURI(mongoUri)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.ToAppError(err, types.InternalServerError, types.Application)
	}
	return mongoClient, nil
}
