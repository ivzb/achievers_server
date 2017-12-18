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

func UserAuth(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
	if r.Method != "POST" {
		return response.SendError(w, http.StatusMethodNotAllowed, MethodNotAllowedErrorMessage)
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" {
		return response.SendError(w, http.StatusBadRequest, MissingEmailErrorMessage)
	}

	if password == "" {
		return response.SendError(w, http.StatusBadRequest, MissingPasswordErrorMessage)
	}

	uID, err := env.DB.UserAuth(email, password)

	if err != nil {
		return response.SendError(w, http.StatusUnauthorized, Unauthorized)
	}

	token, err := env.Tokener.Encrypt(uID)

	if err != nil {
		return response.SendError(w, http.StatusInternalServerError, FriendlyErrorMessage)
	}

	return response.Send(w, http.StatusOK, "authorized", 1, token)
}

func UserCreate(env *model.Env, w http.ResponseWriter, r *http.Request) response.Message {
	if r.Method != "POST" {
		return response.SendError(w, http.StatusMethodNotAllowed, MethodNotAllowedErrorMessage)
	}

	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if first_name == "" {
		return response.SendError(w, http.StatusBadRequest, MissingFirstNameErrorMessage)
	}

	if last_name == "" {
		return response.SendError(w, http.StatusBadRequest, MissingLastNameErrorMessage)
	}

	if email == "" {
		return response.SendError(w, http.StatusBadRequest, MissingEmailErrorMessage)
	}

	if password == "" {
		return response.SendError(w, http.StatusBadRequest, MissingPasswordErrorMessage)
	}

	exists, err := env.DB.Exists("user", "email", email)

	if err != nil {
		return response.SendError(w, http.StatusInternalServerError, FriendlyErrorMessage)
	}

	if exists {
		return response.SendError(w, http.StatusBadRequest, EmailAlreadyExistsErrorMessage)
	}

	id, err := env.DB.UserCreate(first_name, last_name, email, password)

	if err != nil {
		return response.SendError(w, http.StatusInternalServerError, FriendlyErrorMessage)
	}

	return response.Send(w, http.StatusCreated, UserCreated, 1, id)
}
