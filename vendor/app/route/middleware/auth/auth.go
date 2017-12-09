package auth

import (
	"errors"
	"net/http"

	"app/model/user"
	"app/shared/response"
	"app/shared/token"

	"github.com/gorilla/context"
)

const (
	authorize        = "authorize"
	authTokenHeader  = "auth_token"
	authToken        = "auth_token"
	authTokenMissing = "auth_token is missing"
	authTokenInvalid = "auth_token is invalid"
)

// Handler will authorize HTTP requests
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at, err := tokenFromHeader(r, authTokenHeader)

		if err != nil {
			response.Send(w, http.StatusUnauthorized, authTokenMissing, 0, nil)
			return
		}

		t, err := token.Decrypt(at)

		if err != nil {
			response.Send(w, http.StatusUnauthorized, authTokenInvalid, 0, nil)
			return
		}

		uID := string(t)
		u, err := user.Exist(uID)
		if err != nil || u == false {
			response.Send(w, http.StatusUnauthorized, authTokenInvalid, 0, nil)
			return
		}

		context.Set(r, "userID", uID)

		next.ServeHTTP(w, r)
	})
}

func tokenFromHeader(r *http.Request, header string) (string, error) {
	token := r.Header.Get(header)

	l := len(token)

	if l == 0 {
		return "", errors.New("Missing or invalid token in the request header")
	}

	return token, nil
}
