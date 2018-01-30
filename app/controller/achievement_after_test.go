package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func achievementsAfterForm() *map[string]string {
	return &map[string]string{
		consts.ID: mockID,
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
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(achievementsAfterForm(), consts.ID),
		db: &mock.DB{
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Err: mockDbErr},
			},
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
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Err: mockDbErr},
			},
		},
		args: achievementsAfterArgs,
	}),
	constructAchievementsAfterTest(&testInput{
		purpose:            "id AchievementExists does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.ID),
		form:               achievementsAfterForm(),
		db: &mock.DB{
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: false},
			},
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
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
				AfterMock:  mock.AchievementsAfter{Err: mockDbErr},
			},
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
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
				AfterMock:  mock.AchievementsAfter{Achs: generate.Achievements(0)},
			},
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
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
				AfterMock:  mock.AchievementsAfter{Achs: generate.Achievements(4)},
			},
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
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
				AfterMock:  mock.AchievementsAfter{Achs: generate.Achievements(9)},
			},
		},
		args: []string{"9"},
	}),
}

func constructAchievementsAfterTest(testInput *testInput) *test {
	achievementsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(generate.Achievements(achievementsSize))
	}

	return constructTest(AchievementsAfter, testInput, responseResults)
}
