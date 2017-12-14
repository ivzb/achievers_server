package middleware

import (
	"errors"
	"net/http"

	"app/model"
	"app/shared/response"
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
func AuthHandler(env *model.Env, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at, err := valueFromHeader(r, authTokenHeader)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, authTokenMissing)
			return
		}

		t, err := env.Token.Decrypt(at)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
			return
		}

		uID := string(t)
		u, err := env.DB.Exist("user", "id", uID)
		if err != nil {
			response.SendError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if u == false {
			response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
			return
		}

		env.UserId = uID

		next.ServeHTTP(w, r)
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
