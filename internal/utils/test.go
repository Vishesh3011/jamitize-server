package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collections = []string{
	"users",
}

func CleanupCollections(ctx context.Context, db *mongo.Database) error {
	for _, c := range collections {
		if _, err := db.Collection(c).DeleteMany(ctx, bson.M{}); err != nil {
			return err
		}
	}
	return nil
}
