package logger

import (
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

const (
	friendlyError = "an error occurred, please try again later"
)

// Handler will log the HTTP requests
func Handler(handler app.Handler) app.Handler {
	prevH := handler.H

	handler.H = func(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
		message := fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.Env.Logger.Message(message)

		return prevH(env, w, r)
	}

	return handler
}
