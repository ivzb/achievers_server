package request

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

const (
	headerMissing    = "header is missing"
	formatMissing    = "missing %s"
	formatInvalid    = "invalid %s"
	methodNotAllowed = "method not allowed"
)

// GetHeader returns header value by given key, error if nil or empty
func GetHeader(r *http.Request, key string) (string, error) {
	value := r.Header.Get(key)

	if value == "" {
		return "", errors.New(headerMissing)
	}

	return value, nil
}

// Is returns whether actual request method equals expected one
func Is(r *http.Request, expected string) bool {
	actual := r.Method

	return actual == expected
}

// FormValue returns key by value and error if not found
func FormValue(r *http.Request, key string) (string, error) {
	value := r.FormValue(key)

	if value == "" {
		return "", errors.New(fmt.Sprintf(formatMissing, key))
	}

	return value, nil
}

// FormValueInt returns key by value casted to int and error if not found
func FormIntValue(r *http.Request, key string, min int) (int, error) {
	value, err := FormValue(r, key)

	if err != nil {
		return 0, err
	}

	castedValue, err := strconv.Atoi(value)

	if err != nil || castedValue < min {
		return 0, errors.New(fmt.Sprintf(formatInvalid, key))
	}

	return castedValue, nil
}
