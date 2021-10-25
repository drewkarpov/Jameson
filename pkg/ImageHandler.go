package pkg

import (
	"bytes"
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

func (h ImageHandler) GetResult(c *gin.Context) {

	buff := h.Service.DownloadImage("ref1.png")
	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", strconv.Itoa(len(buff)))
	c.Writer.Write(buff)
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
