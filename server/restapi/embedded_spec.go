// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json",
    "text/html"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "blackbox",
    "title": "blackbox",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/user/{id}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get user account\n",
        "tags": [
          "users"
        ],
        "summary": "Get user",
        "parameters": [
          {
            "type": "string",
            "description": "The unique user reference id\n",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully got user",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "description": "A human-readable error message",
          "type": "string",
          "example": "You are not authorized to access this resource."
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "name": {
          "description": "The user's name",
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json",
    "text/html"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "blackbox",
    "title": "blackbox",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/user/{id}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get user account\n",
        "tags": [
          "users"
        ],
        "summary": "Get user",
        "parameters": [
          {
            "type": "string",
            "description": "The unique user reference id\n",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully got user",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "description": "A human-readable error message",
          "type": "string",
          "example": "You are not authorized to access this resource."
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "name": {
          "description": "The user's name",
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
