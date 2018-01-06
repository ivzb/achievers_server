package auth

import (
	"net/http"

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
func Handler(handler app.Handler) app.Handler {
	prevH := handler.H

	handler.H = func(env *model.Env, w http.ResponseWriter, r *http.Request) *response.Message {
		at, err := request.GetHeader(r, authTokenHeader)

		if err != nil {
			return response.Unauthorized(authTokenMissing)
		}

		uID, err := handler.Env.Token.Decrypt(at)

		if err != nil {
			return response.Unauthorized(authTokenInvalid)
		}

		exists, err := handler.Env.DB.UserExists(uID)
		if err != nil {
			return response.Unauthorized(authTokenInvalid)
		}

		if exists == false {
			return response.Unauthorized(authTokenInvalid)
		}

		handler.Env.UserId = uID

		return prevH(env, w, r)
	}

	return handler
}
