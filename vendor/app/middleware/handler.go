package middleware

import (
	"app/model"
	"net/http"
)

// type He func(e *model.Env, w http.ResponseWriter, r *http.Request)

// type Handler struct {
// 	*model.Env
// 	H He
// }

// ServeHTTP allows our Handler type to satisfy http.Handler.
// func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h.H(h.Env, w, r)
// }

// type AppHandler func(env *model.Env, w http.ResponseWriter, r *http.Request)

type AppHandler struct {
	Env *model.Env
	H   func(env *model.Env, w http.ResponseWriter, r *http.Request)
}

// type AppHandler func(*model.Env, http.ResponseWriter, *http.Request)

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn.H(fn.Env, w, r)
}
