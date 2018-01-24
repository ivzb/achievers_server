package response

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type mock struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func mocks() []*mock {
	mocks := make([]*mock, 0)

	mocks = append(mocks, &mock{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"title",
	})

	mocks = append(mocks, &mock{
		"93821a67-9c82-96e4-dc3c-423e5581d036",
		"another title",
	})

	return mocks
}

func fail(t *testing.T, method string, expected interface{}, actual interface{}) {
	t.Fatalf("%v returned unexpected value:\nexpected %#v,\nactual %#v",
		method, expected, actual)
}

func TestSend_MultipleResults(t *testing.T) {
	status := http.StatusOK
	message := "response_message"
	results := mocks()
	length := len(results)

	expectedResult := &Retrieve{
		Message: message,
		Results: results,
	}

	response := send(status, message, length, results)

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "Send", status, response.StatusCode)
	}

	switch actualResult := response.Result.(type) {
	case *Retrieve:
		if !cmp.Equal(expectedResult, actualResult) {
			fail(t, "Send", expectedResult, actualResult)
		}
	default:
		fail(t, "Send", "Retrive", actualResult)
	}
}

func TestSend_NoResults(t *testing.T) {
	status := http.StatusOK
	message := "response_message"
	var results interface{}
	length := 5

	expectedResult := &Change{
		Message:  message,
		Affected: length,
	}

	response := send(status, message, length, results)

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "Send", status, response.StatusCode)
	}

	switch actualResult := response.Result.(type) {
	case *Change:
		if !cmp.Equal(expectedResult, actualResult) {
			fail(t, "Send", expectedResult, actualResult)
		}
	default:
		fail(t, "Send", "Change", actualResult)
	}
}

func TestSend_ZeroLength(t *testing.T) {
	status := http.StatusOK
	message := "response_message"
	var results interface{}
	length := 0

	expectedResult := &Core{
		Message: message,
	}

	response := send(status, message, length, results)

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "Send", status, response.StatusCode)
	}

	switch actualResult := response.Result.(type) {
	case *Core:
		if !cmp.Equal(expectedResult, actualResult) {
			fail(t, "Send", expectedResult, actualResult)
		}
	default:
		fail(t, "Send", "Core", actualResult)
	}
}

func TestSendError(t *testing.T) {
	status := http.StatusBadRequest
	message := "response_message"

	expectedResult := &Core{
		Message: message,
	}

	response := sendError(status, message)

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "Send", status, response.StatusCode)
	}

	switch actualResult := response.Result.(type) {
	case *Core:
		if !cmp.Equal(expectedResult, actualResult) {
			fail(t, "Send", expectedResult, actualResult)
		}
	default:
		fail(t, "Send", "Core", actualResult)
	}
}

func TestOk(t *testing.T) {
	status := http.StatusOK

	response := Ok("", 0, nil)

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "OK", status, response.StatusCode)
	}
}

func TestFile(t *testing.T) {
	status := http.StatusOK
	path := "path_to_file"
	expectedResult := &Core{Message: path}

	response := File(path)

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "OK", status, response.StatusCode)
	}

	actualResult := response.Result.(*Core)

	if !cmp.Equal(expectedResult, actualResult) {
		fail(t, "response.File", expectedResult, actualResult)
	}
}

func TestCreated(t *testing.T) {
	status := http.StatusCreated

	response := Created("", nil)

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "Created", status, response.StatusCode)
	}
}

func TestBadRequest(t *testing.T) {
	status := http.StatusBadRequest

	response := BadRequest("")

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "BadRequest", status, response.StatusCode)
	}
}

func TestUnauthorized(t *testing.T) {
	status := http.StatusUnauthorized

	response := Unauthorized("")

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "Unauthorized", status, response.StatusCode)
	}
}

func TestNotFound(t *testing.T) {
	status := http.StatusNotFound

	response := NotFound("")

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "NotFound", status, response.StatusCode)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	status := http.StatusMethodNotAllowed

	response := MethodNotAllowed()

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "MethodNotAllowed", status, response.StatusCode)
	}
}

func TestInternalServerError(t *testing.T) {
	status := http.StatusInternalServerError

	response := InternalServerError()

	// Check the status code is what we expect.
	if status != response.StatusCode {
		fail(t, "InternalServerError", status, response.StatusCode)
	}
}
