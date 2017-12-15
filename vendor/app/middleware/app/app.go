package app

import (
	"app/model"
	"net/http"
)

type Handler struct {
	Env *model.Env
	H   func(env *model.Env, w http.ResponseWriter, r *http.Request)
}

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn.H(fn.Env, w, r)
}
