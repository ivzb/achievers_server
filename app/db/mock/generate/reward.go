package generate

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

func Rewards(size int) []interface{} {
	rwds := make([]interface{}, size)

	for i := 0; i < size; i++ {
		rwds[i] = Reward()
	}

	return rwds
}

func Reward() interface{} {
	rwd := &model.Reward{
		"mock_id",
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
