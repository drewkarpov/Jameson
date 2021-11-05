package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ImageHandler struct {
	Service MongoImageService
}

// @Summary image
// @Description get image by id
// @ID set image
// @Accept  json
// @Produce  json
// @Accept  multipart/form-data
// @Param image path string true "image"
// @Success 200
// @Failure 422,404 {object} errorResponse
// @Failure 500 {object} string
// @Failure default {object} string
// @Router /image/{image} [get]
func (h ImageHandler) GetImage(c *gin.Context) {
	imageId := c.Param("image")
	buff, err := h.Service.DownloadImage(imageId + ".png")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "cannot download image from db", err)
		return
	}

	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", strconv.Itoa(len(buff)))

	_, err = c.Writer.Write(buff)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "cannot writing image bytes to response", err)
	}
}
