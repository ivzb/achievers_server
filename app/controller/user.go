package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func UserAuth(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.POST) {
		return response.MethodNotAllowed()
	}

	auth := &model.Auth{}
	err := form.ModelValue(env.Request, auth)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if auth.Email == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Email))
	}

	if auth.Password == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Password))
	}

	exists, err := env.DB.UserEmailExists(auth.Email)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(consts.Email)
	}

	uID, err := env.DB.UserAuth(auth)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	token, err := env.Token.Encrypt(uID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Created(consts.AuthToken, token)
}

func UserCreate(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.POST) {
		return response.MethodNotAllowed()
	}

	usr := &model.User{}
	err := form.ModelValue(env.Request, usr)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	prfl := &model.Profile{}
	_ = form.ModelValue(env.Request, prfl)

	if usr.Email == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Email))
	}

	if usr.Password == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Password))
	}

	exists, err := env.DB.UserEmailExists(usr.Email)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if exists {
		return response.BadRequest(fmt.Sprintf(consts.FormatAlreadyExists, consts.Email))
	}

	userID, err := env.DB.UserCreate(usr)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	_, err = env.DB.Profile().ProfileCreate(prfl, userID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Created(consts.User, userID)
}
