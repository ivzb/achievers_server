package controller

import (
	"log"
	"net/http"

	"app/model"
	"app/shared/response"
	"app/shared/token"
)

func UserAuth(env *model.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			response.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" {
			response.SendError(w, http.StatusBadRequest, "missing email")
			return
		}

		if password == "" {
			response.SendError(w, http.StatusBadRequest, "missing password")
			return
		}

		uID, err := env.DB.UserAuth(email, password)

		if err != nil {
			response.SendError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		t, err := token.Encrypt(uID)

		if err != nil {
			response.SendError(w, http.StatusInternalServerError, FriendlyError)
			return
		}

		response.Send(w, http.StatusOK, "authorized", 1, t.AuthToken)
	})
}

func UserCreate(env *model.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			response.SendError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if first_name == "" {
			response.SendError(w, http.StatusBadRequest, "missing first_name")
			return
		}

		if last_name == "" {
			response.SendError(w, http.StatusBadRequest, "missing last_name")
			return
		}

		if email == "" {
			response.SendError(w, http.StatusBadRequest, "missing email")
			return
		}

		if password == "" {
			response.SendError(w, http.StatusBadRequest, "missing password")
			return
		}

		exist, err := env.DB.UserExist("email", email)

		if err != nil {
			response.SendError(w, http.StatusInternalServerError, FriendlyError)
			return
		}

		if exist {
			response.SendError(w, http.StatusBadRequest, "email already exists")
			return
		}

		id, err := env.DB.UserCreate(first_name, last_name, email, password)

		if err != nil {
			log.Println(err)
			response.SendError(w, http.StatusInternalServerError, FriendlyError)
			return
		}

		response.Send(w, http.StatusCreated, ItemCreated, 1, id)
	})
}
