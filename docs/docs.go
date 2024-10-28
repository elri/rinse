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
        "/jobs": {
            "get": {
                "description": "Get a list of all jobs.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "List jobs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/rinser.Job"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add job with either a file using multipart/form-data or a URL using json.",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Add a job",
                "parameters": [
                    {
                        "description": "Add job by URL",
                        "name": "addjoburl",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/rinser.AddJobURL"
                        }
                    },
                    {
                        "type": "file",
                        "description": "this is a test file",
                        "name": "file",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "eng",
                        "name": "lang",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "2048",
                        "name": "maxsizemb",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "600",
                        "name": "maxtimesec",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "600",
                        "name": "cleanupsec",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "true",
                        "name": "cleanupgotten",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rinser.Job"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "415": {
                        "description": "Unsupported Media Type",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    }
                }
            }
        },
        "/jobs/{uuid}": {
            "get": {
                "description": "Get job metadata by UUID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Get job metadata.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "49d1e304-d2b8-46bf-b6a6-f1e9b797e1b0",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rinser.Job"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by job UUID",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Delete a job",
                "parameters": [
                    {
                        "type": "string",
                        "description": "49d1e304-d2b8-46bf-b6a6-f1e9b797e1b0",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rinser.Job"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    }
                }
            }
        },
        "/jobs/{uuid}/meta": {
            "get": {
                "description": "Get the jobs document metadata.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Get the jobs document metadata.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "49d1e304-d2b8-46bf-b6a6-f1e9b797e1b0",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "202": {
                        "description": "Metadata not yet ready.",
                        "schema": {
                            "$ref": "#/definitions/rinser.Job"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "410": {
                        "description": "Job failed.",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    }
                }
            }
        },
        "/jobs/{uuid}/preview": {
            "get": {
                "description": "show job preview image by UUID",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/html",
                    "image/jpeg"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Show a job preview image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "49d1e304-d2b8-46bf-b6a6-f1e9b797e1b0",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "pages",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "172",
                        "name": "width",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "jpeg"
                        }
                    },
                    "202": {
                        "description": "Preview not yet ready.",
                        "schema": {
                            "$ref": "#/definitions/rinser.Job"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "410": {
                        "description": "Job failed.",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    }
                }
            }
        },
        "/jobs/{uuid}/rinsed": {
            "get": {
                "description": "Get the jobs rinsed document.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/pdf",
                    "application/json"
                ],
                "tags": [
                    "jobs"
                ],
                "summary": "Get the jobs rinsed document.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "49d1e304-d2b8-46bf-b6a6-f1e9b797e1b0",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "202": {
                        "description": "Rinsed version not yet ready.",
                        "schema": {
                            "$ref": "#/definitions/rinser.Job"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "410": {
                        "description": "Job failed.",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rinser.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "rinser.AddJobURL": {
            "type": "object",
            "properties": {
                "cleanupgotten": {
                    "type": "boolean",
                    "example": true
                },
                "cleanupsec": {
                    "type": "integer",
                    "example": 86400
                },
                "lang": {
                    "type": "string",
                    "example": "auto"
                },
                "maxsizemb": {
                    "type": "integer",
                    "example": 2048
                },
                "maxtimesec": {
                    "type": "integer",
                    "example": 3600
                },
                "url": {
                    "type": "string",
                    "example": "https://getsamplefiles.com/download/pdf/sample-1.pdf"
                }
            }
        },
        "rinser.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "rinser.Job": {
            "type": "object",
            "properties": {
                "cleanupgotten": {
                    "type": "boolean",
                    "example": true
                },
                "cleanupsec": {
                    "type": "integer",
                    "example": 600
                },
                "created": {
                    "type": "string",
                    "format": "dateTime",
                    "example": "2024-01-01T12:00:00+00:00"
                },
                "diskuse": {
                    "type": "integer",
                    "example": 1234
                },
                "done": {
                    "type": "boolean",
                    "example": false
                },
                "downloads": {
                    "description": "` + "`" + `json:\"downloads,omitempty\" example:\"0\"` + "`" + `",
                    "type": "integer"
                },
                "error": {},
                "lang": {
                    "type": "string",
                    "example": "auto"
                },
                "maxsizemb": {
                    "type": "integer",
                    "example": 2048
                },
                "maxtimesec": {
                    "type": "integer",
                    "example": 600
                },
                "name": {
                    "type": "string",
                    "example": "example.docx"
                },
                "pages": {
                    "type": "integer",
                    "example": 1
                },
                "pdfname": {
                    "description": "rinsed PDF file name",
                    "type": "string",
                    "example": "example-docx-rinsed.pdf"
                },
                "uuid": {
                    "type": "string",
                    "format": "uuid",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                },
                "workdir": {
                    "type": "string",
                    "example": "/tmp/rinse-550e8400-e29b-41d4-a716-446655440000"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "rinse REST API",
	Description:      "Document cleaning service API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
