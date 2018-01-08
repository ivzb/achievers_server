package model

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	headerMissing = "header is missing"
)

//type Requester interface {
//IsMethod(expected string) bool
//Header(key string) (string, error)
//}

type Request struct {
	UserID     string
	Form       Former
	RemoteAddr string
	Method     string
	URL        *url.URL

	httpRequest *http.Request
}

func NewRequest(r *http.Request, f Former) *Request {
	return &Request{
		UserID:     "", // to be determined by auth middleware
		Form:       f,
		RemoteAddr: r.RemoteAddr,
		Method:     r.Method,
		URL:        r.URL,

		httpRequest: r,
	}
}

// IsMethod returns whether actual request method equals expected one
func (r *Request) IsMethod(expected string) bool {
	actual := r.httpRequest.Method

	return actual == expected
}

// HeaderValue returns header value by given key, error if nil or empty
func (r *Request) HeaderValue(key string) (string, error) {
	value := r.httpRequest.Header.Get(key)

	if value == "" {
		return "", errors.New(headerMissing)
	}

	return value, nil
}
