package handlers

import (
	"Jameson/pkg/model"
	mock_service "Jameson/pkg/service/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateProject(t *testing.T) {
	type mockBehavior func(s mock_service.MockImageService, project model.Project)

	tests := []struct {
		name                 string
		inputBody            string
		inputProject         model.Project
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"id":"id","name": "testname"}`,
			inputProject: model.Project{
				Name: "testname",
				ID:   "id",
			},
			mockBehavior: func(s mock_service.MockImageService, project model.Project) {
				s.EXPECT().CreateProject(project).Return(&project, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":"id","name":"testname"}`,
		},
		{
			name:         "missing field name",
			inputBody:    "",
			inputProject: model.Project{},
			mockBehavior: func(s mock_service.MockImageService, project model.Project) {
				s.EXPECT().CreateProject(project).Return(nil, errors.New("field name is required"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"cannot create project","error":"field name is required"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_service.NewMockImageService(c)
			handler := Handler{service}
			test.mockBehavior(*service, test.inputProject)
			// Init Endpoint
			r := gin.New()
			r.POST("/api/v1/projects", handler.CreateProject)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/projects", bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			expect := test.expectedResponseBody
			actual := w.Body.String()

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, actual, expect)
		})
	}

}
