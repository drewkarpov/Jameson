package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/drewkarpov/Jameson/pkg/image"
	img "github.com/drewkarpov/Jameson/pkg/image"
	mdl "github.com/drewkarpov/Jameson/pkg/model"
	"github.com/drewkarpov/Jameson/pkg/utils"
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
// @Router /container/{container}/approve [patch]
func (h *Handler) ApproveReference(c *gin.Context) {
	containerId := c.Param("container")
	update, err := h.Service.ApproveReferenceForContainer(containerId)
	if err != nil || !update {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot approve reference for container id "+containerId, err)
		return
	}
	mdl.NewSuccessResponse(c, "reference for container "+containerId+" is approved")
}

// @Summary set new reference for container
// @Description string container_id
// @ID set new_reference_container
// @Accept  json
// @Produce  json
// @Param container path string true "container_id"
// @Param mock body mdl.Reference true "body"
// @Success 200 {object} mdl.successResponse
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /container/{container}/change/reference [patch]
func (h *Handler) SetNewReference(c *gin.Context) {
	containerId := c.Param("container")
	var ref mdl.Reference
	err := json.NewDecoder(c.Request.Body).Decode(&ref)

	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	update, err := h.Service.SetNewReferenceForContainer(containerId, ref)
	if err != nil || !update {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot set new reference reference for container id "+containerId, err)
		return
	}
	mdl.NewSuccessResponse(c, "reference for container "+containerId+" is changed")
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

	if container == nil || !container.Approved {
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

	var resultImage []byte
	var percentage float64

	if container.VoidZones != nil && len(container.VoidZones) > 0 {
		resultImage, percentage, err = image.GetImageDifference(
			img.GetImageWithVoidZones(reference, container.VoidZones),
			img.GetImageWithVoidZones(candidate, container.VoidZones))
	} else {
		resultImage, percentage, err = image.GetImageDifference(reference, candidate)
	}

	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "images have difference by bound", err)
		return
	}

	result, err := h.Service.WritingTestResultToContainer(candidate, resultImage, percentage, container.ID)
	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot writing test to container", err)
		return
	}

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, result)
}

// @Summary container by id
// @ID set container by id
// @Accept  json
// @Produce  json
// @Param container path string true "container_id"
// @Success 200 {object} mdl.TestContainer
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /container/{container} [get]
func (h *Handler) GetContainerById(c *gin.Context) {
	containerId := c.Param("container")
	container, _ := h.Service.GetContainerById(containerId)

	if container == nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot find container with id "+containerId, nil)
		return
	}
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, container)
}

// @Summary all containers
// @ID get_containers
// @Accept  json
// @Produce  json
// @Success 200 {object} []mdl.TestContainer
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /containers [get]
func (h *Handler) GetContainers(c *gin.Context) {
	containers := h.Service.GetContainers()
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, containers)
}

// @Summary all projects
// @ID delete_container
// @Accept  json
// @Produce  json
// @Param container path string true "container_id"
// @Success 200 {object} mdl.successResponse
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /container/{container}/delete [delete]
func (h *Handler) DeleteContainer(c *gin.Context) {
	containerId := c.Param("container")
	isSuccess, err := h.Service.DeleteContainerById(containerId)

	if err != nil || !isSuccess {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot delete container from db", err)
		return
	}
	mdl.NewSuccessResponse(c, "container with id "+containerId+" is deleted")
}

// @Summary get result test data
// @Description get test by id
// @ID get result test data
// @Accept  json
// @Produce  json
// @Param test path string true "test"
// @Success 200 {object} mdl.ResultContainer
// @Failure 422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Failure default {object} string
// @Router /test/{test} [get]
func (h *Handler) GetPreparedTestData(c *gin.Context) {
	testId := c.Param("test")
	container, isExists := h.Service.GetContainerByTestId(testId)

	if container == nil || !isExists {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot find container", nil)
		return
	}

	var result mdl.ResultContainer
	for _, test := range container.Tests {
		if test.ID == testId {
			result = mdl.ResultContainer{
				Percentage: test.Result.Percentage,
				Images:     mdl.ImagesContainer{DiffId: test.Result.ID, CandidateId: test.CandidateId, ReferenceId: container.ReferenceId},
			}
		}
	}

	if result.Images.DiffId == "" {
		mdl.NewErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("cannot test with id %s", testId), nil)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary add voidzone for reference
// @Description string container_id
// @ID add voidzone for reference
// @Accept  json
// @Produce  json
// @Param container path string true "container_id"
// @Param voidzones_array body []img.VoidZone true "body"
// @Success 200 {object} mdl.successResponse
// @Failure 400,422,404 {object} mdl.errorResponse
// @Failure 500 {object} string
// @Router /container/{container}/add/voidzone [post]
func (h *Handler) AddVoidZoneForReference(c *gin.Context) {
	containerId := c.Param("container")
	var voidZone []img.VoidZone
	err := json.NewDecoder(c.Request.Body).Decode(&voidZone)

	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	err = h.Service.AddVoidZonesForReference(containerId, voidZone)
	if err != nil {
		mdl.NewErrorResponse(c, http.StatusBadRequest, "cannot set new reference reference for container id "+containerId, err)
		return
	}
	mdl.NewSuccessResponse(c, "add void zone for container reference")
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
