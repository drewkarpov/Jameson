package main

import (
	"Jameson/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

const mainDirectory = "./images"

func main() {
	mongoService := pkg.InitMongoService()
	shutdown := make(chan error, 1)

	router := gin.Default()
	imgHandler := pkg.ImageHandler{mongoService}
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.StaticFS("./images", http.Dir("images"))

	api := router.Group("/api/v1")
	{
		api.GET("/result", imgHandler.GetResult)
		api.GET("/image/:image", imgHandler.GetImage)
		api.GET("/projects", imgHandler.GetProjects)
		api.GET("/containers", imgHandler.GetContainers)
		api.GET("/container", imgHandler.GetContainerByName)
		api.PUT("/container/:container/approve", imgHandler.ApproveReference)
		api.POST("/create/project", imgHandler.CreateProject)
		api.POST("/project/:project/test/init", imgHandler.CreateTest)
	}

	err := http.ListenAndServe(":3333", router)
	shutdown <- err

}
