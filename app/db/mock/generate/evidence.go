package generate

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

func Evidences(size int) []interface{} {
	evds := make([]interface{}, size)

	for i := 0; i < size; i++ {
		evds[i] = Evidence()
	}

	return evds
}

func Evidence() interface{} {
	evd := &model.Evidence{
		"mock_id",
		"desc",
		"http://preview-url.jpg",
		"http://url.jpg",
		3,
		"65hjl4aa-719c-ca7c-fb66-80ab235c8e39",
		"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return evd
}
