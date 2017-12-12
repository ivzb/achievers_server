package auth

import (
	"errors"
	"net/http"
	// "crypto/rsa"

	"app/model"
	"app/shared/response"
	"app/shared/token"
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
func Handler(env *model.Env, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at, err := valueFromHeader(r, authTokenHeader)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, authTokenMissing)
			return
		}

		t, err := token.Decrypt(env.Token.GetPrivateKey(), at)

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

		env.Store["user_id"] = uID

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