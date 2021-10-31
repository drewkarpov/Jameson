package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ImageHandler struct {
	Service MongoImageService
}

func (h ImageHandler) GetResult(c *gin.Context) {

	buff := h.Service.DownloadImage("ref1.png")
	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", strconv.Itoa(len(buff)))
	wrapper := ImageWrapper{}
	wrapper.SetReference("./images/ref2.png")

	img, _, err := image.Decode(bytes.NewReader(buff))
	if err != nil {
		log.Fatalln(err)
	}
	img2, perc := GetImageDifference(wrapper.Reference.Body, img)

	println(perc)
	buf := new(bytes.Buffer)
	errr := png.Encode(buf, img2)

	if errr != nil {
		log.Fatalln(err)
	}
	send_s3 := buf.Bytes()

	c.Writer.Write(send_s3)
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
	filename := GetNewId()
	testContainer := TestContainer{
		ID: GetNewId(), ProjectId: projectId, Tests: []Test{}, Name: testName,
	}
	testContainer.ProjectId = projectId
	testContainer.Tests = []Test{}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Fatal(err)
	}
	h.Service.UploadImage(buf.Bytes(), filename)
	testContainer.ReferenceId = filename
	h.Service.WritingTestContainer(testContainer)
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
