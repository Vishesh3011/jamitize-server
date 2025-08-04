package clients

import (
	"context"
	"example/internal/core/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Client interface {
	DBClient() *mongo.Client
}

type client struct {
	*mongo.Client
}

func NewClient(appConfig config.AppConfig) (Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	mongoUri := "mongodb://" + appConfig.DBConfig().Username() + ":" + appConfig.DBConfig().Password() + "@" + appConfig.DBConfig().Host() + ":" + appConfig.DBConfig().Port() + "/" + appConfig.DBConfig().Database() + "?authSource=" + appConfig.DBConfig().Database()
	clientOptions := options.Client().ApplyURI(mongoUri)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
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
