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
            "name": "Connor Davis",
            "url": "https://thusa.co.za",
            "email": "connor.davis@thusa.co.za"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/authentication/check": {
            "get": {
                "description": "Check if user is authenticated",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Check",
                "responses": {
                    "200": {
                        "description": "User object",
                        "schema": {
                            "$ref": "#/definitions/postgres.User"
                        }
                    }
                }
            }
        },
        "/authentication/login": {
            "post": {
                "description": "Authenticate a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Authenticate User",
                "parameters": [
                    {
                        "description": "Login payload.",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authentication.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Authenticated.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request body.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authentication/mfa/disable": {
            "patch": {
                "description": "Disable MFA for the current user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Disable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid query parameters.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "The user was not found.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authentication/mfa/enable": {
            "get": {
                "description": "Enable MFA for the current user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Enable",
                "responses": {
                    "200": {
                        "description": "QR Code",
                        "schema": {
                            "type": "file"
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
        "/authentication/mfa/verify": {
            "post": {
                "description": "Verify MFA for the current user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Verify",
                "parameters": [
                    {
                        "description": "Verify MFA Payload.",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authentication.VerifyMfaPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid MFA code.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authentication.Login": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "authentication.VerifyMfaPayload": {
            "type": "object",
            "required": [
                "code"
            ],
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "postgres.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "role"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "mfaEnabled": {
                    "type": "boolean"
                },
                "mfaVerified": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:6173",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "ZingFibre Reports Portal API",
	Description:      "This is the ZingFibre Reports Portal API built with Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
