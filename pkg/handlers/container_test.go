package handlers

import (
	"bytes"
	"errors"
	"github.com/drewkarpov/Jameson/pkg/model"
	mockservice "github.com/drewkarpov/Jameson/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler_GetContainers(t *testing.T) {
	type mockBehavior func(s mockservice.MockImageService)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "POSITIVE",
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().GetContainers().Return([]model.TestContainer{
					{
						ID:          "43141",
						Name:        "project_name",
						ReferenceId: "reference",
						Approved:    false,
						Tests:       []model.Test{},
					},
				})
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"id":"43141","name":"project_name","project_id":"","reference_id":"reference","approved":false,"tests":[]}]`,
		},
		{
			name: "POSITIVE EMPTY",
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().GetContainers().Return([]model.TestContainer{})
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[]`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mockservice.NewMockImageService(c)
			handler := Handler{service}
			test.mockBehavior(*service)
			r := gin.New()
			r.GET("/api/v1/containers", handler.GetContainers)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/containers", nil)

			r.ServeHTTP(w, req)

			expect := test.expectedResponseBody
			actual := w.Body.String()

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, actual, expect)
		})
	}
}

func TestHandler_ApproveReference(t *testing.T) {
	type mockBehavior func(s mockservice.MockImageService)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "positive",
			expectedStatusCode:   200,
			expectedResponseBody: `{"result":"reference for container some-id is approved"}`,
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().ApproveReferenceForContainer("some-id").Return(true, nil)
			},
		},
		{
			name:                 "return false without error",
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot approve reference for container id some-id","error":"\u003cnil\u003e"}`,
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().ApproveReferenceForContainer("some-id").Return(false, nil)
			},
		},
		{
			name:                 "return true with error",
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot approve reference for container id some-id","error":"some error"}`,
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().ApproveReferenceForContainer("some-id").Return(true, errors.New("some error"))
			},
		},
		{
			name:                 "return false with error",
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot approve reference for container id some-id","error":"some error"}`,
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().ApproveReferenceForContainer("some-id").Return(false, errors.New("some error"))
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
			r.PATCH("/container/approve", handler.ApproveReference)

			context.Params = []gin.Param{
				{
					Key:   "container",
					Value: "some-id",
				},
			}

			handler.ApproveReference(context)

			expect := test.expectedResponseBody
			actual := w.Body.String()

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, actual, expect)
		})
	}
}

func TestHandler_SetNewReference(t *testing.T) {
	type mockBehavior func(s mockservice.MockImageService, ref model.Reference)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		inputBody            string
		inputReference       model.Reference
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "positive",
			inputBody:            `{"reference_id": "some_id"}`,
			inputReference:       model.Reference{ID: "some_id"},
			expectedStatusCode:   200,
			expectedResponseBody: `{"result":"reference for container some-id is changed"}`,
			mockBehavior: func(s mockservice.MockImageService, ref model.Reference) {
				s.EXPECT().SetNewReferenceForContainer("some-id", ref).Return(true, nil)
			},
		},
		{
			name:                 "error encode body , incorrect json",
			inputBody:            `nil`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid request body","error":"invalid character 'i' in literal null (expecting 'u')"}`,
			mockBehavior:         func(s mockservice.MockImageService, ref model.Reference) {},
		},
		{
			name:                 "error encode body , incorrect type of reference id - int",
			inputBody:            `{"reference_id": 3434}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid request body","error":"json: cannot unmarshal number into Go struct field Reference.reference_id of type string"}`,
			mockBehavior:         func(s mockservice.MockImageService, ref model.Reference) {},
		},
		{
			name:                 "error encode body , incorrect type of reference id - bool",
			inputBody:            `{"reference_id": true}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid request body","error":"json: cannot unmarshal bool into Go struct field Reference.reference_id of type string"}`,
			mockBehavior:         func(s mockservice.MockImageService, ref model.Reference) {},
		},
		{
			name:                 "return false without error",
			inputBody:            `{"reference_id": "some_id"}`,
			inputReference:       model.Reference{ID: "some_id"},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot set new reference reference for container id some-id","error":"\u003cnil\u003e"}`,
			mockBehavior: func(s mockservice.MockImageService, ref model.Reference) {
				s.EXPECT().SetNewReferenceForContainer("some-id", ref).Return(false, nil)
			},
		},
		{
			name:                 "return false with error",
			inputBody:            `{"reference_id": "some_id"}`,
			inputReference:       model.Reference{ID: "some_id"},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot set new reference reference for container id some-id","error":"some error"}`,
			mockBehavior: func(s mockservice.MockImageService, ref model.Reference) {
				s.EXPECT().SetNewReferenceForContainer("some-id", ref).Return(false, errors.New("some error"))
			},
		},
		{
			name:                 "return true with error",
			inputBody:            `{"reference_id": "some_id"}`,
			inputReference:       model.Reference{ID: "some_id"},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot set new reference reference for container id some-id","error":"some error"}`,
			mockBehavior: func(s mockservice.MockImageService, ref model.Reference) {
				s.EXPECT().SetNewReferenceForContainer("some-id", ref).Return(true, errors.New("some error"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mockservice.NewMockImageService(c)
			test.mockBehavior(*service, test.inputReference)

			handler := Handler{service}

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(w)
			r := gin.New()
			r.PATCH("/container/new/reference", handler.SetNewReference)

			context.Params = []gin.Param{
				{
					Key:   "container",
					Value: "some-id",
				},
			}
			req := httptest.NewRequest("PATCH", "/container/new/reference", bytes.NewBufferString(test.inputBody))

			context.Request = req

			handler.SetNewReference(context)

			expect := test.expectedResponseBody
			actual := w.Body.String()

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, actual, expect)
		})
	}
}

func TestHandler_DeleteContainer(t *testing.T) {
	type mockBehavior func(s mockservice.MockImageService)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "POSITIVE",
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().DeleteContainerById("some-id").Return(true, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"result":"container with id some-id is deleted"}`,
		},
		{
			name: "returned true with error",
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().DeleteContainerById("some-id").Return(true, errors.New("some error"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot delete container from db","error":"some error"}`,
		},
		{
			name: "returned false without error",
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().DeleteContainerById("some-id").Return(false, nil)
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot delete container from db","error":"\u003cnil\u003e"}`,
		},
		{
			name: "returned false with error",
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().DeleteContainerById("some-id").Return(false, errors.New("some error"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot delete container from db","error":"some error"}`,
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
			r.DELETE("/container/new/reference", handler.SetNewReference)

			context.Params = []gin.Param{
				{
					Key:   "container",
					Value: "some-id",
				},
			}

			handler.DeleteContainer(context)

			expect := test.expectedResponseBody
			actual := w.Body.String()

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, actual, expect)
		})
	}
}

func TestHandler_PerformTest(t *testing.T) {
	type mockBehavior func(s mockservice.MockImageService, candidate []byte, resultId *string)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		resultID             string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "POSITIVE",
			resultID: "test-1",
			mockBehavior: func(s mockservice.MockImageService, candidate []byte, resultId *string) {
				s.EXPECT().GetContainerById("some-id").Return(
					&model.TestContainer{
						ID:          "some-id",
						ProjectId:   "project-id",
						ReferenceId: "reference-id",
						Approved:    true,
						Tests:       []model.Test{},
					}, true)

				s.EXPECT().DownloadImage("reference-id.png").Return(candidate, nil)
				s.EXPECT().UploadImage(candidate).Return(resultId, nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mockservice.NewMockImageService(c)

			handler := Handler{service}

			gin.SetMode(gin.TestMode)

			r := gin.New()
			r.POST("/container/perform/test", handler.PerformTest)

			w := httptest.NewRecorder()

			context, _ := gin.CreateTestContext(w)
			context.Params = []gin.Param{
				{
					Key:   "container",
					Value: "some-id",
				},
			}

			b, wr := createMultipartFormData(t, "file", "./ref1.png")
			req := httptest.NewRequest("POST", "/container/perform/test", &b)
			req.Header.Set("Content-Type", wr.FormDataContentType())

			context.Request = req

			bt, _ := excludeFileBytes(context)

			test.mockBehavior(*service, bt, &test.resultID)

			handler.PerformTest(context)

			expect := test.expectedResponseBody
			actual := w.Body.String()

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, actual, expect)
		})
	}
}

func createMultipartFormData(t *testing.T, fieldName, fileName string) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file := mustOpen(fileName)
	if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		t.Errorf("Error creating writer: %v", err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		t.Errorf("Error with io.Copy: %v", err)
	}
	w.Close()
	return b, w
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
	}
	return r
}
