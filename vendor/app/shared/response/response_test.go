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
		Status:  status,
		Message: message,
		Length:  length,
		Results: results,
	}

	response := Send(status, message, length, results)

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
		Status:   status,
		Message:  message,
		Affected: length,
	}

	response := Send(status, message, length, results)

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
		Status:  status,
		Message: message,
	}

	response := Send(status, message, length, results)

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
	status := http.StatusOK
	message := "response_message"

	expectedResult := &Core{
		Status:  status,
		Message: message,
	}

	response := SendError(status, message)

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
