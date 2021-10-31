package main

import (
	cfg "Jameson/config"
	"Jameson/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config := cfg.InitConfig("mongo_db")
	mongoService := pkg.InitMongoService(config)
	shutdown := make(chan error, 1)

	router := gin.Default()
	imgHandler := pkg.ImageHandler{mongoService}
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.StaticFS("./images", http.Dir("images"))

	api := router.Group("/api/v1")
	{
		api.GET("/image/:image", imgHandler.GetImage)
		api.GET("/projects", imgHandler.GetProjects)
		api.GET("/containers", imgHandler.GetContainers)
		api.GET("/container", imgHandler.GetContainerByName)
		api.PUT("/container/:container/approve", imgHandler.ApproveReference)
		api.POST("/container/:container/perform/test", imgHandler.PerformTest)
		api.POST("/create/project", imgHandler.CreateProject)
		api.POST("/project/:project/test/init", imgHandler.CreateTest)
	}

	err := http.ListenAndServe(":3333", router)
	shutdown <- err

}
