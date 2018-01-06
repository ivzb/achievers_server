package controller

import (
	"net/http"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func HomeIndex(
	env *model.Env,
	r *http.Request) *response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed()
	}

	return response.Ok(
		home,
		1,
		welcome)
}
