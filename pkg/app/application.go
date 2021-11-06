package app

import (
	"Jameson/config"
	cfg "Jameson/config"
	handler "Jameson/pkg/handlers"
	srv "Jameson/pkg/service"
)

type Application struct {
	Config  config.Config
	Service srv.ImageService
	Handler *handler.Handler
}

func InitApplication() *Application {
	mongoConfig := cfg.InitConfig("mongo_db")
	mongoService := srv.InitMongoService(mongoConfig)
	imageHandler := handler.NewHandler(mongoService)
	return &Application{Config: mongoConfig,
		Service: mongoService,
		Handler: imageHandler}
}
