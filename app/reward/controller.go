package reward

import (
	"github.com/ivzb/achievers_server/app"
	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/controller"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
	"github.com/ivzb/achievers_server/app/shared/validator"
)

func Last(env *app.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, err := env.DB.Reward().LastID()

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	rwds, err := env.DB.Reward().After(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Rewards,
		len(rwds),
		rwds)
}

func After(env *app.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, respErr := controller.GetFormString(env, consts.ID, env.DB.Reward())

	if respErr != nil {
		return respErr
	}

	rwds, err := env.DB.Reward().After(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Rewards,
		len(rwds),
		rwds)
}

func Single(env *app.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, respErr := controller.GetFormString(env, consts.ID, env.DB.Reward())

	if respErr != nil {
		return respErr
	}

	rwd, err := env.DB.Reward().Single(id)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Ok(
		consts.Reward,
		1,
		rwd)
}

func Create(env *app.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.POST) {
		return response.MethodNotAllowed()
	}

	rwd := &model.Reward{}
	err := form.ModelValue(env.Request, rwd)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	err = validator.Validate(*rwd)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	rewardTypeExists, err := env.DB.RewardType().Exists(rwd.RewardTypeID)

	if err != nil {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	if !rewardTypeExists {
		return response.NotFound(consts.RewardTypeID)
	}

	rwd.UserID = env.UserID

	id, err := env.DB.Reward().Create(rwd)

	if err != nil || id == "" {
		env.Log.Error(err)
		return response.InternalServerError()
	}

	return response.Created(
		consts.Reward,
		id)
}
