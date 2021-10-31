package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ImageHandler struct {
	Service MongoImageService
}

func (h ImageHandler) GetProjects(c *gin.Context) {
	projects := h.Service.GetProjects()
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, projects)
}

func (h ImageHandler) GetContainers(c *gin.Context) {
	containers := h.Service.GetContainers()
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, containers)
}

func (h ImageHandler) PerformTest(c *gin.Context) {
	containerId := c.Param("container")
	container, _ := h.Service.GetContainerById(containerId)

	reference := h.Service.DownloadImage(container.ReferenceId + ".png")
	candidate, err := excludeFileBytes(c)
	if err != nil {
		c.String(422, "cannot exclude file from request body")
	}

	resultImage, percentage := GetImageDifference(reference, candidate)
	candidateId := h.Service.UploadImage(candidate)
	resultId := h.Service.UploadImage(resultImage)

	testResult := TestResult{ID: resultId, Percentage: percentage}

	h.Service.WritingTestResultToContainer(container.ID, Test{CandidateId: candidateId,
		Result: TestResult{ID: resultId, Percentage: percentage}})

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, testResult)
}

func (h ImageHandler) GetContainerByName(c *gin.Context) {
	testName := c.Request.URL.Query().Get("test_name")
	container, isExists := h.Service.GetContainerByName(testName)
	c.Header("content-type", "application/json")

	if isExists {
		c.JSON(http.StatusOK, container)
	}
}

func (h ImageHandler) GetImage(c *gin.Context) {
	imageId := c.Param("image")
	buff := h.Service.DownloadImage(imageId + ".png")
	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", strconv.Itoa(len(buff)))

	_, err := c.Writer.Write(buff)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
	}
}
func (h ImageHandler) ApproveReference(c *gin.Context) {
	containerId := c.Param("container")
	h.Service.ApproveReferenceForContainer(containerId)

}
func (h ImageHandler) CreateTest(c *gin.Context) {
	projectId := c.Param("project")
	testName := c.Request.URL.Query().Get("test_name")
	testContainer := TestContainer{
		ID: GetNewId(), ProjectId: projectId, Tests: []Test{}, Name: testName,
	}
	testContainer.ProjectId = projectId

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Fatal(err)
	}
	filename := h.Service.UploadImage(buf.Bytes())
	testContainer.ReferenceId = filename
	h.Service.CreateNewTestContainer(testContainer)

	c.JSON(http.StatusOK, testContainer)
}

func (h ImageHandler) CreateProject(c *gin.Context) {
	project := Project{}
	err := json.NewDecoder(c.Request.Body).Decode(&project)

	if err != nil {
		c.String(http.StatusUnprocessableEntity, "invalid body")
		return
	}
	h.Service.CreateProject(project)
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
