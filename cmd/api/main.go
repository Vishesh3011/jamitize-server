package main

import (
	"example/errors"
	"example/internal/core/application"
	"example/internal/core/config"
	"example/internal/routes"
	"example/internal/types"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(errors.ToAppError(err, types.InternalServerError, types.Application).Error())
	}

	appConfig, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(errors.ToAppError(err, types.InternalServerError, types.Application).Error())
	}

	app, err := application.NewApplication(appConfig)
	if err != nil {
		log.Fatal(errors.ToAppError(err, types.InternalServerError, types.Application).Error())
	}

	if err := routes.NewServer(app).Start(); err != nil {
		log.Fatal(errors.ToAppError(err, types.InternalServerError, types.Application).Error())
	}
	
}
