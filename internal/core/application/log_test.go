package application

import (
	config2 "example/internal/core/config"
	"github.com/joho/godotenv"
	"testing"
)

func TestLogger(t *testing.T) {
	if err := godotenv.Load("../../../test.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	config, err := config2.NewAppConfig()
	if err != nil {
		t.Fatalf("Error creating app config: %v", err)
	}
	app, err := NewApplication(config)
	if err != nil {
		t.Fatalf("Error creating application: %v", err)
	}

	t.Run("TestLoggerInfo001", func(t *testing.T) {
		app.Logger().Info("This is an info message for logger test")
	})

	t.Run("TestLoggerError005", func(t *testing.T) {
		app.Logger().Error("This is an error message for logger test")
	})
}
