package main

import (
	"embed"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/drewkarpov/Jameson/docs"
	"github.com/drewkarpov/Jameson/pkg/app"
)

// @title Swagger Jameson API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host testing.rezero.pro
// @BasePath /api/v1
// @query.collection.format multi

//go:embed frontend/public/*
var fs embed.FS

func main() {
	application := app.InitApplication()

	shutdown := make(chan error, 1)

	router := application.Handler.InitRoutes(fs)

	go func() {
		err := http.ListenAndServe(":3000", router)
		shutdown <- err
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
