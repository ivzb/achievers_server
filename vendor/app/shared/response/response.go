package response

import (
	"net/http"
)

type Message struct {
	StatusCode int
	Result     interface{}
}

// Core Response
type Core struct {
	Status  http.ConnState `json:"status"`
	Message string         `json:"message"`
}

// Change Response
type Change struct {
	Status   http.ConnState `json:"status"`
	Message  string         `json:"message"`
	Affected int            `json:"affected"`
}

// Retrieve Response
type Retrieve struct {
	Status  http.ConnState `json:"status"`
	Message string         `json:"message"`
	Count   int            `json:"count"`
	Results interface{}    `json:"results"`
}

// SendError calls Send by without a count or results
func SendError(w http.ResponseWriter, status http.ConnState, message string) Message {
	return Send(w, status, message, 0, nil)
}

// Send writes struct to the writer using a format
func Send(
	w http.ResponseWriter,
	status http.ConnState,
	message string,
	count int,
	results interface{}) Message {

	var result interface{}

	// Determine the best format
	if count < 1 {
		result = &Core{
			Status:  status,
			Message: message,
		}
	} else if results == nil {
		result = &Change{
			Status:   status,
			Message:  message,
			Affected: count,
		}
	} else {
		result = &Retrieve{
			Status:  status,
			Message: message,
			Count:   count,
			Results: results,
		}
	}

	statusCode := int(status)

	return Message{
		StatusCode: statusCode,
		Result:     result}
}

// SendJSON writes a struct to the writer
// func SendJSON(w http.ResponseWriter, i interface{}) {
// 	js, err := json.Marshal(i)
// 	if err != nil {
// 		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)
// }
