// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/action": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Создать новый action и привязаться к нему по токену\n- привязка по **id** или создание по остальным полям",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actions"
                ],
                "summary": "Создать новый action",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/actionModel.ActionCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/actions": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actions"
                ],
                "summary": "Найти action по неполному совпадению имени",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/actionModel.Action"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/typeInfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получить описание типов и вложеннымх в него сабтипов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TypeInfo"
                ],
                "summary": "Получить информацию по типам",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/typeInfoModel.TypeInfoResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/unitInfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TypeInfo"
                ],
                "summary": "Единицы измерения",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/unitModel.Unit"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}/actions": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получить userActions юзера (private=false) по id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Получить userActions юзера (private=false)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/userActionModel.UserActionsResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/userAction/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Удалить userAction по id по токену",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserActions"
                ],
                "summary": "Удалить userAction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/userAction/{id}/done": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "**action.type=1**; **required**: user_action_id *int*; **not required**: description *string*;\n**action.type=2**; **required**: user_action_id *int*; **not required**: description *string*;\n**action.type=3**; **required**: user_action_id *int*, count *float*; **not required**: description *string*;",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserActions"
                ],
                "summary": "Сделал action",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/logModel.CreateLogRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/userAction/{id}/private": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Редактировать userAction по id по токену",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserActions"
                ],
                "summary": "Редактировать userAction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userActionModel.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/userAction/{id}/stats": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получить статистику по id action по токену",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserActions"
                ],
                "summary": "Получить статистику",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date",
                        "name": "filter[log_date][gte]",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "date",
                        "name": "filter[log_date][lte]",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "field[eq]",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/logModel.GetStatsResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/userActions/my": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получить мои actions по токену",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserActions"
                ],
                "summary": "Получить мои actions",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "deleted",
                        "name": "deleted",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "field[eq]",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/userActionModel.MyActionsResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        },
        "/userActions/my/active": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получить мои активные actions с типом интервал по токену",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserActions"
                ],
                "summary": "Получить мои активные actions с типом интервал",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "deleted",
                        "name": "deleted",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "field[eq]",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userActionModel.MyActiveActions"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/defaultModel.FailedResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "actionModel.Action": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "subtype": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                },
                "unit": {
                    "type": "integer"
                },
                "unit_info": {
                    "$ref": "#/definitions/unitModel.Unit"
                }
            }
        },
        "actionModel.ActionCreateRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "subtype": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                },
                "unit": {
                    "type": "integer"
                }
            }
        },
        "defaultModel.FailedResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                }
            }
        },
        "defaultModel.SuccessResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                }
            }
        },
        "logModel.CreateLogRequest": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "start_start": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "logModel.GetStatsResponse": {
            "type": "object",
            "properties": {
                "action": {
                    "$ref": "#/definitions/actionModel.Action"
                },
                "count": {
                    "type": "integer"
                },
                "stats": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/logModel.Log"
                    }
                },
                "user_action_id": {
                    "type": "integer"
                }
            }
        },
        "logModel.Log": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "log_date": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                },
                "user_action_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "typeInfoModel.TypeInfoResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "sub_type": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "description": {
                                "type": "string"
                            },
                            "id": {
                                "type": "integer"
                            },
                            "name": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "unitModel.Unit": {
            "type": "object",
            "properties": {
                "description": {
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
        "userActionModel.MyActionsResponse": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "object",
                    "properties": {
                        "id": {
                            "type": "integer"
                        },
                        "name": {
                            "type": "string"
                        },
                        "subtype": {
                            "type": "integer"
                        },
                        "type": {
                            "type": "integer"
                        },
                        "unit_info": {
                            "type": "object",
                            "properties": {
                                "name": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                },
                "deleted": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "private": {
                    "type": "boolean"
                }
            }
        },
        "userActionModel.MyActiveActions": {
            "type": "object",
            "properties": {
                "action": {
                    "$ref": "#/definitions/actionModel.Action"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "string"
                }
            }
        },
        "userActionModel.UpdateRequest": {
            "type": "object",
            "properties": {
                "private": {
                    "type": "boolean"
                }
            }
        },
        "userActionModel.UserActionsResponse": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "object",
                    "properties": {
                        "id": {
                            "type": "integer"
                        },
                        "name": {
                            "type": "string"
                        },
                        "subtype": {
                            "type": "integer"
                        },
                        "type": {
                            "type": "integer"
                        },
                        "unit_info": {
                            "type": "object",
                            "properties": {
                                "name": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                },
                "id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Pood - just pood)",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
