package handlers

import (
	"bytes"
	mockservice "github.com/drewkarpov/Jameson/pkg/service/mocks"
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
	type mockBehavior func(s mockservice.MockImageService)

	tests := []struct {
		name               string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedBuffLength int
	}{
		{
			name:               "images difference by bounds",
			expectedStatusCode: 200,
			expectedBuffLength: 11550,
			mockBehavior: func(s mockservice.MockImageService) {
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
			test.mockBehavior(*service)

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
