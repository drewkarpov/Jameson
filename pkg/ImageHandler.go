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
	"os"
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
	jsonPessoal, err := json.Marshal(projects)
	if err != nil {
		log.Fatal(err)
	}
	c.Header("content-type", "application/json")
	fmt.Fprintf(os.Stdout, "%s", jsonPessoal) // still fine here .
	// it is fine because you are formating []byte into string using fmt and
	// printing it on console. `%s` makes sures that it echos as string.

	c.JSON(http.StatusOK, projects)
}

func (h ImageHandler) CreateTest(c *gin.Context) {
	testContainer := TestContainerDTO{}
	err := json.NewDecoder(c.Request.Body).Decode(&testContainer)

	if err != nil {
		c.String(http.StatusUnprocessableEntity, "invalid body")
		return
	}
	h.Service.AddTestContainerToProject(testContainer.ProjectId, TestContainer{Name: testContainer.Name, ReferenceId: testContainer.ReferenceId, Tests: []Test{}})
}

func (h ImageHandler) GetOriginImage(c *gin.Context) {
	path := c.Param("path")
	//image := h.Service.GetReference("")
	http.ServeFile(c.Writer, c.Request, "/Users/akarpov/Projects/Jameson/images/"+path+".png")
}

func (h ImageHandler) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Fatal(err)
	}
	h.Service.UploadImage(buf.Bytes(), filename)
	c.JSON(http.StatusOK, gin.H{"filepath": filename})
}
