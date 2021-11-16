package app

import (
	"github.com/drewkarpov/Jameson/config"
	cfg "github.com/drewkarpov/Jameson/config"
	handler "github.com/drewkarpov/Jameson/pkg/handlers"
	srv "github.com/drewkarpov/Jameson/pkg/service"
)

type Application struct {
	Config  config.Config
	Service srv.ImageService
	Handler *handler.Handler
}

func InitApplication() *Application {
	mongoConfig := cfg.InitConfig("mongo_db")
	mongoService := srv.InitMongoService(mongoConfig)
	imageHandler := handler.NewHandler(&mongoService)
	return &Application{Config: mongoConfig,
		Service: &mongoService,
		Handler: imageHandler}
}
