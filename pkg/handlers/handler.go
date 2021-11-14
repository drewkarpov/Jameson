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

func NewHandler(service service.ImageService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.GET("/projects", h.GetProjects)
		api.GET("/containers", h.GetContainers)
		api.GET("/image/:image", h.GetImage)

		container := api.Group("/container")
		{
			container.POST("/:container/perform/test", h.PerformTest)
			container.PATCH("/:container/approve", h.ApproveReference)
			container.PATCH("/:container/change/reference", h.SetNewReference)
			container.DELETE("/:container/delete", h.DeleteContainer)
		}

		project := api.Group("/project")
		{
			project.POST("/create", h.CreateProject)
			project.POST("/:project/test/create", h.CreateNewTestContainer)

		}
	}
	return router
}
