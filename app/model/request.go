package model

import (
	"net/http"
)

type Requester interface {
	IsMethod(expected string) bool
}

type Request struct {
	httpRequest *http.Request
}

func NewRequester(r *http.Request) *Request {
	return &Request{
		httpRequest: r,
	}
}

// IsMethod returns whether actual request method equals expected one
func (r *Request) IsMethod(expected string) bool {
	actual := r.httpRequest.Method

	return actual == expected
}
