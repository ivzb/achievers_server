package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/env"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func RewardsLast(env *env.Env) *response.Message {
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

func RewardsAfter(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, respErr := getFormString(env, consts.ID, env.DB.Reward())

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

func RewardSingle(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	id, respErr := getFormString(env, consts.ID, env.DB.Reward())

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

func RewardCreate(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.POST) {
		return response.MethodNotAllowed()
	}

	rwd := &model.Reward{}
	err := form.ModelValue(env.Request, rwd)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if rwd.Title == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Title))
	}

	if rwd.Description == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.Description))
	}

	if rwd.PictureURL == "" {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.PictureURL))
	}

	if rwd.RewardTypeID == 0 {
		return response.BadRequest(fmt.Sprintf(consts.FormatMissing, consts.RewardTypeID))
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
