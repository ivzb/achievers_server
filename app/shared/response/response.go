package response

import "net/http"

// Message is the return type of all handlers
type Message struct {
	StatusCode int
	Result     interface{}
}

// Core Response
type Core struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Change Response
type Change struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Affected int    `json:"affected"`
}

// Retrieve Response
type Retrieve struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Length  int         `json:"length"`
	Results interface{} `json:"results"`
}

// Ok sends response with status code 200
func Ok(
	message string,
	length int,
	results interface{}) Message {

	return send(http.StatusOK, message, length, results)
}

// Created sends response with status code 201
func Created(
	message string,
	results interface{}) Message {

	return send(http.StatusCreated, message, 1, results)
}

// BadRequest sends response with status code 400
func BadRequest(message string) Message {
	return sendError(http.StatusBadRequest, message)
}

// Unauthorized sends response with status code 401
func Unauthorized(message string) Message {
	return sendError(http.StatusUnauthorized, message)
}

// NotFound sends response with status code 404
func NotFound(message string) Message {
	return sendError(http.StatusNotFound, message)
}

// MethodNotAllowed sends response with status code 405
func MethodNotAllowed(message string) Message {
	return sendError(http.StatusMethodNotAllowed, message)
}

// InternalServerError sends response with status code 500
func InternalServerError(message string) Message {
	return sendError(http.StatusInternalServerError, message)
}

// sendError calls Send by without a count or results
func sendError(status int, message string) Message {
	return send(status, message, 0, nil)
}

// send writes struct to the writer using a format
func send(
	status int,
	message string,
	length int,
	results interface{}) Message {

	var result interface{}

	// Determine the best format
	if length < 1 {
		result = &Core{
			Status:  status,
			Message: message,
		}
	} else if results == nil {
		result = &Change{
			Status:   status,
			Message:  message,
			Affected: length,
		}
	} else {
		result = &Retrieve{
			Status:  status,
			Message: message,
			Length:  length,
			Results: results,
		}
	}

	return Message{status, result}
}
