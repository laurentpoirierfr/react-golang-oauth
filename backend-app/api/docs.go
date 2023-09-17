// Code generated by swaggo/swag. DO NOT EDIT.

package api

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/api/menus": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "Get versioni",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.MenuItem"
                            }
                        }
                    }
                }
            }
        },
        "/ops/info": {
            "get": {
                "description": "Informations sur le service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ops"
                ],
                "summary": "Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Info"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Info": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "domain.MenuItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "submenu": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.MenuItem"
                    }
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "homezone.swagger.io:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "bff",
	Description:      "This is a bff server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
