package auth

import (
	"app/middleware/app"
	"app/model"
	"app/shared/request"
	"app/shared/response"
	"net/http"
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

	handler.H = func(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
		at, err := request.GetHeader(r, authTokenHeader)

		if err != nil {
			return response.SendError(w, http.StatusUnauthorized, authTokenMissing)
		}

		uID, err := handler.Env.Tokener.Decrypt(at)

		if err != nil {
			return response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
		}

		exists, err := handler.Env.DB.Exists("user", "id", uID)
		if err != nil {
			return response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
		}

		if exists == false {
			return response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
		}

		handler.Env.UserId = uID

		return prevH(env, w, r)
	}

	return handler
}
