package controller

import (
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

const (
	user = "user"

	firstName = "first_name"
	lastName  = "last_name"
	email     = "email"
	password  = "password"
)

func UserAuth(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "POST" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	eml := r.FormValue("email")
	pwd := r.FormValue("password")

	if eml == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, email))
	}

	if pwd == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, password))
	}

	uID, err := env.DB.UserAuth(eml, pwd)

	if err != nil {
		return response.Unauthorized(unauthorized)
	}

	token, err := env.Tokener.Encrypt(uID)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Created(authorized, token)
}

func UserCreate(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "POST" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	fnm := r.FormValue("first_name")
	lnm := r.FormValue("last_name")
	eml := r.FormValue("email")
	pwd := r.FormValue("password")

	if fnm == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, firstName))
	}

	if lnm == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, lastName))
	}

	if eml == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, email))
	}

	if pwd == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, password))
	}

	exists, err := env.DB.Exists("user", "email", email)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if exists {
		return response.BadRequest(fmt.Sprintf(formatAlreadyExists, email))
	}

	id, err := env.DB.UserCreate(fnm, lnm, eml, pwd)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Created(fmt.Sprintf(formatCreated, user), id)
}
