package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func achievementsByQuestIDForm() *map[string]string {
	return &map[string]string{
		page: mockPage,
		id:   mockID,
	}
}

var achievementsByQuestIDArgs = []string{"9"}

var achievementsByQuestIDTests = []*test{
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
		args:               achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "missing page",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, page),
		form:               mapWithout(achievementsByQuestIDForm(), page),
		args:               achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "invalid page",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatInvalid, page),
		form: &map[string]string{
			page: "-1",
		},
		args: achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "missing id",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, id),
		form:               mapWithout(achievementsByQuestIDForm(), id),
		args:               achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "questExists db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Err: mockDbErr},
		},
		args: achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "id does not exist",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, id),
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: false},
		},
		args: achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "db error",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock:           mock.QuestExists{Bool: true},
			AchievementsByQuestIDMock: mock.AchievementsByQuestID{Err: mockDbErr},
		},
		args: achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "no results on page",
		requestMethod:      GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, page),
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock:           mock.QuestExists{Bool: true},
			AchievementsByQuestIDMock: mock.AchievementsByQuestID{Achs: mock.Achievements(0)},
		},
		args: []string{"0"},
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "4 results on page",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, achievements),
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock:           mock.QuestExists{Bool: true},
			AchievementsByQuestIDMock: mock.AchievementsByQuestID{Achs: mock.Achievements(4)},
		},
		args: []string{"4"},
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "9 results on page",
		requestMethod:      GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, achievements),
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock:           mock.QuestExists{Bool: true},
			AchievementsByQuestIDMock: mock.AchievementsByQuestID{Achs: mock.Achievements(9)},
		},
		args: []string{"9"},
	}),
}

func constructAchievementsByQuestIDTest(testInput *testInput) *test {
	achievementsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(mock.Achievements(achievementsSize))
	}

	return constructTest(AchievementsByQuestID, testInput, responseResults)
}
