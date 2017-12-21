package response

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

// SendError calls Send by without a count or results
func SendError(status int, message string) Message {
	return Send(status, message, 0, nil)
}

// Send writes struct to the writer using a format
func Send(
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
