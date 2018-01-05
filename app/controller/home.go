package controller

import (
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func HomeIndex(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed(methodNotAllowed)
	}

	return response.Ok(
		fmt.Sprintf(formatFound, home),
		1,
		welcome)
}
