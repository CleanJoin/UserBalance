// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/CleanJoin/USERBALANCE/",
        "contact": {
            "name": "Github.com",
            "url": "https://github.com/CleanJoin/USERBALANCE/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/add": {
            "post": {
                "description": "Зачислить средства",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "addMoneyHandler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.RequestMoveMoney"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/health": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Info"
                ],
                "summary": "Show the status of server.",
                "responses": {}
            }
        },
        "/api/money": {
            "post": {
                "description": "Получить данные о балансе",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "getMoneyUserHadler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.RequestUser"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/reduce": {
            "post": {
                "description": "Списать денежные средства",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "reduceMoneyHandler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.RequestMoveMoney"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/transfer": {
            "post": {
                "description": "Перевести деньги от пользователя к пользователю",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "transferMoneyHandler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.RequestMoveMoney"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/user": {
            "post": {
                "description": "Регистрация пользователя",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "userHandler",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/balance.RequestUser"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "balance.RequestMoveMoney": {
            "type": "object",
            "properties": {
                "money": {
                    "type": "number"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "balance.RequestUser": {
            "type": "object",
            "properties": {
                "paswword": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger USERBALANCE",
	Description:      "This is a sample server USERBALANCE",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
