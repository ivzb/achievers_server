package middleware

import (
    "net/http"
    "app/model"
)

type Handler struct {
	*model.Env
	H func(e *model.Env, w http.ResponseWriter, r *http.Request)
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.H(h.Env, w, r)
}