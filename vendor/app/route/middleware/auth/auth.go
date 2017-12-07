package auth

import (
	"encoding/json"
	"net/http"

	mt "app/model/token"
	"app/model/user"
	"app/shared/response"
	"app/shared/token"

	"github.com/gorilla/context"
)

const (
	authorize        = "authorize"
	authToken        = "auth_token"
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
		var t mt.Entity

		if err := json.Unmarshal([]byte(s), &t); err != nil {
			response.Send(w, http.StatusUnauthorized, authTokenInvalid, 0, nil)
			return
		}

		v, err := token.Extract(t.AuthToken)

		if err != nil {
			response.Send(w, http.StatusUnauthorized, authTokenInvalid, 0, nil)
			return
		}

		uID := string(v)
		u, err := user.Exist(uID)
		if err != nil || u == false {
			response.Send(w, http.StatusUnauthorized, authTokenInvalid, 0, nil)
			return
		}

		context.Set(r, "userID", uID)

		next.ServeHTTP(w, r)
	})
}
