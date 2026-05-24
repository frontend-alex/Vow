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
                  "$ref": "#/components/schemas/APIResponse"
                },
                "examples": {
                  "ok": {
                    "value": {
                      "success": true,
                      "message": "Server is running and healthy",
                      "error_message": null,
                      "error_status": null,
                      "data": null
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
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/LoginRequest"
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "Logged in successfully",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/AuthAPIResponse"
                }
              }
            }
          },
          "400": {
            "$ref": "#/components/responses/BadRequest"
          },
          "401": {
            "$ref": "#/components/responses/Unauthorized"
          }
        }
      }
    },
    "/v1/api/auth/register": {
      "post": {
        "summary": "Register",
        "operationId": "register",
        "tags": ["Auth"],
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/RegisterRequest"
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "Registered successfully",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/AuthAPIResponse"
                }
              }
            }
          },
          "400": {
            "$ref": "#/components/responses/BadRequest"
          },
          "409": {
            "$ref": "#/components/responses/Conflict"
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
                  "$ref": "#/components/schemas/APIResponse"
                },
                "examples": {
                  "notImplemented": {
                    "value": {
                      "success": false,
                      "message": null,
                      "error_message": "logout not implemented",
                      "error_status": 501,
                      "data": null
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/v1/api/onboarding/start": {
      "post": {
        "summary": "Start onboarding",
        "operationId": "startOnboarding",
        "tags": ["Onboarding"],
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "responses": {
          "200": {
            "description": "Onboarding started",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/OnboardingAPIResponse"
                }
              }
            }
          },
          "401": {
            "$ref": "#/components/responses/Unauthorized"
          },
          "500": {
            "$ref": "#/components/responses/InternalServerError"
          }
        }
      }
    },
    "/v1/api/onboarding/complete": {
      "post": {
        "summary": "Complete onboarding",
        "operationId": "completeOnboarding",
        "tags": ["Onboarding"],
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/CompleteOnboardingRequest"
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "Onboarding completed",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/APIResponse"
                },
                "examples": {
                  "completed": {
                    "value": {
                      "success": true,
                      "message": "onboarding completed",
                      "error_message": null,
                      "error_status": null,
                      "data": null
                    }
                  }
                }
              }
            }
          },
          "400": {
            "$ref": "#/components/responses/BadRequest"
          },
          "401": {
            "$ref": "#/components/responses/Unauthorized"
          },
          "500": {
            "$ref": "#/components/responses/InternalServerError"
          }
        }
      }
    }
  },
  "components": {
    "securitySchemes": {
      "bearerAuth": {
        "type": "http",
        "scheme": "bearer",
        "bearerFormat": "JWT"
      }
    },
    "responses": {
      "BadRequest": {
        "description": "Bad request",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/APIResponse"
            }
          }
        }
      },
      "Unauthorized": {
        "description": "Missing, invalid, or expired credentials",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/APIResponse"
            }
          }
        }
      },
      "Conflict": {
        "description": "Conflict",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/APIResponse"
            }
          }
        }
      },
      "InternalServerError": {
        "description": "Internal server error",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/APIResponse"
            }
          }
        }
      }
    },
    "schemas": {
      "APIResponse": {
        "type": "object",
        "required": ["success", "message", "error_message", "error_status", "data"],
        "properties": {
          "success": {
            "type": "boolean"
          },
          "message": {
            "type": ["string", "null"]
          },
          "error_message": {
            "type": ["string", "null"]
          },
          "error_status": {
            "type": ["integer", "null"]
          },
          "data": true
        }
      },
      "LoginRequest": {
        "type": "object",
        "required": ["email", "password"],
        "properties": {
          "email": {
            "type": "string",
            "format": "email"
          },
          "password": {
            "type": "string",
            "format": "password"
          }
        }
      },
      "RegisterRequest": {
        "type": "object",
        "required": ["email", "name", "password"],
        "properties": {
          "email": {
            "type": "string",
            "format": "email"
          },
          "name": {
            "type": "string"
          },
          "password": {
            "type": "string",
            "format": "password"
          }
        }
      },
      "AuthUser": {
        "type": "object",
        "required": ["id", "email", "name", "onboardingCompleted"],
        "properties": {
          "id": {
            "type": "integer",
            "format": "int64"
          },
          "email": {
            "type": "string",
            "format": "email"
          },
          "name": {
            "type": "string"
          },
          "onboardingCompleted": {
            "type": "boolean"
          }
        }
      },
      "AuthResponse": {
        "type": "object",
        "required": ["accessToken", "tokenType", "user"],
        "properties": {
          "accessToken": {
            "type": "string"
          },
          "tokenType": {
            "type": "string",
            "example": "Bearer"
          },
          "user": {
            "$ref": "#/components/schemas/AuthUser"
          }
        }
      },
      "AuthAPIResponse": {
        "allOf": [
          {
            "$ref": "#/components/schemas/APIResponse"
          },
          {
            "type": "object",
            "properties": {
              "data": {
                "$ref": "#/components/schemas/AuthResponse"
              }
            }
          }
        ]
      },
      "UserOnboarding": {
        "type": "object",
        "required": ["ID", "UserID", "Status", "Version", "StartedAt", "CreatedAt", "UpdatedAt"],
        "properties": {
          "ID": {
            "type": "integer",
            "format": "int64"
          },
          "UserID": {
            "type": "integer",
            "format": "int64"
          },
          "Status": {
            "type": "string",
            "enum": ["in_progress", "completed", "abandoned"]
          },
          "Version": {
            "type": "integer"
          },
          "CurrentStep": {
            "type": ["string", "null"]
          },
          "StartedAt": {
            "type": "string",
            "format": "date-time"
          },
          "CompletedAt": {
            "type": ["string", "null"],
            "format": "date-time"
          },
          "CreatedAt": {
            "type": "string",
            "format": "date-time"
          },
          "UpdatedAt": {
            "type": "string",
            "format": "date-time"
          }
        }
      },
      "OnboardingAPIResponse": {
        "allOf": [
          {
            "$ref": "#/components/schemas/APIResponse"
          },
          {
            "type": "object",
            "properties": {
              "data": {
                "$ref": "#/components/schemas/UserOnboarding"
              }
            }
          }
        ]
      },
      "CompleteOnboardingRequest": {
        "type": "object",
        "required": [
          "improvementGoals",
          "riskMoment",
          "blockedApps",
          "wakeTime",
          "repeatDays",
          "unlockTasks",
          "difficulty",
          "priority1",
          "priority2",
          "priority3",
          "morningIntention",
          "desiredMorningState",
          "autoUnlockEnabled",
          "autoUnlockAfterMinutes"
        ],
        "properties": {
          "improvementGoals": {
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          "riskMoment": {
            "type": "string"
          },
          "blockedApps": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/BlockedAppInput"
            }
          },
          "wakeTime": {
            "type": "string",
            "example": "07:00"
          },
          "repeatDays": {
            "type": "array",
            "items": {
              "type": "string",
              "example": "monday"
            }
          },
          "unlockTasks": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/UnlockTaskInput"
            }
          },
          "difficulty": {
            "type": "string",
            "enum": ["gentle", "balanced", "strong"]
          },
          "priority1": {
            "type": "string"
          },
          "priority2": {
            "type": "string"
          },
          "priority3": {
            "type": "string"
          },
          "morningIntention": {
            "type": "string"
          },
          "desiredMorningState": {
            "type": "string"
          },
          "autoUnlockEnabled": {
            "type": "boolean"
          },
          "autoUnlockAfterMinutes": {
            "type": "integer"
          }
        }
      },
      "BlockedAppInput": {
        "type": "object",
        "required": ["appIdentifier", "appName", "appCategory", "platform"],
        "properties": {
          "appIdentifier": {
            "type": "string"
          },
          "appName": {
            "type": "string"
          },
          "appCategory": {
            "type": "string"
          },
          "platform": {
            "type": "string",
            "enum": ["ios", "android", "web", "desktop"]
          }
        }
      },
      "UnlockTaskInput": {
        "type": "object",
        "required": ["taskType", "title", "description", "metadata"],
        "properties": {
          "taskType": {
            "type": "string",
            "enum": ["walk_steps", "drink_water", "stretch", "scan_qr", "write_intention", "breathing", "custom"]
          },
          "title": {
            "type": "string"
          },
          "description": {
            "type": "string"
          },
          "requiredValue": {
            "type": ["integer", "null"]
          },
          "metadata": {
            "type": "object",
            "additionalProperties": true
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
