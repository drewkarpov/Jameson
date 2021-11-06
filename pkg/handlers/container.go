package handlers

import (
	mdl "Jameson/pkg/model"
	"Jameson/pkg/utils"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// @Summary create new test container
// @ID set new_container
// @Accept  json
// @Produce  json
// @Accept  multipart/form-data
// @Param project path string true "project_id"
// @Param test_name query string true "test_name"
// @Param   file formData file true  "this is a test file"
// @Success 200 {object} mdl.TestContainer
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /project/{project}/test/create [post]
func (h *Handler) CreateNewTestContainer(c *gin.Context) {
	projectId := c.Param("project")
	testName := c.Request.URL.Query().Get("test_name")
	testContainer := mdl.TestContainer{ID: utils.GetNewId(), ProjectId: projectId, Tests: []mdl.Test{}, Name: testName}
	_, isExists := h.Service.GetContainerByName(testName)
	if isExists {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "container with name "+testName+" is exist", nil)
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot processing image from request", err)
		return
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot processing image from request", err)
	}
	filename, err := h.Service.UploadImage(buf.Bytes())
	if err != nil {
		c.String(http.StatusBadRequest, "cannot upload file to db")
		return
	}
	testContainer.ReferenceId = *filename
	createdContainer, err := h.Service.CreateNewTestContainer(testContainer)
	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot create container", err)
		return
	}

	c.JSON(http.StatusOK, createdContainer)
}

// @Summary approve reference for container
// @Description string container_id
// @ID set container_id
// @Accept  json
// @Produce  json
// @Param container path string true "container_id"
// @Success 200 {object} mdl.successResponse
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /container/{container}/approve [put]
func (h *Handler) ApproveReference(c *gin.Context) {
	containerId := c.Param("container")
	update, err := h.Service.ApproveReferenceForContainer(containerId)
	if err != nil || !update {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot approve reference for container id "+containerId, err)
		return
	}
	mdl.NewSuccessResponse(c, "reference for container "+containerId+" is approved")
}

// @Summary perform test
// @ID set perform_test
// @Accept  json
// @Produce  json
// @Accept  multipart/form-data
// @Param container path string true "container_id"
// @Param   file formData file true  "this is a test file"
// @Success 200 {object} mdl.TestResult
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /container/{container}/perform/test [post]
func (h *Handler) PerformTest(c *gin.Context) {
	containerId := c.Param("container")
	container, _ := h.Service.GetContainerById(containerId)

	if !container.Approved {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "reference for this container is not approved", nil)
		return
	}

	reference, err := h.Service.DownloadImage(container.ReferenceId + ".png")
	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot download image with id "+container.ReferenceId+" from db", err)
		return
	}

	candidate, err := excludeFileBytes(c)
	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot exclude file from request body", err)
		return
	}

	resultImage, percentage := utils.GetImageDifference(reference, candidate)
	candidateId, err1 := h.Service.UploadImage(candidate)
	if err1 != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot upload  candidate image to db", err1)
		return
	}
	resultId, err2 := h.Service.UploadImage(resultImage)
	if err2 != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot upload  result image to db", err2)
		return
	}
	testResult := mdl.TestResult{ID: *resultId, Percentage: percentage}

	_, err3 := h.Service.WritingTestResultToContainer(container.ID, mdl.Test{CandidateId: *candidateId,
		Result: mdl.TestResult{ID: *resultId, Percentage: percentage}})
	if err3 != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot upload image to db", err3)
		return
	}

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, testResult)
}

// @Summary all projects
// @ID get_projects
// @Accept  json
// @Produce  json
// @Success 200 {object} []mdl.TestContainer
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /projects [get]
func (h *Handler) GetContainers(c *gin.Context) {
	containers := h.Service.GetContainers()
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, containers)
}

func excludeFileBytes(c *gin.Context) ([]byte, error) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
