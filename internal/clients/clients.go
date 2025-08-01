package clients

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Client interface {
	DBClient() *mongo.Client
}

type client struct {
	*mongo.Client
}

func NewClient() (Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

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
