package request

import (
	"errors"
	"net/http"
)

const (
	headerMissing    = "header is missing"
	formatMissing    = "missing %s"
	formatInvalid    = "invalid %s"
	methodNotAllowed = "method not allowed"
)

// IsMethod returns whether actual request method equals expected one
func IsMethod(r *http.Request, expected string) bool {
	actual := r.Method

	return actual == expected
}

// HeaderValue returns header value by given key, error if nil or empty
func HeaderValue(r *http.Request, key string) (string, error) {
	value := r.Header.Get(key)

	if value == "" {
		return "", errors.New(headerMissing)
	}

	return value, nil
}
