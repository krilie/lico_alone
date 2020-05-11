// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-11 21:12:13.1803729 +0800 CST m=+0.295052301

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
        "/api/common/icp_info": {
            "get": {
                "description": "Icp信息",
                "tags": [
                    "Common"
                ],
                "summary": "Icp信息",
                "operationId": "Icp信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.IcpInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
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
        "/api/user/login": {
            "post": {
                "description": "用户用密码登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
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
                        "description": "2000 {token:\"asb\"}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/com_model.CommonReturn"
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
        "/api/user/register": {
            "post": {
                "description": "用户注册",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
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
                            "$ref": "#/definitions/com_model.CommonReturn"
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
        "/api/user/send_sms": {
            "post": {
                "description": "用户发短信",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
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
                            "$ref": "#/definitions/com_model.CommonReturn"
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
        "/health": {
            "get": {
                "description": "健康检查2",
                "tags": [
                    "基本信息"
                ],
                "summary": "健康检查2",
                "operationId": "健康检查2",
                "responses": {
                    "200": {
                        "description": "pong start time up time",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "Version",
                "tags": [
                    "基本信息"
                ],
                "summary": "Version",
                "operationId": "Version",
                "responses": {
                    "200": {
                        "description": "version build_time git_commit go_version",
                        "schema": {
                            "type": "string"
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
        }
    },
    "definitions": {
        "com_model.CommonReturn": {
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
        },
        "model.IcpInfo": {
            "type": "object",
            "properties": {
                "label": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "name": {
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
