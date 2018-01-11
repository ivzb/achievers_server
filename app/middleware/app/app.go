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

func (app App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	size := app.Env.Config.Server.MaxBytesReader // in Mb's
	req.Body = http.MaxBytesReader(w, req.Body, size)

	app.Env.Request = req

	resp := app.Handler(app.Env)

	switch resp.Type {
	case response.TypeFile:
		serveFile(w, req, resp)
	default:
		serveJSON(w, resp)
	}
}

func serveJSON(w http.ResponseWriter, resp *response.Message) {
	js, err := json.Marshal(resp.Result)

	if err != nil {
		http.Error(w, "JSON Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(js)
}

func serveFile(w http.ResponseWriter, r *http.Request, msg *response.Message) {
	filepath := msg.Result.(string)
	http.ServeFile(w, r, filepath)
}
