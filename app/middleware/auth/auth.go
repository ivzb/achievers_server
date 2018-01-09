package auth

import (
	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

const (
	authorize        = "authorize"
	authTokenHeader  = "auth_token"
	authToken        = "auth_token"
	authTokenMissing = "auth_token is missing"
	authTokenInvalid = "auth_token is invalid"
)

// Handler will authorize HTTP requests
func Handler(app app.App) app.App {
	prevHandler := app.Handler

	app.Handler = func(env *model.Env) *response.Message {
		at, err := request.HeaderValue(app.Env.Request, authTokenHeader)

		if err != nil {
			return response.Unauthorized(authTokenMissing)
		}

		uID, err := app.Env.Token.Decrypt(at)

		if err != nil {
			return response.Unauthorized(authTokenInvalid)
		}

		exists, err := app.Env.DB.UserExists(uID)
		if err != nil {
			return response.Unauthorized(authTokenInvalid)
		}

		if exists == false {
			return response.Unauthorized(authTokenInvalid)
		}

		app.Env.UserID = uID

		return prevHandler(env)
	}

	return app
}
