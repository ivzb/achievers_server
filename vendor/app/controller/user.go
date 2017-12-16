package controller

import (
	"net/http"

	"app/model"
	"app/shared/response"
)

const (
	UserCreated = "user created"

	MissingFirstNameErrorMessage = "missing first_name"
	MissingLastNameErrorMessage  = "missing last_name"
	MissingEmailErrorMessage     = "missing email"
	MissingPasswordErrorMessage  = "missing password"

	EmailAlreadyExistsErrorMessage = "email already exists"
)

func UserAuth(env *model.Env, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response.SendError(w, http.StatusMethodNotAllowed, MethodNotAllowedErrorMessage)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" {
		response.SendError(w, http.StatusBadRequest, MissingEmailErrorMessage)
		return
	}

	if password == "" {
		response.SendError(w, http.StatusBadRequest, MissingPasswordErrorMessage)
		return
	}

	uID, err := env.DB.UserAuth(email, password)

	if err != nil {
		response.SendError(w, http.StatusUnauthorized, Unauthorized)
		return
	}

	token, err := env.Tokener.Encrypt(uID)

	if err != nil {
		response.SendError(w, http.StatusInternalServerError, FriendlyErrorMessage)
		return
	}

	response.Send(w, http.StatusOK, "authorized", 1, token)
}

func UserCreate(env *model.Env, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response.SendError(w, http.StatusMethodNotAllowed, MethodNotAllowedErrorMessage)
		return
	}

	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if first_name == "" {
		response.SendError(w, http.StatusBadRequest, MissingFirstNameErrorMessage)
		return
	}

	if last_name == "" {
		response.SendError(w, http.StatusBadRequest, MissingLastNameErrorMessage)
		return
	}

	if email == "" {
		response.SendError(w, http.StatusBadRequest, MissingEmailErrorMessage)
		return
	}

	if password == "" {
		response.SendError(w, http.StatusBadRequest, MissingPasswordErrorMessage)
		return
	}

	exists, err := env.DB.Exists("user", "email", email)

	if err != nil {
		response.SendError(w, http.StatusInternalServerError, FriendlyErrorMessage)
		return
	}

	if exists {
		response.SendError(w, http.StatusBadRequest, EmailAlreadyExistsErrorMessage)
		return
	}

	id, err := env.DB.UserCreate(first_name, last_name, email, password)

	if err != nil {
		response.SendError(w, http.StatusInternalServerError, FriendlyErrorMessage)
		return
	}

	response.Send(w, http.StatusCreated, UserCreated, 1, id)
}
