package controller

import (
	"fmt"
	"log"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func UserAuth(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, POST) {
		return response.MethodNotAllowed()
	}

	auth := &model.Auth{}
	err := form.ModelValue(env.Request, auth)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if auth.Email == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, email))
	}

	if auth.Password == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, password))
	}

	exists, err := env.DB.UserEmailExists(auth.Email)

	if err != nil {
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(email)
	}

	uID, err := env.DB.UserAuth(auth)

	if err != nil {
		log.Println(err)
		return response.InternalServerError()
	}

	token, err := env.Token.Encrypt(uID)

	if err != nil {
		return response.InternalServerError()
	}

	return response.Created(authToken, token)
}

func UserCreate(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, POST) {
		return response.MethodNotAllowed()
	}

	usr := &model.User{}
	err := form.ModelValue(env.Request, usr)

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
		return response.InternalServerError()
	}

	if exists {
		return response.BadRequest(fmt.Sprintf(formatAlreadyExists, email))
	}

	id, err := env.DB.UserCreate(usr)

	if err != nil {
		return response.InternalServerError()
	}

	return response.Created(user, id)
}
