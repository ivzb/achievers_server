package middleware

import (
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

// AuthHandler will authorize HTTP requests
func AuthHandler(h AppHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at, err := valueFromHeader(r, authTokenHeader)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, authTokenMissing)
			return
		}

		t, err := h.Env.Token.Decrypt(at)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
			return
		}

		uID := string(t)
		u, err := h.Env.DB.Exist("user", "id", uID)
		if err != nil {
			response.SendError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if u == false {
			response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
			return
		}

		h.Env.UserId = uID

		h.ServeHTTP(w, r)
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
