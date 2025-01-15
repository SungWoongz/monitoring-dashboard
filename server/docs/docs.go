// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cpu": {
            "get": {
                "description": "Retrieves CPU status data based on the specified interval and limit.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CPU"
                ],
                "summary": "Get CPU Status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Interval in seconds (default: 1)",
                        "name": "interval",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Number of records to fetch (default: 30)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.HttpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HttpResponse"
                        }
                    }
                }
            }
        },
        "/memory": {
            "get": {
                "description": "Retrieves memory status data based on the specified interval and limit.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Memory"
                ],
                "summary": "Get Memory Status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Interval in seconds (default: 1)",
                        "name": "interval",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Number of records to fetch (default: 30)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.HttpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.HttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.HttpResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "status": {
                    "$ref": "#/definitions/models.HttpStatus"
                }
            }
        },
        "models.HttpStatus": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "example": "OK"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MonDash API",
	Description:      "MonDash API SERVER Swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
