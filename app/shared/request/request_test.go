package request

import (
	"net/http/httptest"
	"testing"
)

func TestIsMethod_ValidMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)

	actualValue := IsMethod(req, "GET")

	expectedValue := true

	// Check the status code is what we expect.
	if expectedValue != actualValue {
		t.Fatalf("GetHeader returned wrong value: expected %v, actual %v",
			expectedValue, actualValue)
	}
}

func TestIsMethod_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)

	actualValue := IsMethod(req, "POST")

	expectedValue := false

	// Check the status code is what we expect.
	if expectedValue != actualValue {
		t.Fatalf("GetHeader returned wrong value: expected %v, actual %v",
			expectedValue, actualValue)
	}
}

func TestGetHeader_ValidHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)

	key := "header_key"
	expectedValue := "header_value"

	req.Header.Add(key, expectedValue)

	actualValue, err := HeaderValue(req, key)

	if err != nil {
		t.Fatalf("GetHeader returned error: %v",
			err)
	}

	// Check the status code is what we expect.
	if expectedValue != actualValue {
		t.Fatalf("GetHeader returned wrong value: expected %v, actual %v",
			expectedValue, actualValue)
	}
}

func TestGetHeader_MissingHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/auth", nil)

	expectedValue := headerMissing

	key := "header_key"

	_, err := HeaderValue(req, key)

	if err == nil {
		t.Fatalf("GetHeader expected error, but it was nil")
	}

	actualValue := err.Error()

	// Check the status code is what we expect.
	if expectedValue != actualValue {
		t.Fatalf("GetHeader returned wrong value: expected %v, actual %v",
			expectedValue, actualValue)
	}
}
