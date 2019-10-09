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
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Address Book",
    "version": "0.1.1"
  },
  "host": "localhost:8765",
  "paths": {
    "/contacts": {
      "get": {
        "summary": "List contacts",
        "responses": {
          "200": {
            "description": "Contacts list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Contact"
              }
            }
          },
          "default": {
            "$ref": "#/responses/Error"
          }
        }
      },
      "post": {
        "summary": "Add contact",
        "parameters": [
          {
            "name": "contact",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Contact"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Added contact",
            "schema": {
              "$ref": "#/definitions/Contact"
            }
          },
          "default": {
            "$ref": "#/responses/Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Contact": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int32",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600 with HTTP Status Code 422",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "Error": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "API-Key",
      "in": "header"
    }
  },
  "security": [
    {
      "api_key": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Address Book",
    "version": "0.1.1"
  },
  "host": "localhost:8765",
  "paths": {
    "/contacts": {
      "get": {
        "summary": "List contacts",
        "responses": {
          "200": {
            "description": "Contacts list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Contact"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "summary": "Add contact",
        "parameters": [
          {
            "name": "contact",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Contact"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Added contact",
            "schema": {
              "$ref": "#/definitions/Contact"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Contact": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int32",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600 with HTTP Status Code 422",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "Error": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "API-Key",
      "in": "header"
    }
  },
  "security": [
    {
      "api_key": []
    }
  ]
}`))
}
