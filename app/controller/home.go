package controller

import (
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func HomeIndex(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	return response.Ok(
		consts.Home,
		1,
		consts.Welcome)
}
