package controller

import (
	"fmt"
	"path/filepath"

	"github.com/ivzb/achievers_server/app/model"
	f "github.com/ivzb/achievers_server/app/shared/file"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func FileSingle(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, GET) {
		return response.MethodNotAllowed()
	}

	id, err := form.StringValue(env.Request, id)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	path := fmt.Sprintf("%s/%s.jpg", env.Config.Server.FileStorage, id)
	exists := f.Exists(path)

	if !exists {
		return response.NotFound(id)
	}

	return response.File(file, path)
}

func FileCreate(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, POST) {
		return response.MethodNotAllowed()
	}

	multipart, header, err := form.MultipartFile(env.Request, "file")

	if err != nil {
		env.Log.Error(err)
		return response.BadRequest(err.Error())
	}

	filename, err := env.DB.UUID()

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	ext := filepath.Ext(header.Filename)

	path := fmt.Sprintf("%s/%s%s", env.Config.Server.FileStorage, filename, ext)
	err = f.Create(path, multipart)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Created(file, filename)
}
