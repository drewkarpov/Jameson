package handlers

import (
	mockservice "Jameson/pkg/service/mocks"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"image"
	"image/png"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler_GetImage(t *testing.T) {
	type mockBehavior func(s mockservice.MockImageService, path string)

	tests := []struct {
		name               string
		imagePath          string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedBuffLength int
	}{
		{
			name:               "images difference by bounds",
			imagePath:          "./ref1.png",
			expectedStatusCode: 200,
			expectedBuffLength: 11550,
			mockBehavior: func(s mockservice.MockImageService, path string) {
				imageBuff, _ := getBytesFromImage("./ref1.png")
				s.EXPECT().DownloadImage("some_path.png").Return(imageBuff, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mockservice.NewMockImageService(c)
			test.mockBehavior(*service, test.imagePath)

			handler := Handler{service}

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(w)
			r := gin.New()
			r.GET("/image/", handler.GetImage)

			context.Params = []gin.Param{
				{
					Key:   "image",
					Value: "some_path",
				},
			}

			handler.GetImage(context)

			expect := test.expectedBuffLength
			actual := len(w.Body.Bytes())

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, actual, expect)
		})
	}
}

func getBytesFromImage(filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, _ := image.Decode(f)

	buff := new(bytes.Buffer)
	err2 := png.Encode(buff, img)
	if err2 != nil {
		return nil, err2
	}
	return buff.Bytes(), nil
}
