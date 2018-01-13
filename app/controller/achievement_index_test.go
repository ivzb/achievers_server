package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func achievementsIndexForm() *map[string]string {
	return &map[string]string{
		consts.Page: mockPage,
	}
}

var achievementsIndexArgs = []string{"9"}

var achievementsIndexTests = []*test{
	constructAchievementsIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               achievementsIndexArgs,
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "missing form consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Page),
		form:               mapWithout(achievementsIndexForm(), consts.Page),
		args:               achievementsIndexArgs,
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "invalid consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatInvalid, consts.Page),
		form: &map[string]string{
			consts.Page: "-1",
		},
		args: achievementsIndexArgs,
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsIndexForm(),
		db: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{Err: mockDbErr},
		},
		args: achievementsIndexArgs,
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "no results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Page),
		form:               achievementsIndexForm(),
		db: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{Achs: mock.Achievements(0)},
		},
		args: []string{"0"},
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "4 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsIndexForm(),
		db: &mock.DB{
			AchievementsAllMock: mock.AchievementsAll{Achs: mock.Achievements(4)},
		},
		args: []string{"4"},
	}),
	constructAchievementsIndexTest(&testInput{
		purpose:            "9 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
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
