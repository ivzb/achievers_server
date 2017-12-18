package app

import (
	"app/model"
	"app/shared/response"
	"encoding/json"
	"net/http"
)

type Handle func(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message

type Handler struct {
	Env *model.Env
	H   Handle
}

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := fn.H(fn.Env, w, r)

	js, err := json.Marshal(response.Result)
	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(js)
}
