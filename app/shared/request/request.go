package request

import (
	"errors"
	"net/http"
)

const (
	headerMissing = "header is missing"
)

// GetHeader returns header value by given key, error if nil or empty
func GetHeader(r *http.Request, key string) (string, error) {
	value := r.Header.Get(key)

	if value == "" {
		return "", errors.New(headerMissing)
	}

	return value, nil
}
