// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/container/{container}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "container by id",
                "operationId": "set container by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container_id",
                        "name": "container",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TestContainer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/container/{container}/add/voidzone": {
            "post": {
                "description": "string container_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "add voidzone for reference",
                "operationId": "add voidzone for reference",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container_id",
                        "name": "container",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "voidzones_array",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/image.VoidZone"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.successResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/container/{container}/approve": {
            "patch": {
                "description": "string container_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "approve reference for container",
                "operationId": "set container_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container_id",
                        "name": "container",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.successResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/container/{container}/change/reference": {
            "patch": {
                "description": "string container_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "set new reference for container",
                "operationId": "set new_reference_container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container_id",
                        "name": "container",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "mock",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Reference"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.successResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/container/{container}/delete": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "all projects",
                "operationId": "delete_container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container_id",
                        "name": "container",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.successResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/container/{container}/perform/test": {
            "post": {
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "perform test",
                "operationId": "set perform_test",
                "parameters": [
                    {
                        "type": "string",
                        "description": "container_id",
                        "name": "container",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "this is a test file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TestResult"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/containers": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "all containers",
                "operationId": "get_containers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TestContainer"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/image/{image}": {
            "get": {
                "description": "get image by id",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "image",
                "operationId": "set image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image",
                        "name": "image",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/project/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create project",
                "operationId": "creat_project",
                "parameters": [
                    {
                        "description": "body",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Project"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/project/{project}/test/create": {
            "post": {
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create new test container",
                "operationId": "set new_container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project_id",
                        "name": "project",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "test_name",
                        "name": "test_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "this is a test file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TestContainer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/projects": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get all containers",
                "operationId": "get_projects",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Project"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/test/{test}": {
            "get": {
                "description": "get test by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get result test data",
                "operationId": "get result test data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "test",
                        "name": "test",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResultContainer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "image.VoidZone": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer"
                },
                "offset_x": {
                    "type": "integer"
                },
                "offset_y": {
                    "type": "integer"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "model.ImagesContainer": {
            "type": "object",
            "properties": {
                "candidate": {
                    "type": "string"
                },
                "diff": {
                    "type": "string"
                },
                "reference": {
                    "type": "string"
                }
            }
        },
        "model.Project": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Reference": {
            "type": "object",
            "properties": {
                "reference_id": {
                    "type": "string"
                }
            }
        },
        "model.ResultContainer": {
            "type": "object",
            "properties": {
                "images": {
                    "$ref": "#/definitions/model.ImagesContainer"
                },
                "percentage": {
                    "type": "number"
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "model.Test": {
            "type": "object",
            "properties": {
                "candidate_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "result": {
                    "$ref": "#/definitions/model.TestResult"
                }
            }
        },
        "model.TestContainer": {
            "type": "object",
            "properties": {
                "approved": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "reference_id": {
                    "type": "string"
                },
                "tests": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Test"
                    }
                },
                "void_zones": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/image.VoidZone"
                    }
                }
            }
        },
        "model.TestResult": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "percentage": {
                    "type": "number"
                }
            }
        },
        "model.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.successResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "testing.rezero.pro",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Swagger Jameson API",
	Description: "This is a sample server celler server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
