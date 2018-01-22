package generate

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

func Rewards(size int) []*model.Reward {
	rwds := make([]*model.Reward, size)

	for i := 0; i < size; i++ {
		rwds[i] = Reward()
	}

	return rwds
}

func Reward() *model.Reward {
	rwd := &model.Reward{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"name",
		"desc",
		"http://picture.jpg",
		3,
		"user_id",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return rwd
}
