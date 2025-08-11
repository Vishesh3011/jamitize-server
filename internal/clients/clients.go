package clients

import (
	"context"
	"example/errors"
	"example/internal/core/config"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Client interface {
	DBClient() *mongo.Client
}

type client struct {
	*mongo.Client
}

func NewClient(appConfig config.AppConfig) (Client, errors.AppError) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoC, err := getMongoClient(ctx, appConfig)
	if err != nil {
		return nil, err
	}

	return &client{
		Client: mongoC,
	}, nil
}

func (c *client) DBClient() *mongo.Client {
	return c.Client
}
