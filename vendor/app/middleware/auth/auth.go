package auth

import (
	"errors"
	"net/http"

	"app/model"
	"app/shared/response"
	"app/shared/token"
	// "github.com/gorilla/context"
)

const (
	authorize        = "authorize"
	authTokenHeader  = "auth_token"
	authToken        = "auth_token"
	authTokenMissing = "auth_token is missing"
	authTokenInvalid = "auth_token is invalid"
)

// Handler will authorize HTTP requests
func Handler(env *model.Env, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at, err := tokenFromHeader(r, authTokenHeader)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, err.Error())
			return
		}

		t, err := token.Decrypt(at)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
			return
		}

		uID := string(t)
		u, err := env.DB.UserExist("id", uID)
		if err != nil {
			response.SendError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if u == false {
			response.SendError(w, http.StatusUnauthorized, authTokenInvalid)
			return
		}

		// todo: add to context
		//context.Set(r, "userID", uID)

		next.ServeHTTP(w, r)
	})
}

func tokenFromHeader(r *http.Request, header string) (string, error) {
	token := r.Header.Get(header)

	l := len(token)

	if l == 0 {
		return "", errors.New(authTokenMissing)
	}

	return token, nil
}
