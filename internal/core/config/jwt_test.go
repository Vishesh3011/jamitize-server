package config

import (
	"example/internal/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestJWTConfig(t *testing.T) {
	if err := godotenv.Load("../../../test.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	config, err := newJWTConfig()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("TestGenerateAndVerifyJWT", func(t *testing.T) {
		tokenStr, err := utils.GenerateJWT(primitive.NewObjectID(), config.secret)
		if err != nil {
			t.Fatal(err)
		}

		err = utils.VerifyJWT(tokenStr, config.secret)
		if err != nil {
			t.Fatal(err)
		}
	})
}
