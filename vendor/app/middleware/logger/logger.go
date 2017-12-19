package logger

import (
	"app/middleware/app"
	"app/model"
	"app/shared/response"
	"fmt"
	"net/http"
)

const (
	FriendlyError = "an error occurred, please try again later"
)

// Handler will log the HTTP requests
func Handler(handler app.Handler) app.Handler {
	prevH := handler.H

	handler.H = func(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
		log := fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		err := handler.Env.Logger.Log(log)

		if err != nil {
			return response.SendError(http.StatusInternalServerError, FriendlyError)
		}

		return prevH(env, w, r)
	}

	return handler
}
