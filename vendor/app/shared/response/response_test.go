package response

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type mock struct {
	Id    string `json:"id"`
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

func fail(t *testing.T, value string, expected interface{}, actual interface{}) {
	t.Fatalf("Send returned unexpected %v:\nexpected %#v,\nactual %#v",
		value, expected, actual)
}

func TestSendMultipleResults(t *testing.T) {
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
		t.Fatalf("Send returned wrong value: expected %v, actual %v",
			status, response.StatusCode)
	}

	switch actualResult := response.Result.(type) {
	case *Retrieve:
		if !cmp.Equal(expectedResult, actualResult) {
			fail(t, "Results", expectedResult, actualResult)
		}
	default:
		t.Fatalf("Send returned unexpected type: expected %v, actual %v",
			"Retrive", actualResult)
	}
}

func TestSendNoResults(t *testing.T) {
	status := http.StatusOK
	message := "response_message"
	var results interface{} = nil
	length := 5

	expectedResult := &Change{
		Status:   status,
		Message:  message,
		Affected: length,
	}

	response := Send(status, message, length, results)

	// Check the status code is what we expect.
	if status != response.StatusCode {
		t.Fatalf("Send returned wrong value: expected %v, actual %v",
			status, response.StatusCode)
	}

	switch actualResult := response.Result.(type) {
	case *Change:
		if !cmp.Equal(expectedResult, actualResult) {
			fail(t, "Results", expectedResult, actualResult)
		}
	default:
		t.Fatalf("Send returned unexpected type: expected %v, actual %v",
			"Retrive", actualResult)
	}
}
