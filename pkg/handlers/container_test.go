package handlers

import (
	"Jameson/pkg/model"
	mockservice "Jameson/pkg/service/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
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
			r.GET("/container/approve", handler.ApproveReference)

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
