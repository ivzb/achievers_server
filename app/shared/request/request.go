package request

import (
	"errors"
	"net/http"
)

const (
	headerMissing = "header is missing"
)

func GetHeader(r *http.Request, key string) (string, error) {
	value := r.Header.Get(key)

	if value == "" {
		return "", errors.New(headerMissing)
	}

	return value, nil
}
