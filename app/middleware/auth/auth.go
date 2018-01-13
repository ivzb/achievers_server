package auth

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/middleware/app"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

// Handler will authorize HTTP requests
func Handler(app app.App) app.App {
	prevHandler := app.Handler

	app.Handler = func(env *model.Env) *response.Message {
		at, err := request.HeaderValue(app.Env.Request, consts.AuthToken)

		if err != nil {
			return response.Unauthorized(fmt.Sprintf(consts.FormatMissing, consts.AuthToken))
		}

		uID, err := app.Env.Token.Decrypt(at)

		if err != nil {
			return response.Unauthorized(fmt.Sprintf(consts.FormatInvalid, consts.AuthToken))
		}

		exists, err := app.Env.DB.UserExists(uID)

		if err != nil {
			return response.InternalServerError()
		}

		if exists == false {
			return response.Unauthorized(fmt.Sprintf(consts.FormatInvalid, consts.AuthToken))
		}

		app.Env.UserID = uID

		return prevHandler(env)
	}

	return app
}
