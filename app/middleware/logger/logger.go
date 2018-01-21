package logger

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/shared/env"
	"github.com/ivzb/achievers_server/app/shared/response"
)

// Handler will log the HTTP requests
func Handler(app app.App) app.App {
	prevHandler := app.Handler

	app.Handler = func(env *env.Env) *response.Message {
		message := fmt.Sprintf("%s %s %s",
			env.Request.RemoteAddr,
			env.Request.Method,
			env.Request.URL)

		app.Env.Log.Message(message)

		return prevHandler(env)
	}

	return app
}
