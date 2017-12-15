package auth

import (
	"app/middleware/app"
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
func Handler(handler app.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at, err := request.GetHeader(r, authTokenHeader)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, authTokenMissing)
			return
		}

		uID, err := handler.Env.Token.Decrypt(at)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
			return
		}

		exists, err := handler.Env.DB.Exists("user", "id", uID)
		if err != nil {
			response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
			return
		}

		if exists == false {
			response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
			return
		}

		handler.Env.UserId = uID

		handler.ServeHTTP(w, r)
	})
}
