package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

var achievementsLatestArgs = []string{"9"}

var achievementsLatestTests = []*test{
	constructAchievementsLatestTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               achievementsLatestArgs,
	}),
	constructAchievementsLatestTest(&testInput{
		purpose:            "id AchievementsLastID db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		db: &mock.DB{
			AchievementsLastIDMock: mock.AchievementsLastID{Err: mockDbErr},
		},
		args: achievementsLatestArgs,
	}),
	constructAchievementsLatestTest(&testInput{
		purpose:            "AchievementsAfterMock db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		db: &mock.DB{
			AchievementsLastIDMock: mock.AchievementsLastID{ID: mockID},
			AchievementsAfterMock:  mock.AchievementsAfter{Err: mockDbErr},
		},
		args: achievementsLatestArgs,
	}),
	constructAchievementsLatestTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		db: &mock.DB{
			AchievementsLastIDMock: mock.AchievementsLastID{ID: mockID},
			AchievementsAfterMock:  mock.AchievementsAfter{Achs: mock.Achievements(0)},
		},
		args: []string{"0"},
	}),
	constructAchievementsLatestTest(&testInput{
		purpose:            "4 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		db: &mock.DB{
			AchievementsLastIDMock: mock.AchievementsLastID{ID: mockID},
			AchievementsAfterMock:  mock.AchievementsAfter{Achs: mock.Achievements(4)},
		},
		args: []string{"4"},
	}),
	constructAchievementsLatestTest(&testInput{
		purpose:            "9 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		db: &mock.DB{
			AchievementsLastIDMock: mock.AchievementsLastID{ID: mockID},
			AchievementsAfterMock:  mock.AchievementsAfter{Achs: mock.Achievements(9)},
		},
		args: []string{"9"},
	}),
}

func constructAchievementsLatestTest(testInput *testInput) *test {
	achievementsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(mock.Achievements(achievementsSize))
	}

	return constructTest(AchievementsLast, testInput, responseResults)
}
