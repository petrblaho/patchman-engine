// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-25 14:57:38.382663372 +0100 CET m=+0.052983360

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/patch/v1/advisories": {
            "get": {
                "description": "Show me all applicable errata for all my systems",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show me all applicable errata for all my systems",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SystemDetailResponse"
                        }
                    }
                }
            }
        },
        "/api/patch/v1/advisories/{advisory_id}": {
            "get": {
                "description": "Show me details an advisory by given advisory name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show me details an advisory by given advisory name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Advisory ID",
                        "name": "advisory_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.AdvisoryDetailResponse"
                        }
                    }
                }
            }
        },
        "/api/patch/v1/advisories/{advisory_id}/applicable_systems": {
            "get": {
                "description": "Show me all my systems",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show me all my systems",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Advisory ID",
                        "name": "advisory_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SystemsApplicableResponse"
                        }
                    }
                }
            }
        },
        "/api/patch/v1/systems": {
            "get": {
                "description": "Show me all my systems",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show me all my systems",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SystemDetailResponse"
                        }
                    }
                }
            }
        },
        "/api/patch/v1/systems/{inventory_id}": {
            "get": {
                "description": "Show me details about a system by given inventory id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show me details about a system by given inventory id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Inventory ID",
                        "name": "inventory_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SystemDetailResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.AdvisoryDetailAttributes": {
            "type": "object",
            "properties": {
                "cves": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "fixes": {
                    "type": "string"
                },
                "modified_date": {
                    "type": "string"
                },
                "public_date": {
                    "type": "string"
                },
                "references": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "severity": {
                    "type": "string"
                },
                "solution": {
                    "type": "string"
                },
                "synopsis": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                }
            }
        },
        "controllers.AdvisoryDetailResponse": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.AdvisoryDetailAttributes"
                },
                "id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "controllers.AdvisoryItem": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.AdvisoryItemAttributes"
                },
                "id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "controllers.AdvisoryItemAttributes": {
            "type": "object",
            "properties": {
                "advisory_type": {
                    "type": "integer"
                },
                "applicable_systems": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "public_date": {
                    "type": "string"
                },
                "severity": {
                    "type": "string"
                },
                "synopsis": {
                    "type": "string"
                }
            }
        },
        "controllers.AdvisoryMeta": {
            "type": "object",
            "properties": {
                "data_format": {
                    "type": "string"
                },
                "filter": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "pages": {
                    "type": "integer"
                },
                "public_from": {
                    "type": "integer"
                },
                "public_to": {
                    "type": "integer"
                },
                "severity": {
                    "type": "string"
                },
                "show_all": {
                    "type": "boolean"
                },
                "sort": {
                    "type": "boolean"
                },
                "total_items": {
                    "type": "integer"
                }
            }
        },
        "controllers.Links": {
            "type": "object",
            "properties": {
                "first": {
                    "type": "string"
                },
                "last": {
                    "type": "string"
                },
                "next": {
                    "type": "string"
                },
                "previous": {
                    "type": "string"
                }
            }
        },
        "controllers.SystemDetailResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "advisories items",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.AdvisoryItem"
                    }
                },
                "links": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.Links"
                },
                "meta": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.AdvisoryMeta"
                }
            }
        },
        "controllers.SystemItem": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.SystemItemAttributes"
                },
                "id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "controllers.SystemItemAttributes": {
            "type": "object",
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "last_evaluation": {
                    "type": "string"
                },
                "last_upload": {
                    "type": "string"
                },
                "rhba_count": {
                    "type": "integer"
                },
                "rhea_count": {
                    "type": "integer"
                },
                "rhsa_count": {
                    "type": "integer"
                }
            }
        },
        "controllers.SystemsApplicableMeta": {
            "type": "object",
            "properties": {
                "advisory": {
                    "type": "string"
                },
                "data_format": {
                    "type": "string"
                },
                "enabled": {
                    "type": "boolean"
                },
                "filter": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "pages": {
                    "type": "integer"
                },
                "total_items": {
                    "type": "integer"
                }
            }
        },
        "controllers.SystemsApplicableResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.SystemItem"
                    }
                },
                "links": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.Links"
                },
                "meta": {
                    "type": "object",
                    "$ref": "#/definitions/controllers.SystemsApplicableMeta"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
