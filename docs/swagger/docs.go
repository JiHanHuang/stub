// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swagger

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
        "termsOfService": "https://github.com/JiHanHuang/stub",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/set/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Set"
                ],
                "summary": "自定义返回列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/set/response": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Set"
                ],
                "summary": "设置自定义返回",
                "parameters": [
                    {
                        "description": "设自定义返回结构",
                        "name": "setResponse",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/set.SetResponseForm"
                        }
                    },
                    {
                        "type": "string",
                        "default": "set_response",
                        "description": "自定义返回名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/tool/fingerprint": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool"
                ],
                "summary": "获取数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "appkey",
                        "name": "app_key",
                        "in": "query",
                        "required": true
                    },
                    {
                        "default": "{\"app_id\":\"xxxxx\",...}",
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/data": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "获取一定量的数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据量(k)默认0k",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/delay": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "延时返回",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "延时时长(默认5s)",
                        "name": "delay",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/download": {
            "get": {
                "tags": [
                    "Test"
                ],
                "summary": "下载文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file name",
                        "name": "filename",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.Context"
                        }
                    }
                }
            }
        },
        "/api/v1/download2": {
            "get": {
                "tags": [
                    "Test"
                ],
                "summary": "下载文件(不可靠)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file name",
                        "name": "filename",
                        "in": "query",
                        "required": true
                    }
                ]
            }
        },
        "/api/v1/get": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "获取数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "自定义返回(可选)",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/geturl": {
            "get": {
                "tags": [
                    "Test"
                ],
                "summary": "get url信息获取",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/pdata": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "获取一定量的数据(并发)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据量(0,1,10)默认0",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/post": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "上传数据",
                "parameters": [
                    {
                        "default": "{\"data\":\"helllo\"}",
                        "description": "post",
                        "name": "post",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "自定义返回(可选)",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/posturl": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "post url信息获取",
                "parameters": [
                    {
                        "default": "{\"data\":\"helllo\"}",
                        "description": "Data",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/show": {
            "get": {
                "tags": [
                    "Test"
                ],
                "summary": "get url信息获取",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/upload": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "gin.Context": {
            "type": "object",
            "properties": {
                "accepted": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "errors": {
                    "type": "errorMsgs"
                },
                "keys": {
                    "type": "object",
                    "additionalProperties": true
                },
                "params": {
                    "type": "Params"
                },
                "request": {
                    "type": "string"
                },
                "writer": {
                    "type": "ResponseWriter"
                }
            }
        },
        "set.SetResponseForm": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "contentType": {
                    "type": "string",
                    "example": "json"
                },
                "data": {
                    "type": "string",
                    "example": "your response data"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Gin-VUE API",
	Description: "An example of gin+vue",
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
