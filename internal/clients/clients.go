package clients

import (
	"context"
	"example/errors"
	"example/internal/core/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Client interface {
	DBClient() *mongo.Client
	Cancel() context.CancelFunc
}

type client struct {
	*mongo.Client
	context.CancelFunc
}

func NewClient(appConfig config.AppConfig) (Client, errors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoC, err := getMongoClient(ctx, appConfig)
	if err != nil {
		return &client{CancelFunc: cancel}, err
	}

	return &client{
		Client:     mongoC,
		CancelFunc: cancel,
	}, nil
}

func (c *client) DBClient() *mongo.Client {
	return c.Client
}

func (c *client) Cancel() context.CancelFunc {
	return c.CancelFunc
}
