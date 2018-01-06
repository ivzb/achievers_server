package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func achievementsIndexForm() *map[string]string {
	return &map[string]string{
		page: mockPage,
	}
}

var achievementsIndexArgs = []string{"9"}

var achievementsIndexTests = []*test{
	constructAchievementsIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
		args:               achievementsIndexArgs,
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "missing page",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, page),
		form:               mapWithout(achievementsIndexForm(), page),
		args:               achievementsIndexArgs,
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "invalid page",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatInvalid, page),
		form: &map[string]string{
			page: "-1",
		},
		args: achievementsIndexArgs,
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               achievementsIndexForm(),
		db: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{Err: mockDbErr},
		},
		args: achievementsIndexArgs,
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "no results on page",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, page),
		form:               achievementsIndexForm(),
		db: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{Achs: mock.Achievements(0)},
		},
		args: []string{"0"},
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "4 results on page",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, achievements),
		form:               achievementsIndexForm(),
		db: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{Achs: mock.Achievements(4)},
		},
		args: []string{"4"},
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "9 results on page",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, achievements),
		form:               achievementsIndexForm(),
		db: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{Achs: mock.Achievements(9)},
		},
		args: []string{"9"},
	}),
}

func constructAchievementsIndexTest(testInput *testInput) *test {
	achievementsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(mock.Achievements(achievementsSize))
	}

	return constructTest(AchievementsIndex, testInput, responseResults)
}
