package handlers

import (
	mdl "Jameson/pkg/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary get all containers
// @ID get_containers
// @Accept  json
// @Produce  json
// @Success 200 {object} []Project
// @Failure 422,404 {object} errorResponse
// @Failure 500 {object} string
// @Router /containers [get]
func (h *Handler) GetProjects(c *gin.Context) {
	projects := h.Service.GetProjects()
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, projects)
}

func (h *Handler) CreateProject(c *gin.Context) {
	project := mdl.Project{}
	err := json.NewDecoder(c.Request.Body).Decode(&project)

	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}
	createdProject, err := h.Service.CreateProject(project)
	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot create project", err)
		return
	}

	c.JSON(http.StatusOK, createdProject)
}
