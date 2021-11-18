package handlers

import (
	_ "github.com/drewkarpov/Jameson/docs"
	"github.com/drewkarpov/Jameson/pkg/service"
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
	router.Use(CORSMiddleware())
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
