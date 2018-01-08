package controller

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func HomeIndex(env *model.Env) *response.Message {
	if !env.Request.IsMethod(GET) {
		return response.MethodNotAllowed()
	}

	return response.Ok(
		home,
		1,
		welcome)
}
