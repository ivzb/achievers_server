package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func achievementsByQuestIDForm() *map[string]string {
	return &map[string]string{
		consts.Page: mockPage,
		consts.ID:   mockID,
	}
}

var achievementsByQuestIDArgs = []string{"9"}

var achievementsByQuestIDTests = []*test{
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "invalconsts.ID request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "missing consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Page),
		form:               mapWithout(achievementsByQuestIDForm(), consts.Page),
		args:               achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "invalconsts.ID consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatInvalid, consts.Page),
		form: &map[string]string{
			consts.Page: "-1",
		},
		args: achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "missing consts.ID",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(achievementsByQuestIDForm(), consts.ID),
		args:               achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "questExists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Err: mockDbErr},
		},
		args: achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "consts.ID does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.ID),
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: false},
		},
		args: achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock:           mock.QuestExists{Bool: true},
			AchievementsByQuestIDMock: mock.AchievementsByQuestID{Err: mockDbErr},
		},
		args: achievementsByQuestIDArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "no results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Page),
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock:           mock.QuestExists{Bool: true},
			AchievementsByQuestIDMock: mock.AchievementsByQuestID{Achs: mock.Achievements(0)},
		},
		args: []string{"0"},
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "4 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsByQuestIDForm(),
		db: &mock.DB{
			QuestExistsMock:           mock.QuestExists{Bool: true},
			AchievementsByQuestIDMock: mock.AchievementsByQuestID{Achs: mock.Achievements(4)},
		},
		args: []string{"4"},
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "9 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
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
