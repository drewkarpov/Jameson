package main

import (
	cfg "Jameson/config"
	"Jameson/pkg"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "Jameson/docs"
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
	config := cfg.InitConfig("mongo_db")
	mongoService := pkg.InitMongoService(config)
	shutdown := make(chan error, 1)

	router := gin.Default()
	imgHandler := pkg.ImageHandler{Service: mongoService}
	projectHandler := pkg.ProjectHandler{Service: mongoService}
	testHandler := pkg.TestHandler{Service: mongoService}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.GET("/projects", imgHandler.GetProjects)
		api.GET("/containers", testHandler.GetContainers)
		api.GET("/image/:image", imgHandler.GetImage)

		api.POST("/project/create", projectHandler.CreateProject)
		api.POST("/project/:project/test/create", testHandler.CreateNewTestContainer)
		api.POST("/container/:container/perform/test", testHandler.PerformTest)

		api.GET("/container", testHandler.GetContainerByName)
		api.PUT("/container/:container/approve", testHandler.ApproveReference)
	}
	go func() {
		err := http.ListenAndServe(":3333", router)
		shutdown <- err
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
