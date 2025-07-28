package clients

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client interface {
	DBClient() *mongo.Client
}

type client struct {
	*mongo.Client
}

func NewClient(ctx context.Context) (Client, error) {
	mongoClient, err := mongo.Connect(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &client{
		Client: mongoClient,
	}, nil
}

func (c *client) DBClient() *mongo.Client {
	return c.Client
}
