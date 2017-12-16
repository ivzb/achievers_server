package app

import (
	"app/model"
	"net/http"
)

type Handle func(env *model.Env, w http.ResponseWriter, r *http.Request)

type Handler struct {
	Env *model.Env
	H   Handle
}

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn.H(fn.Env, w, r)
}
