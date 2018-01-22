package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func achievementSingleForm() *map[string]string {
	return &map[string]string{
		consts.ID: mockID,
	}
}

var achievementSingleTests = []*test{
	constructAchievementSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(achievementSingleForm(), consts.ID),
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "achievement exists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementSingleForm(),
		db: &mock.DB{
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Err: mockDbErr},
			},
		},
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "achievement does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Achievement),
		form:               achievementSingleForm(),
		db: &mock.DB{
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: false},
			},
		},
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "achievement single db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementSingleForm(),
		db: &mock.DB{
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
				SingleMock: mock.AchievementSingle{Err: mockDbErr},
			},
		},
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "achievement single OK",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievement),
		form:               achievementSingleForm(),
		db: &mock.DB{
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
				SingleMock: mock.AchievementSingle{Ach: generate.Achievement()},
			},
		},
	}),
}

func constructAchievementSingleTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(generate.Achievement())

	return constructTest(AchievementSingle, testInput, responseResults)
}
