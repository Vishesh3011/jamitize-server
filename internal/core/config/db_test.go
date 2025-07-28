package config

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestDBConnection(t *testing.T) {
	if err := godotenv.Load("../../../test.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	config, err := newDBConfig()
	if err != nil {
		t.Fatalf("Error creating DB config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("TestDBConnection001", func(t *testing.T) {
		uri := "mongodb://" + config.database + ":" + config.password + "@" + config.host + ":" + config.port + "/" + config.database + "?authSource=admin"
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			t.Fatalf("Error connecting to MongoDB: %v", err)
		}
		defer func(client *mongo.Client, ctx context.Context) {
			err := client.Disconnect(ctx)
			if err != nil {
				t.Fatalf("Error disconnecting from MongoDB: %v", err)
			}
		}(client, ctx)
	})
}
