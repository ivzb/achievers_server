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

func RewardsIndex(env *env.Env) *response.Message {
	if !request.IsMethod(env.Request, consts.GET) {
		return response.MethodNotAllowed()
	}

	pg, err := form.IntValue(env.Request, consts.Page)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(consts.FormatInvalid, consts.Page))
	}

	rwds, err := env.DB.Reward().All(pg)

	if err != nil {
		return response.InternalServerError()
	}

	if len(rwds) == 0 {
		return response.NotFound(consts.Page)
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

	rwdID, err := form.StringValue(env.Request, consts.ID)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.Reward().Exists(rwdID)

	if err != nil {
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(consts.Reward)
	}

	rwd, err := env.DB.Reward().Single(rwdID)

	if err != nil {
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
		return response.InternalServerError()
	}

	if !rewardTypeExists {
		return response.NotFound(consts.RewardTypeID)
	}

	rwd.UserID = env.UserID

	id, err := env.DB.Reward().Create(rwd)

	if err != nil || id == "" {
		return response.InternalServerError()
	}

	return response.Created(
		consts.Reward,
		id)
}
