package controller

import (
	"fmt"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func RewardsIndex(env *model.Env) *response.Message {
	if !env.Request.IsMethod(GET) {
		return response.MethodNotAllowed()
	}

	pg, err := env.Request.Form.IntValue(page)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	rwds, err := env.DB.RewardsAll(pg)

	if err != nil {
		return response.InternalServerError()
	}

	if len(rwds) == 0 {
		return response.NotFound(fmt.Sprintf(formatNotFound, page))
	}

	return response.Ok(
		fmt.Sprintf(formatFound, rewards),
		len(rwds),
		rwds)
}

func RewardSingle(env *model.Env) *response.Message {
	if !env.Request.IsMethod(GET) {
		return response.MethodNotAllowed()
	}

	rwdID, err := env.Request.Form.StringValue(id)

	if err != nil {
		return response.BadRequest(err.Error())
	}

	exists, err := env.DB.RewardExists(rwdID)

	if err != nil {
		return response.InternalServerError()
	}

	if !exists {
		return response.NotFound(fmt.Sprintf(formatNotFound, reward))
	}

	rwd, err := env.DB.RewardSingle(rwdID)

	if err != nil {
		return response.InternalServerError()
	}

	return response.Ok(
		fmt.Sprintf(formatFound, reward),
		1,
		rwd)
}

func RewardCreate(env *model.Env) *response.Message {
	if !env.Request.IsMethod(POST) {
		return response.MethodNotAllowed()
	}

	rwd := &model.Reward{}
	err := env.Request.Form.Map(rwd)

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
		return response.NotFound(fmt.Sprintf(formatNotFound, rewardTypeID))
	}

	rwd.AuthorID = env.Request.UserID

	id, err := env.DB.RewardCreate(rwd)

	if err != nil || id == "" {
		return response.InternalServerError()
	}

	return response.Ok(
		fmt.Sprintf(formatCreated, reward),
		1,
		id)
}
