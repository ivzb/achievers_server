package generate

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

func Quests(size int) []interface{} {
	qsts := make([]interface{}, size)

	for i := 0; i < size; i++ {
		qsts[i] = Quest()
	}

	return qsts
}

func Quest() interface{} {
	qst := &model.Quest{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"name",
		"http://picture.jpg",
		3,
		3,
		"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return qst
}
