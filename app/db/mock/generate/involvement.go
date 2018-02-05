package generate

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

func Involvement() interface{} {
	inv := &model.Involvement{
		1,
		"title",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return inv
}
