package response

import (
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/shared/consts"
)

const (
	TypeJSON = 0
	TypeFile = 1
)

// Message is the return type of all handlers
type Message struct {
	StatusCode int
	Result     interface{}
	Type       int
}

// Core Response
type Core struct {
	Message string `json:"message"`
}

// Change Response
type Change struct {
	Message  string `json:"message"`
	Affected int    `json:"affected"`
}

// Retrieve Response
type Retrieve struct {
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

// Ok sends response with status code 200
func Ok(
	key string,
	length int,
	results interface{}) *Message {

	message := fmt.Sprintf(consts.FormatFound, key)
	return send(http.StatusOK, message, length, results)
}

// File sends response with status code 200
func File(results string) *Message {
	result := &Core{
		Message: results,
	}

	return &Message{http.StatusOK, result, TypeFile}
}

// Created sends response with status code 201
func Created(
	key string,
	results interface{}) *Message {

	message := fmt.Sprintf(consts.FormatCreated, key)
	return send(http.StatusCreated, message, 1, results)
}

// BadRequest sends response with status code 400
func BadRequest(message string) *Message {
	return sendError(http.StatusBadRequest, message)
}

// Unauthorized sends response with status code 401
func Unauthorized(message string) *Message {
	return sendError(http.StatusUnauthorized, message)
}

// NotFound sends response with status code 404
func NotFound(key string) *Message {
	message := fmt.Sprintf(consts.FormatNotFound, key)
	return sendError(http.StatusNotFound, message)
}

// MethodNotAllowed sends response with status code 405
func MethodNotAllowed() *Message {
	return sendError(http.StatusMethodNotAllowed, consts.MethodNotAllowed)
}

// InternalServerError sends response with status code 500
func InternalServerError() *Message {
	return sendError(http.StatusInternalServerError, consts.FriendlyErrorMessage)
}

// sendError calls Send by without a count or results
func sendError(status int, message string) *Message {
	return send(status, message, 0, nil)
}

// send writes struct to the writer using a format
func send(
	status int,
	message string,
	length int,
	results interface{}) *Message {

	var result interface{}

	// Determine the best format
	if length < 1 {
		result = &Core{
			Message: message,
		}
	} else if results == nil {
		result = &Change{
			Message:  message,
			Affected: length,
		}
	} else {
		result = &Retrieve{
			Message: message,
			Results: results,
		}
	}

	return &Message{status, result, TypeJSON}
}
