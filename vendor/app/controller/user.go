package controller

import (
	"log"
	"net/http"

	"app/model"
	"app/shared/response"
)

func UserAuth(env *model.Env, w http.ResponseWriter, r *http.Request) {
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

	t, err := env.Token.Encrypt(uID)

	if err != nil {
		log.Println(err)
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	response.Send(w, http.StatusOK, "authorized", 1, t)
}

func UserCreate(env *model.Env, w http.ResponseWriter, r *http.Request) {
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

	exists, err := env.DB.Exists("user", "email", email)

	if err != nil {
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	if exists {
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
}
