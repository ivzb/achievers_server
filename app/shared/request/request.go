package request

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/shared/consts"
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
		return "", errors.New(fmt.Sprintf(consts.FormatMissing, consts.Header))
	}

	return value, nil
}
