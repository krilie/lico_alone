// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "/api/common/article/query_sample": {
            "get": {
                "description": "查询文章列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共接口"
                ],
                "summary": "查询文章列表",
                "operationId": "查询文章列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "搜索内容",
                        "name": "search_key",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页索引",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/com_model.CommonReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/com_model.PageData"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "Data": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/model.QueryArticleModel"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
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
        "/api/common/icp_info": {
            "get": {
                "description": "Icp信息",
                "tags": [
                    "公共接口"
                ],
                "summary": "Icp信息",
                "operationId": "Icp信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
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
        "/api/manage/article/get_by_id": {
            "get": {
                "description": "获取一个文章信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "通过id获取文章",
                "operationId": "获取一个文章信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
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
        "/api/manage/article/update": {
            "post": {
                "description": "更新文章内容",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "更新文章内容",
                "operationId": "更新文章内容",
                "parameters": [
                    {
                        "description": "文章内容",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateArticleModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
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
        "/api/manage/setting/get_setting_all": {
            "post": {
                "description": "获取所有配置项",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "配置项"
                ],
                "summary": "获取所有配置项",
                "operationId": "获取所有配置项",
                "parameters": [
                    {
                        "type": "string",
                        "description": "搜索值",
                        "name": "searchKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/com_model.CommonReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Config"
                                            }
                                        }
                                    }
                                }
                            ]
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
        "/api/manage/setting/update_config": {
            "post": {
                "description": "更新配置项",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "配置项"
                ],
                "summary": "更新配置项",
                "operationId": "更新配置项",
                "parameters": [
                    {
                        "type": "string",
                        "description": "配置名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "配置值",
                        "name": "value",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
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
                "description": "健康检查Hello 返回hello字样",
                "tags": [
                    "健康检查"
                ],
                "summary": "健康检查Hello",
                "operationId": "健康检查Hello",
                "responses": {
                    "200": {
                        "description": "hello",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health/ping": {
            "get": {
                "description": "健康检查Ping 检查数据库是否正常 并返回启动时间",
                "tags": [
                    "健康检查"
                ],
                "summary": "健康检查Ping",
                "operationId": "健康检查Ping",
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
                    "公共接口"
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
        "com_model.PageData": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "列表数据",
                    "type": "object"
                },
                "page_info": {
                    "description": "分页信息",
                    "type": "object",
                    "$ref": "#/definitions/com_model.PageInfo"
                }
            }
        },
        "com_model.PageInfo": {
            "type": "object",
            "properties": {
                "page_num": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total": {
                    "description": "总数据条数",
                    "type": "integer"
                },
                "total_page": {
                    "type": "integer"
                }
            }
        },
        "model.Config": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
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
        },
        "model.QueryArticleModel": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "pv": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.UpdateArticleModel": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "title": {
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
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
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
