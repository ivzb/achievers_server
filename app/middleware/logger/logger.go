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
func Handler(app app.App) app.App {
	prevHandler := app.Handler

	app.Handler = func(env *model.Env, r *http.Request) *response.Message {
		message := fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		app.Env.Log.Message(message)

		return prevHandler(env, r)
	}

	return app
}
