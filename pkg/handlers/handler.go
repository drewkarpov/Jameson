package handlers

import (
	_ "Jameson/docs"
	"Jameson/pkg/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	Service service.ImageService
}

func NewHandler(services service.ImageService) *Handler {
	return &Handler{Service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.GET("/projects", h.GetProjects)
		api.GET("/containers", h.GetContainers)
		api.GET("/image/:image", h.GetImage)

		api.POST("/project/create", h.CreateProject)
		api.POST("/project/:project/test/create", h.CreateNewTestContainer)
		api.POST("/container/:container/perform/test", h.PerformTest)
		api.PUT("/container/:container/approve", h.ApproveReference)
	}
	return router
}
