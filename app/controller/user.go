package controller

import (
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
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

	usr := &model.User{}
	err := env.Former.Map(r, usr)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if usr.FirstName == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, firstName))
	}

	if usr.LastName == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, lastName))
	}

	if usr.Email == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, email))
	}

	if usr.Password == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, password))
	}

	exists, err := env.DB.UserEmailExists(usr.Email)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	if exists {
		return response.BadRequest(fmt.Sprintf(formatAlreadyExists, email))
	}

	id, err := env.DB.UserCreate(usr)

	if err != nil {
		return response.InternalServerError(friendlyErrorMessage)
	}

	return response.Created(fmt.Sprintf(formatCreated, user), id)
}
