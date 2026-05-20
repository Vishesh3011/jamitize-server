package test

import (
	"context"
	"example/internal/core/application"
	"example/internal/core/config"
	"log"

	"github.com/joho/godotenv"
)

type AppTest interface {
	application.Application
	Context() context.Context
	Cancel() context.CancelFunc
}

type appTest struct {
	application.Application
	ctx    context.Context
	cancel context.CancelFunc
}

func (a appTest) Context() context.Context {
	return a.ctx
}

func (a appTest) Cancel() context.CancelFunc {
	return a.cancel
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

	ctx, cancel := context.WithCancel(context.Background())

	return appTest{
		Application: app,
		ctx:         ctx,
		cancel:      cancel,
	}
}
