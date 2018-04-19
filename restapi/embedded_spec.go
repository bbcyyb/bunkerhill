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
  "swagger": "2.0",
  "info": {
    "description": "RestAPI supporting lifecycle management for blog system named Missouri",
    "title": "Bunker Hill",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/apiversion": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiversion"
        ],
        "summary": "Return the bunkerhill version",
        "operationId": "GetAPIVersion",
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "$ref": "#/definitions/apiversion"
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      }
    },
    "/blogs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "GetBlogs",
        "parameters": [
          {
            "type": "string",
            "name": "author_id",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "page",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "pre_page",
            "in": "query"
          },
          {
            "type": "string",
            "description": "e.g. sortby=+title,-timestamp",
            "name": "sortby",
            "in": "query"
          },
          {
            "type": "string",
            "description": "e.g. select=title,id,body,body_html",
            "name": "select",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/blog"
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      },
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "InsertBlog",
        "parameters": [
          {
            "description": "Content of blog to be saved",
            "name": "blog",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/blog"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "The request has been fulfilled, resulting in the creation of a new resource.",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      }
    },
    "/blogs/{blogId}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "GetBlogById",
        "parameters": [
          {
            "type": "string",
            "description": "Id of the blog to return",
            "name": "blogId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/blog"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      },
      "put": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "UpdateBlog",
        "parameters": [
          {
            "type": "string",
            "description": "Id of the blog to update",
            "name": "blogId",
            "in": "path",
            "required": true
          },
          {
            "description": "Content of blog to be saved",
            "name": "blog",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/blog"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/blog"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "DeleteBlog",
        "parameters": [
          {
            "type": "string",
            "description": "Id of the blog to delete",
            "name": "blogId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "apiversion": {
      "type": "object",
      "required": [
        "apiVersion"
      ],
      "properties": {
        "apiVersion": {
          "type": "string"
        }
      }
    },
    "blog": {
      "type": "object",
      "properties": {
        "author": {
          "$ref": "#/definitions/user"
        },
        "body": {
          "type": "string"
        },
        "body_html": {
          "type": "string"
        },
        "comment_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "id": {
          "type": "string"
        },
        "timestamp": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "comment": {
      "type": "object",
      "properties": {
        "author": {
          "$ref": "#/definitions/user"
        },
        "body": {
          "type": "string"
        },
        "body_html": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "post_id": {
          "type": "integer",
          "format": "int32"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "generic.deleteResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "generic.error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Restful API Version",
      "name": "apiversion"
    },
    {
      "description": "Article information in blog",
      "name": "blog"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "RestAPI supporting lifecycle management for blog system named Missouri",
    "title": "Bunker Hill",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/apiversion": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "apiversion"
        ],
        "summary": "Return the bunkerhill version",
        "operationId": "GetAPIVersion",
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/getApiVersionOKBody"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      }
    },
    "/blogs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "GetBlogs",
        "parameters": [
          {
            "type": "string",
            "name": "author_id",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "page",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "pre_page",
            "in": "query"
          },
          {
            "type": "string",
            "description": "e.g. sortby=+title,-timestamp",
            "name": "sortby",
            "in": "query"
          },
          {
            "type": "string",
            "description": "e.g. select=title,id,body,body_html",
            "name": "select",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/blog"
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      },
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "InsertBlog",
        "parameters": [
          {
            "description": "Content of blog to be saved",
            "name": "blog",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/blog"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "The request has been fulfilled, resulting in the creation of a new resource.",
            "schema": {
              "$ref": "#/definitions/insertBlogCreatedBody"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      }
    },
    "/blogs/{blogId}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "GetBlogById",
        "parameters": [
          {
            "type": "string",
            "description": "Id of the blog to return",
            "name": "blogId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/blog"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      },
      "put": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "UpdateBlog",
        "parameters": [
          {
            "type": "string",
            "description": "Id of the blog to update",
            "name": "blogId",
            "in": "path",
            "required": true
          },
          {
            "description": "Content of blog to be saved",
            "name": "blog",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/blog"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "$ref": "#/definitions/blog"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "blog"
        ],
        "operationId": "DeleteBlog",
        "parameters": [
          {
            "type": "string",
            "description": "Id of the blog to delete",
            "name": "blogId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/generic.error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "apiversion": {
      "type": "object",
      "required": [
        "apiVersion"
      ],
      "properties": {
        "apiVersion": {
          "type": "string"
        }
      }
    },
    "blog": {
      "type": "object",
      "properties": {
        "author": {
          "$ref": "#/definitions/user"
        },
        "body": {
          "type": "string"
        },
        "body_html": {
          "type": "string"
        },
        "comment_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "id": {
          "type": "string"
        },
        "timestamp": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "comment": {
      "type": "object",
      "properties": {
        "author": {
          "$ref": "#/definitions/user"
        },
        "body": {
          "type": "string"
        },
        "body_html": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "post_id": {
          "type": "integer",
          "format": "int32"
        },
        "timestamp": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "generic.deleteResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "generic.error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "getApiVersionOKBody": {
      "type": "object",
      "properties": {
        "data": {
          "$ref": "#/definitions/apiversion"
        }
      },
      "x-go-gen-location": "operations"
    },
    "insertBlogCreatedBody": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    },
    "user": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Restful API Version",
      "name": "apiversion"
    },
    {
      "description": "Article information in blog",
      "name": "blog"
    }
  ]
}`))
}
