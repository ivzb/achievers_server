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
//Form() Former
//IsMethod(expected string) bool
//HeaderValue(key string) (string, error)
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
		Form:        f,
		httpRequest: r,
		RemoteAddr:  r.RemoteAddr,
		Method:      r.Method,
		URL:         r.URL,
	}
}

//// Form returns request form
//func (r *Request) Form() Former {
//return *r.form
//}

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
