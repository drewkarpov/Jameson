package main

import (
	"Jameson/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const mainDirectory = "./images"

func main() {
	if _, err := os.Stat(mainDirectory); os.IsNotExist(err) {
		err := os.Mkdir(mainDirectory, os.ModeDir)
		fmt.Errorf("cannot create folder\n error : %s", err)
	}

	mongoService := pkg.InitMongoService()

	//mongoService.CreateProject("huntbox")

	wrapper := pkg.ImageWrapper{}
	wrapper.SetReference("./images/refs4.png")
	wrapper.SetCandidate("./images/ref2.png")

	//file := "./images/ref3.png"
	//filename := path.Base("ref1.png")

	//mongoService.UploadImage(wrapper.GetReferenceBytes(),filename)
	//mongoService.DownloadImage(filename)
	shutdown := make(chan error, 1)

	router := gin.Default()
	imgHandler := pkg.ImageHandler{mongoService}
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.StaticFS("./images", http.Dir("images"))

	api := router.Group("/api/v1")
	{
		api.GET("/path/:path", imgHandler.GetOriginImage)
		api.GET("/result", imgHandler.GetResult)
		api.GET("/projects", imgHandler.GetProjects)
		api.POST("/upload", imgHandler.Upload)
	}

	err := http.ListenAndServe(":3333", router)
	shutdown <- err

}
