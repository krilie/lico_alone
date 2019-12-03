// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-30 10:02:08.0208152 +0800 CST m=+0.274425301

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "Swagger Example API",
        "contact": {},
        "license": {},
        "version": "0.0.1"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/user/login": {
            "post": {
                "description": "用户用密码登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户账号"
                ],
                "summary": "用户登录",
                "operationId": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户手机号",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{jwt:\"asb\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    }
                }
            }
        },
        "/v1/user/register": {
            "post": {
                "description": "用户注册",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户账号"
                ],
                "summary": "用户注册",
                "operationId": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户手机号",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码",
                        "name": "valid_code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    }
                }
            }
        },
        "/v1/user/send_sms": {
            "post": {
                "description": "用户发短信",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户发短信"
                ],
                "summary": "用户发短信",
                "operationId": "用户发短信",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户手机号",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "register login change_password",
                        "name": "send_type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/cmodel.CommonReturn"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cmodel.CommonReturn": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "data": {
                    "description": "数据值",
                    "type": "object"
                },
                "detail": {
                    "description": "可由运行模式控制是否显示",
                    "type": "string",
                    "example": "错误的详细信息，用于排查错误"
                },
                "message": {
                    "type": "string",
                    "example": "错误信息"
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
var SwaggerInfo = swaggerInfo{Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
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
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
