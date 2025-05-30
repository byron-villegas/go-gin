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
            "name": "Byron Villegas Moya",
            "url": "https://github.com/byron-villegas",
            "email": "byronvillegasm@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/byron-villegas/go-gin/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "description": "Get all products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    }
                }
            }
        },
        "/products/{sku}": {
            "get": {
                "description": "Get product by SKU",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get product by SKU",
                "parameters": [
                    {
                        "type": "string",
                        "description": "SKU",
                        "name": "sku",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Characteristic": {
            "type": "object",
            "properties": {
                "titulo": {
                    "type": "string"
                },
                "valor": {
                    "type": "string"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "caracteristicas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Characteristic"
                    }
                },
                "descripcion": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imagen": {
                    "type": "string"
                },
                "marca": {
                    "type": "string"
                },
                "nombre": {
                    "type": "string"
                },
                "precio": {
                    "type": "integer"
                },
                "sku": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Go Gin",
	Description:      "Proyecto base para aplicaciones Gin con ejemplos de configuración, testing y buenas prácticas.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
