package config

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"syscall"
	"testing"
	"time"
)

func TestDBConnection(t *testing.T) {
	if err := godotenv.Load("../../../test.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	host, ok := syscall.Getenv("MONGODB_HOST")
	if !ok {
		t.Fatalf("Error getting MONGODB_HOST")
	}
	port, ok := syscall.Getenv("MONGODB_PORT")
	if !ok {
		t.Fatalf("Error getting MONGODB_PORT")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("TestDBConnection001", func(t *testing.T) {
		uri := "mongodb://" + host + ":" + port
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
