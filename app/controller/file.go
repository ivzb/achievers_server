package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/file"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func FileSingle(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, GET) {
		return response.MethodNotAllowed()
	}

	filename, err := form.StringValue(env.Request, id)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	path := fmt.Sprintf("%s/%s", env.Config.Server.FileStorage, filename)
	exists := file.Exists(path)

	if !exists {
		return response.NotFound(id)
	}

	return response.File(path)
}

func FileCreate(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, POST) {
		return response.MethodNotAllowed()
	}

	multipart, _, err := form.MultipartFile(env.Request, "file")

	if err != nil {
		env.Log.Error(err)
		return response.BadRequest(err.Error())
	}

	filename, err := env.DB.UUID()

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	path := fmt.Sprintf("%s/%s", env.Config.Server.FileStorage, filename)
	err = file.Create(path, multipart)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Created("file", filename)
}
