package app

import (
	"encoding/json"
	"net/http"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

type Handler func(env *model.Env) *response.Message

type App struct {
	Env     *model.Env
	Handler Handler
}

func (app App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	form := model.NewForm(r)
	app.Env.Request = model.NewRequest(r, form)

	response := app.Handler(app.Env)

	js, err := json.Marshal(response.Result)

	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(js)
}
