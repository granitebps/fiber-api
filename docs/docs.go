// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Granite Bagas",
            "url": "https://granitebps.com",
            "email": "info@granitebps.com"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login existing user",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.LoginTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity"
                    }
                }
            }
        },
        "/api/v1/auth/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get authenticated user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.MeTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.RegisterTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity"
                    }
                }
            }
        },
        "/api/v1/blogs": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog"
                ],
                "summary": "Get all blogs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/transformer.BlogTransformer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog"
                ],
                "summary": "Save blog to database",
                "parameters": [
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateBlogRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.BlogTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity"
                    }
                }
            }
        },
        "/api/v1/blogs/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog"
                ],
                "summary": "Get single blog by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "blog id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.BlogTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog"
                ],
                "summary": "Update existing blog by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "blog id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "JSON payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateBlogRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/transformer.BlogTransformer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "422": {
                        "description": "Unprocessable Entity"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blog"
                ],
                "summary": "Delete existing blog by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "blog id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/api/v1/ping": {
            "get": {
                "description": "Do a ping to service just make sure service is working normally.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tools"
                ],
                "summary": "Ping to the service",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.JSONResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateBlogRequest": {
            "type": "object",
            "required": [
                "description",
                "image",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "request.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "request.UpdateBlogRequest": {
            "type": "object",
            "required": [
                "description",
                "image",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "transformer.BlogTransformer": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "transformer.LoginTransformer": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "transformer.MeTransformer": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "transformer.RegisterTransformer": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "utils.JSONResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "timestamp": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Fiber API",
	Description:      "This is a Fiber API Doc",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
