package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/form"
	"github.com/ivzb/achievers_server/app/shared/request"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func RewardsIndex(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, GET) {
		return response.MethodNotAllowed()
	}

	pg, err := form.IntValue(env.Request, page)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(formatInvalid, page))
	}

	rwds, err := env.DB.RewardsAll(pg)

	if err != nil {
		return response.InternalServerError()
	}

	if len(rwds) == 0 {
		return response.NotFound(page)
	}

	return response.Ok(
		rewards,
		len(rwds),
		rwds)
}

func RewardSingle(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, GET) {
		return response.MethodNotAllowed()
	}

	rwdID, err := form.StringValue(env.Request, id)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.RewardExists(rwdID)

	if err != nil {
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(reward)
	}

	rwd, err := env.DB.RewardSingle(rwdID)

	if err != nil {
		return response.InternalServerError()
	}

	return response.Ok(
		reward,
		1,
		rwd)
}

func RewardCreate(env *model.Env) *response.Message {
	if !request.IsMethod(env.Request, POST) {
		return response.MethodNotAllowed()
	}

	rwd := &model.Reward{}
	err := form.ModelValue(env.Request, rwd)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	if rwd.Title == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, title))
	}

	if rwd.Description == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, description))
	}

	if rwd.PictureURL == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, pictureURL))
	}

	if rwd.RewardTypeID == 0 {
		return response.BadRequest(fmt.Sprintf(formatMissing, rewardTypeID))
	}

	rewardTypeExists, err := env.DB.RewardTypeExists(rwd.RewardTypeID)

	if err != nil {
		return response.InternalServerError()
	}

	if !rewardTypeExists {
		return response.NotFound(rewardTypeID)
	}

	rwd.AuthorID = env.UserID

	id, err := env.DB.RewardCreate(rwd)

	if err != nil || id == "" {
		return response.InternalServerError()
	}

	return response.Created(
		reward,
		id)
}
