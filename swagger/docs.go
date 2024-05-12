// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "LocalAI",
            "url": "https://localai.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://raw.githubusercontent.com/mudler/LocalAI/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/assistants": {
            "post": {
                "summary": "Create an assistant with a model and instructions.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/openai.AssistantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/openai.Assistant"
                        }
                    }
                }
            }
        },
        "/v1/audio/speech": {
            "post": {
                "summary": "Generates audio from the input text.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.TTSRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/audio/transcriptions": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "summary": "Transcribes audio into the input language.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model",
                        "name": "model",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/v1/chat/completions": {
            "post": {
                "summary": "Generate a chat completions for a given prompt and model.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.OpenAIRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/schema.OpenAIResponse"
                        }
                    }
                }
            }
        },
        "/v1/completions": {
            "post": {
                "summary": "Generate completions for a given prompt and model.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.OpenAIRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/schema.OpenAIResponse"
                        }
                    }
                }
            }
        },
        "/v1/embeddings": {
            "post": {
                "summary": "Get a vector representation of a given input that can be easily consumed by machine learning models and algorithms.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.OpenAIRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/schema.OpenAIResponse"
                        }
                    }
                }
            }
        },
        "/v1/images/generations": {
            "post": {
                "summary": "Creates an image given a prompt.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.OpenAIRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/schema.OpenAIResponse"
                        }
                    }
                }
            }
        },
        "/v1/text-to-speech/{voice-id}": {
            "post": {
                "summary": "Generates audio from the input text.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "voice-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.TTSRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "functions.Argument": {
            "type": "object",
            "properties": {
                "properties": {
                    "type": "object",
                    "additionalProperties": true
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "functions.Function": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parameters": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "functions.FunctionName": {
            "type": "object",
            "properties": {
                "const": {
                    "type": "string"
                }
            }
        },
        "functions.FunctionProperties": {
            "type": "object",
            "properties": {
                "arguments": {
                    "$ref": "#/definitions/functions.Argument"
                },
                "function": {
                    "$ref": "#/definitions/functions.FunctionName"
                }
            }
        },
        "functions.ItemFunction": {
            "type": "object",
            "properties": {
                "properties": {
                    "$ref": "#/definitions/functions.FunctionProperties"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "functions.ItemName": {
            "type": "object",
            "properties": {
                "properties": {
                    "$ref": "#/definitions/functions.NameProperties"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "functions.JSONFunctionStructureFunction": {
            "type": "object",
            "properties": {
                "$defs": {
                    "type": "object",
                    "additionalProperties": true
                },
                "anyOf": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/functions.ItemFunction"
                    }
                },
                "oneOf": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/functions.ItemFunction"
                    }
                }
            }
        },
        "functions.JSONFunctionStructureName": {
            "type": "object",
            "properties": {
                "$defs": {
                    "type": "object",
                    "additionalProperties": true
                },
                "anyOf": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/functions.ItemName"
                    }
                },
                "oneOf": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/functions.ItemName"
                    }
                }
            }
        },
        "functions.NameProperties": {
            "type": "object",
            "properties": {
                "arguments": {
                    "$ref": "#/definitions/functions.Argument"
                },
                "name": {
                    "$ref": "#/definitions/functions.FunctionName"
                }
            }
        },
        "functions.Tool": {
            "type": "object",
            "properties": {
                "function": {
                    "$ref": "#/definitions/functions.Function"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "openai.Assistant": {
            "type": "object",
            "properties": {
                "created": {
                    "description": "The time at which the assistant was created.",
                    "type": "integer"
                },
                "description": {
                    "description": "The description of the assistant.",
                    "type": "string"
                },
                "file_ids": {
                    "description": "A list of file IDs attached to this assistant.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "description": "The unique identifier of the assistant.",
                    "type": "string"
                },
                "instructions": {
                    "description": "The system instructions that the assistant uses.",
                    "type": "string"
                },
                "metadata": {
                    "description": "Set of key-value pairs attached to the assistant.",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "model": {
                    "description": "The model ID used by the assistant.",
                    "type": "string"
                },
                "name": {
                    "description": "The name of the assistant.",
                    "type": "string"
                },
                "object": {
                    "description": "Object type, which is \"assistant\".",
                    "type": "string"
                },
                "tools": {
                    "description": "A list of tools enabled on the assistant.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/openai.Tool"
                    }
                }
            }
        },
        "openai.AssistantRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "file_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "instructions": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "model": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tools": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/openai.Tool"
                    }
                }
            }
        },
        "openai.Tool": {
            "type": "object",
            "properties": {
                "type": {
                    "$ref": "#/definitions/openai.ToolType"
                }
            }
        },
        "openai.ToolType": {
            "type": "string",
            "enum": [
                "code_interpreter",
                "retrieval",
                "function"
            ],
            "x-enum-varnames": [
                "CodeInterpreter",
                "Retrieval",
                "Function"
            ]
        },
        "schema.ChatCompletionResponseFormat": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                }
            }
        },
        "schema.Choice": {
            "type": "object",
            "properties": {
                "delta": {
                    "$ref": "#/definitions/schema.Message"
                },
                "finish_reason": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "message": {
                    "$ref": "#/definitions/schema.Message"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "schema.FunctionCall": {
            "type": "object",
            "properties": {
                "arguments": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.Item": {
            "type": "object",
            "properties": {
                "b64_json": {
                    "type": "string"
                },
                "embedding": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "index": {
                    "type": "integer"
                },
                "object": {
                    "type": "string"
                },
                "url": {
                    "description": "Images",
                    "type": "string"
                }
            }
        },
        "schema.Message": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "The message content"
                },
                "function_call": {
                    "description": "A result of a function call"
                },
                "name": {
                    "description": "The message name (used for tools calls)",
                    "type": "string"
                },
                "role": {
                    "description": "The message role",
                    "type": "string"
                },
                "string_content": {
                    "type": "string"
                },
                "string_images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "tool_calls": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.ToolCall"
                    }
                }
            }
        },
        "schema.OpenAIRequest": {
            "type": "object",
            "required": [
                "file"
            ],
            "properties": {
                "backend": {
                    "type": "string"
                },
                "batch": {
                    "description": "Custom parameters - not present in the OpenAI API",
                    "type": "integer"
                },
                "clip_skip": {
                    "description": "Diffusers",
                    "type": "integer"
                },
                "echo": {
                    "type": "boolean"
                },
                "file": {
                    "description": "whisper",
                    "type": "string"
                },
                "frequency_penalty": {
                    "type": "number"
                },
                "function_call": {
                    "description": "might be a string or an object"
                },
                "functions": {
                    "description": "A list of available functions to call",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/functions.Function"
                    }
                },
                "grammar": {
                    "description": "A grammar to constrain the LLM output",
                    "type": "string"
                },
                "grammar_json_functions": {
                    "$ref": "#/definitions/functions.JSONFunctionStructureFunction"
                },
                "grammar_json_name": {
                    "$ref": "#/definitions/functions.JSONFunctionStructureName"
                },
                "ignore_eos": {
                    "type": "boolean"
                },
                "input": {},
                "instruction": {
                    "description": "Edit endpoint",
                    "type": "string"
                },
                "language": {
                    "description": "Also part of the OpenAI official spec",
                    "type": "string"
                },
                "max_tokens": {
                    "type": "integer"
                },
                "messages": {
                    "description": "Messages is read only by chat/completion API calls",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Message"
                    }
                },
                "mode": {
                    "description": "Image (not supported by OpenAI)",
                    "type": "integer"
                },
                "model": {
                    "description": "Also part of the OpenAI official spec",
                    "type": "string"
                },
                "model_base_name": {
                    "description": "AutoGPTQ",
                    "type": "string"
                },
                "n": {
                    "description": "Also part of the OpenAI official spec. use it for returning multiple results",
                    "type": "integer"
                },
                "n_keep": {
                    "type": "integer"
                },
                "negative_prompt": {
                    "type": "string"
                },
                "negative_prompt_scale": {
                    "type": "number"
                },
                "presence_penalty": {
                    "type": "number"
                },
                "prompt": {
                    "description": "Prompt is read only by completion/image API calls"
                },
                "repeat_penalty": {
                    "type": "number"
                },
                "response_format": {
                    "description": "whisper/image",
                    "allOf": [
                        {
                            "$ref": "#/definitions/schema.ChatCompletionResponseFormat"
                        }
                    ]
                },
                "rope_freq_base": {
                    "type": "number"
                },
                "rope_freq_scale": {
                    "type": "number"
                },
                "seed": {
                    "type": "integer"
                },
                "size": {
                    "description": "image",
                    "type": "string"
                },
                "step": {
                    "type": "integer"
                },
                "stop": {},
                "stream": {
                    "type": "boolean"
                },
                "temperature": {
                    "type": "number"
                },
                "tfz": {
                    "type": "number"
                },
                "tokenizer": {
                    "description": "RWKV (?)",
                    "type": "string"
                },
                "tool_choice": {},
                "tools": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/functions.Tool"
                    }
                },
                "top_k": {
                    "type": "integer"
                },
                "top_p": {
                    "description": "Common options between all the API calls, part of the OpenAI spec",
                    "type": "number"
                },
                "typical_p": {
                    "type": "number"
                },
                "use_fast_tokenizer": {
                    "description": "AutoGPTQ",
                    "type": "boolean"
                }
            }
        },
        "schema.OpenAIResponse": {
            "type": "object",
            "properties": {
                "choices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Choice"
                    }
                },
                "created": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Item"
                    }
                },
                "id": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "object": {
                    "type": "string"
                },
                "usage": {
                    "$ref": "#/definitions/schema.OpenAIUsage"
                }
            }
        },
        "schema.OpenAIUsage": {
            "type": "object",
            "properties": {
                "completion_tokens": {
                    "type": "integer"
                },
                "prompt_tokens": {
                    "type": "integer"
                },
                "total_tokens": {
                    "type": "integer"
                }
            }
        },
        "schema.TTSRequest": {
            "type": "object",
            "properties": {
                "backend": {
                    "type": "string"
                },
                "input": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "voice": {
                    "type": "string"
                }
            }
        },
        "schema.ToolCall": {
            "type": "object",
            "properties": {
                "function": {
                    "$ref": "#/definitions/schema.FunctionCall"
                },
                "id": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "LocalAI API",
	Description:      "The LocalAI Rest API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
