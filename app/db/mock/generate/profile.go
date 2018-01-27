package generate

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

func Profile() *model.Profile {
	prfl := &model.Profile{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"Ivan Zahariev",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return prfl
}