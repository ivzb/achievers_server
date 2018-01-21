package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func achievementsAfterForm() *map[string]string {
	return &map[string]string{
		consts.AfterID: mockID,
	}
}

var achievementsAfterArgs = []string{"9"}

var achievementsAfterTests = []*test{
	constructAchievementsAfterTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               achievementsAfterArgs,
	}),
	constructAchievementsAfterTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.AfterID),
		form:               mapWithout(achievementsAfterForm(), consts.AfterID),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Err: mockDbErr},
		},
		args: achievementsAfterArgs,
	}),
	constructAchievementsAfterTest(&testInput{
		purpose:            "id AchievementExists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsAfterForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Err: mockDbErr},
		},
		args: achievementsAfterArgs,
	}),
	constructAchievementsAfterTest(&testInput{
		purpose:            "id AchievementExists does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.AfterID),
		form:               achievementsAfterForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: false},
		},
		args: achievementsAfterArgs,
	}),
	constructAchievementsAfterTest(&testInput{
		purpose:            "AchievementsAfterMock db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsAfterForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: true},
			AchievementsAfterMock: mock.AchievementsAfter{Err: mockDbErr},
		},
		args: achievementsAfterArgs,
	}),
	constructAchievementsAfterTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsAfterForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: true},
			AchievementsAfterMock: mock.AchievementsAfter{Achs: mock.Achievements(0)},
		},
		args: []string{"0"},
	}),
	constructAchievementsAfterTest(&testInput{
		purpose:            "4 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsAfterForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: true},
			AchievementsAfterMock: mock.AchievementsAfter{Achs: mock.Achievements(4)},
		},
		args: []string{"4"},
	}),
	constructAchievementsAfterTest(&testInput{
		purpose:            "9 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsAfterForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: true},
			AchievementsAfterMock: mock.AchievementsAfter{Achs: mock.Achievements(9)},
		},
		args: []string{"9"},
	}),
}

func constructAchievementsAfterTest(testInput *testInput) *test {
	achievementsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(mock.Achievements(achievementsSize))
	}

	return constructTest(AchievementsAfter, testInput, responseResults)
}
