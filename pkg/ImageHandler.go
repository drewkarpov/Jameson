package pkg

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ImageHandler struct {
	Service ImageService
}

func (h ImageHandler) GetResult(c *gin.Context) {

	buffer := new(bytes.Buffer)
	image := h.Service.GetReference("")
	if err := png.Encode(buffer, image); err != nil {
		log.Println("unable to encode image.")
	}

	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", strconv.Itoa(len(buffer.Bytes())))

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
	out, err := os.Create("/Users/akarpov/Projects/Jameson/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}
