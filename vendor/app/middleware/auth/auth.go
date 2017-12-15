package auth

import (
	"app/middleware/app"
	"app/shared/response"
	"errors"
	"net/http"
)

const (
	headerMissing    = "header is missing"
	authorize        = "authorize"
	authTokenHeader  = "auth_token"
	authToken        = "auth_token"
	authTokenMissing = "auth_token is missing"
	authTokenInvalid = "auth_token is invalid"
)

// Handler will authorize HTTP requests
func Handler(handler app.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at, err := valueFromHeader(r, authTokenHeader)

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
			response.SendError(w, http.StatusUnauthorized, err.Error())
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

func valueFromHeader(r *http.Request, key string) (string, error) {

	value := r.Header.Get(key)

	l := len(value)

	if l == 0 {
		return "", errors.New(headerMissing)
	}

	return value, nil
}
