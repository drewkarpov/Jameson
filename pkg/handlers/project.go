package handlers

import (
	"encoding/json"
	"errors"
	mdl "github.com/drewkarpov/Jameson/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary get all containers
// @ID get_projects
// @Accept  json
// @Produce  json
// @Success 200 {object} []mdl.Project
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /projects [get]
func (h *Handler) GetProjects(c *gin.Context) {
	projects, err := h.Service.GetProjects()
	if err != nil {
		mdl.NewErrorResponse(c, http.StatusUnprocessableEntity, "cannot exclude projects from db", err)
	}

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, projects)
}

// @Summary create project
// @ID creat_project
// @Accept  json
// @Produce  json
// @Success 200 {object} mdl.Project
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /api/v1/project/create [post]
func (h *Handler) CreateProject(c *gin.Context) {
	project := mdl.Project{}
	err := json.NewDecoder(c.Request.Body).Decode(&project)

	if err != nil {
		mdl.NewErrorResponse(c, http.StatusUnprocessableEntity, "cannot decode body", err)
		return
	}
	if project.Name == "" {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "field name is required", errors.New("invalid value for name, field name is required"))
		return
	}
	createdProject, err := h.Service.CreateProject(&project)

	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot create project", err)
		return
	}

	c.JSON(http.StatusOK, createdProject)
}
