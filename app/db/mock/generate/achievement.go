package generate

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

func Achievements(size int) []*model.Achievement {
	achs := make([]*model.Achievement, size)

	for i := 0; i < size; i++ {
		achs[i] = Achievement()
	}

	return achs
}

func Achievement() *model.Achievement {
	ach := &model.Achievement{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"title",
		"desc",
		"http://picture.jpg",
		"3",
		"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return ach
}
