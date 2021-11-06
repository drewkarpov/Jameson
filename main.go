package main

import (
	_ "Jameson/docs"
	"Jameson/pkg/app"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

// @host localhost:3333
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	application := app.InitApplication()

	shutdown := make(chan error, 1)

	go func() {
		err := http.ListenAndServe(":3333", application.Handler.InitRoutes())
		shutdown <- err
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
