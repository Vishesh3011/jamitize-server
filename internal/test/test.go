package test

import (
	"context"
	"example/internal/core/application"
	"example/internal/core/config"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type AppTest interface {
	application.Application
	Context() context.Context
}

type appTest struct {
	application.Application
	ctx context.Context
}

func (a appTest) Context() context.Context {
	return a.ctx
}

func NewAppTest(path *string) AppTest {
	envPath := "../../test.env"
	if path != nil {
		envPath = *path
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file")
	}

	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("Error creating app config: %v", err)
	}

	app, err := application.NewApplication(cnfg)
	if err != nil {
		log.Fatalf("Error creating application: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return appTest{
		Application: app,
		ctx:         ctx,
	}
}
