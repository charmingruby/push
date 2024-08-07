// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Gustavo Dias",
            "url": "https://github.com/charmingruby",
            "email": "gustavodiasa2121@gmail.com"
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
        "/communication-channels": {
            "post": {
                "description": "Creates a communication channel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Communication Channel"
                ],
                "summary": "Creates a communication channel",
                "parameters": [
                    {
                        "description": "Create Communication Channel Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.CreateCommunicationChannelRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/notifications": {
            "post": {
                "description": "Schedules a notification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notifications"
                ],
                "summary": "Schedules a notification",
                "parameters": [
                    {
                        "description": "Schedule Notification Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.ScheduleNotificationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/notifications/{id}": {
            "get": {
                "description": "Gets a notification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notifications"
                ],
                "summary": "Gets a notification",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get Notification Payload",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.GetNotificationResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/notifications/{id}/cancel": {
            "patch": {
                "description": "Cancel notification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notifications"
                ],
                "summary": "Cancel notification",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cancel Notification Payload",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CancelNotificationResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        },
        "/welcome": {
            "get": {
                "description": "Health Check",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "notification_entity.Notification": {
            "type": "object",
            "required": [
                "communication_channel_id",
                "created_at",
                "date",
                "destination",
                "id",
                "status"
            ],
            "properties": {
                "communication_channel_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "retries": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "rest.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "v1.CancelNotificationResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/notification_entity.Notification"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.CreateCommunicationChannelRequest": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "v1.GetNotificationResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/notification_entity.Notification"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.ScheduleNotificationRequest": {
            "type": "object",
            "required": [
                "communication_channel_id",
                "destination",
                "raw_date"
            ],
            "properties": {
                "communication_channel_id": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "raw_date": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "push",
	Description:      "This is the push for new Go APIs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
