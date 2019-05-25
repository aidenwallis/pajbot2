// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-03-15 18:07:38.867873465 +0100 CET m=+0.100944248

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "API for pajbot2",
        "title": "pajbot2 API",
        "contact": {
            "name": "pajlada",
            "url": "https://pajlada.se",
            "email": "rasmus.karlsson@pajlada.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/pajbot/pajbot2/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "localhost:2355",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/channel/{channelID}/moderation/check_message": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Check a message in a bots filter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of channel to run the test in",
                        "name": "channelID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "message to test against the bots filters",
                        "name": "message",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/moderation.CheckMessageSuccessResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/utils.WebAPIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "moderation.CheckMessageSuccessResponse": {
            "type": "object",
            "properties": {
                "banned": {
                    "type": "boolean"
                },
                "filter_data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/moderation.filterData"
                    }
                },
                "input_message": {
                    "type": "string"
                }
            }
        },
        "moderation.filterData": {
            "type": "object",
            "properties": {
                "action_type": {
                    "type": "integer"
                },
                "reason": {
                    "type": "string"
                }
            }
        },
        "utils.WebAPIError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
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
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
