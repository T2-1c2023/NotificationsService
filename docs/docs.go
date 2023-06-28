// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/": {
            "get": {
                "description": "Returns a 200 code and JSON with a message.",
                "summary": "Check the service's status.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Returns a 200 code and JSON with the date the service started running and a description.",
                "summary": "Get the service's health status.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.HealthResponse"
                        }
                    }
                }
            }
        },
        "/logs": {
            "get": {
                "description": "Get the service's logs.",
                "produces": [
                    "text/plain"
                ],
                "summary": "Get the service's logs.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/new-follower": {
            "post": {
                "description": "Send new follower notification to the given followed user.",
                "summary": "Send new follower notification to the given followed user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Follower user decoded payload info",
                        "name": "user_info",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Followed user ID",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.NewFollowerInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/new-message": {
            "post": {
                "description": "Send new message notification to the given recipient user.",
                "summary": "Send new message notification to the given recipient user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Sender user decoded payload info",
                        "name": "user_info",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Recipient user ID and message",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.NewMessageInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "Get the status of the service.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get the service's blocked status.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Decoded payload of the admin token",
                        "name": "user_info",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ChangeStatusInput"
                        }
                    }
                }
            },
            "post": {
                "description": "Changes the status of the service.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Change the service's blocked status.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Decoded payload of the admin token",
                        "name": "user_info",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Blocked status of the service",
                        "name": "rule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ChangeStatusInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ChangeStatusInput"
                        }
                    }
                }
            }
        },
        "/training-completed": {
            "post": {
                "description": "Send completed training notification to the user.",
                "summary": "Send completed training notification to the user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User decoded payload info",
                        "name": "user_info",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Completed training title",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.TrainingCompletedInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ChangeStatusInput": {
            "type": "object",
            "properties": {
                "blocked": {
                    "type": "boolean"
                }
            }
        },
        "controller.HealthResponse": {
            "type": "object",
            "properties": {
                "creation_date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "last_response": {
                    "type": "string"
                }
            }
        },
        "controller.NewFollowerInput": {
            "type": "object",
            "required": [
                "followed_id"
            ],
            "properties": {
                "followed_id": {
                    "type": "integer"
                }
            }
        },
        "controller.NewMessageInput": {
            "type": "object",
            "required": [
                "message",
                "recipient_id"
            ],
            "properties": {
                "message": {
                    "type": "string"
                },
                "recipient_id": {
                    "type": "integer"
                }
            }
        },
        "controller.TrainingCompletedInput": {
            "type": "object",
            "required": [
                "training_title"
            ],
            "properties": {
                "training_title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Notifications Microservice API",
	Description:      "Notifications microservice for FiuFit's backend.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
