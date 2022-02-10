package handlers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/drewkarpov/Jameson/pkg/model"
	mockservice "github.com/drewkarpov/Jameson/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateProject(t *testing.T) {
	type mockBehavior func(s mockservice.MockImageService, project model.Project)

	tests := []struct {
		name                 string
		inputBody            string
		inputProject         model.Project
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "POSITIVE",
			inputBody: `{"id":"id","name": "testname"}`,
			inputProject: model.Project{
				Name: "testname",
				ID:   "id",
			},
			mockBehavior: func(s mockservice.MockImageService, project model.Project) {
				s.EXPECT().CreateProject(&project).Return(&project, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":"id","name":"testname"}`,
		},
		{
			name:                 "decode body error - empty",
			inputBody:            "",
			inputProject:         model.Project{},
			mockBehavior:         func(s mockservice.MockImageService, project model.Project) {},
			expectedStatusCode:   422,
			expectedResponseBody: `{"message":"cannot decode body","error":"EOF"}`,
		},
		{
			name:                 "decode body error - try nil",
			inputBody:            "nil",
			inputProject:         model.Project{},
			mockBehavior:         func(s mockservice.MockImageService, project model.Project) {},
			expectedStatusCode:   422,
			expectedResponseBody: `{"message":"cannot decode body","error":"invalid character 'i' in literal null (expecting 'u')"}`,
		},
		{
			name:                 "decode body error , incorrect value type for field name - bool",
			inputBody:            `{"id":"id","name": true}`,
			inputProject:         model.Project{},
			mockBehavior:         func(s mockservice.MockImageService, project model.Project) {},
			expectedStatusCode:   422,
			expectedResponseBody: `{"message":"cannot decode body","error":"json: cannot unmarshal bool into Go struct field Project.name of type string"}`,
		},
		{
			name:                 "decode body error , incorrect value type for field name - int",
			inputBody:            `{"id":"id","name": 383843}`,
			inputProject:         model.Project{},
			mockBehavior:         func(s mockservice.MockImageService, project model.Project) {},
			expectedStatusCode:   422,
			expectedResponseBody: `{"message":"cannot decode body","error":"json: cannot unmarshal number into Go struct field Project.name of type string"}`,
		},
		{
			name:      "empty field name",
			inputBody: `{"id":"id"}`,
			inputProject: model.Project{
				ID: "id",
			},
			mockBehavior:         func(s mockservice.MockImageService, project model.Project) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"field name is required","error":"invalid value for name, field name is required"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			service := mockservice.NewMockImageService(c)
			handler := Handler{service}
			test.mockBehavior(*service, test.inputProject)
			// Init Endpoint
			r := gin.New()
			r.POST("/api/v1/project/create", handler.CreateProject)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/project/create", bytes.NewBufferString(test.inputBody))

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

func TestHandler_GetProjects(t *testing.T) {
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
				s.EXPECT().GetProjects().Return([]model.Project{{ID: "43141", Name: "project_name"}}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"id":"43141","name":"project_name"}]`,
		},
		{
			name: "POSITIVE EMPTY",
			mockBehavior: func(s mockservice.MockImageService) {
				s.EXPECT().GetProjects().Return([]model.Project{}, nil)
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
			r.GET("/api/v1/projects", handler.GetProjects)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/projects", nil)

			r.ServeHTTP(w, req)

			expect := test.expectedResponseBody
			actual := w.Body.String()

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, actual, expect)
		})
	}
}
