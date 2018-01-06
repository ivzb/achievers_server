package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model"
	"github.com/ivzb/achievers_server/app/shared/response"
)

func RewardsIndex(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) *response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed()
	}

	pg, err := strconv.Atoi(r.FormValue(page))

	if err != nil {
		return response.BadRequest(fmt.Sprintf(formatMissing, page))
	}

	if pg < 0 {
		return response.BadRequest(fmt.Sprintf(formatInvalid, page))
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

func RewardSingle(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) *response.Message {

	if r.Method != "GET" {
		return response.MethodNotAllowed()
	}

	rwdID := r.FormValue(id)

	if rwdID == "" {
		return response.BadRequest(fmt.Sprintf(formatMissing, id))
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

func RewardCreate(
	env *model.Env,
	w http.ResponseWriter,
	r *http.Request) *response.Message {

	if r.Method != "POST" {
		return response.MethodNotAllowed()
	}

	rwd := &model.Reward{}
	err := env.Former.Map(r, rwd)

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

	rwd.AuthorID = env.UserId

	id, err := env.DB.RewardCreate(rwd)

	if err != nil || id == "" {
		return response.InternalServerError()
	}

	return response.Ok(
		fmt.Sprintf(formatCreated, reward),
		1,
		id)
}
