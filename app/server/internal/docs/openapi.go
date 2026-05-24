package docs

import "net/http"

const openAPISpec = `{
  "openapi": "3.1.0",
  "info": {
    "title": "Vow API",
    "version": "0.1.0",
    "description": "HTTP API for Vow."
  },
  "servers": [
    {
      "url": "http://localhost:8080",
      "description": "Local development"
    }
  ],
  "paths": {
    "/health": {
      "get": {
        "summary": "Health check",
        "operationId": "getHealth",
        "tags": ["System"],
        "responses": {
          "200": {
            "description": "Server is healthy",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HealthResponse"
                },
                "examples": {
                  "ok": {
                    "value": {
                      "status": "ok"
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/v1/api/auth/login": {
      "post": {
        "summary": "Log in",
        "operationId": "login",
        "tags": ["Auth"],
        "responses": {
          "501": {
            "description": "Login is not implemented yet",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/MessageResponse"
                },
                "examples": {
                  "notImplemented": {
                    "value": {
                      "message": "login not implemented"
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/v1/api/auth/register": {
      "post": {
        "summary": "Register",
        "operationId": "register",
        "tags": ["Auth"],
        "responses": {
          "501": {
            "description": "Registration is not implemented yet",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/MessageResponse"
                },
                "examples": {
                  "notImplemented": {
                    "value": {
                      "message": "register not implemented"
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/v1/api/auth/logout": {
      "post": {
        "summary": "Log out",
        "operationId": "logout",
        "tags": ["Auth"],
        "responses": {
          "501": {
            "description": "Logout is not implemented yet",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/MessageResponse"
                },
                "examples": {
                  "notImplemented": {
                    "value": {
                      "message": "logout not implemented"
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "HealthResponse": {
        "type": "object",
        "required": ["status"],
        "properties": {
          "status": {
            "type": "string",
            "example": "ok"
          }
        }
      },
      "MessageResponse": {
        "type": "object",
        "required": ["message"],
        "properties": {
          "message": {
            "type": "string"
          }
        }
      }
    }
  }
}`

func OpenAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(openAPISpec))
}
