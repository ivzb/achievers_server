package auth

import (
	"net/http"
	"encoding/json"

	mtoken "app/model/token"
	"app/shared/response"
	"app/shared/token"
)

const (
	authorize = "authorize"
	authToken = "auth_token"
	authTokenMissing = "auth_token is missing"
	authTokenInvalid = "auth_token is invalid"
)

// Handler will authorize HTTP requests
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at := r.Header.Get(authToken)

		if at == "" {
			response.Send(w, http.StatusUnauthorized, authTokenMissing, 0, nil)
			return
		}

		s := `{"auth_token":"` + at + `"}`
		var t mtoken.Entity

		if err := json.Unmarshal([]byte(s), &t); err != nil {
			response.Send(w, http.StatusUnauthorized, authTokenInvalid, 0, nil)
			return
		}

		v := token.Validate(t.AuthToken)

		if v == false {
			response.Send(w, http.StatusUnauthorized, authTokenInvalid, 0, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}