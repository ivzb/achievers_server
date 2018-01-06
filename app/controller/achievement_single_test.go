package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func achievementSingleForm() *map[string]string {
	return &map[string]string{
		id: mockID,
	}
}

var achievementSingleTests = []*test{
	constructAchievementSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, id),
		form:               mapWithout(achievementSingleForm(), id),
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "achievement exists db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               achievementSingleForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Err: mockDbErr},
		},
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "achievement does not exist",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, achievement),
		form:               achievementSingleForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: false},
		},
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "achievement single db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               achievementSingleForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: true},
			AchievementSingleMock: mock.AchievementSingle{Err: mockDbErr},
		},
	}),
	constructAchievementSingleTest(&testInput{
		purpose:            "achievement single OK",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, achievement),
		form:               achievementSingleForm(),
		db: &mock.DB{
			AchievementExistsMock: mock.AchievementExists{Bool: true},
			AchievementSingleMock: mock.AchievementSingle{Ach: mock.Achievement()},
		},
	}),
}

func constructAchievementSingleTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mock.Achievement())

	return constructTest(AchievementSingle, testInput, responseResults)
}
