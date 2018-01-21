package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func achievementsByQuestIDLastForm() *map[string]string {
	return &map[string]string{
		consts.QuestID: mockID,
	}
}

var achievementsByQuestIDLastArgs = []string{"9"}

var achievementsByQuestIDLastTests = []*test{
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "invalconsts.ID request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               achievementsByQuestIDLastArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "missing QuestID",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.QuestID),
		form:               mapWithout(achievementsByQuestIDLastForm(), consts.QuestID),
		args:               achievementsByQuestIDLastArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "questExists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsByQuestIDLastForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Err: mockDbErr},
		},
		args: achievementsByQuestIDLastArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "QuestID does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.QuestID),
		form:               achievementsByQuestIDLastForm(),
		db: &mock.DB{
			QuestExistsMock: mock.QuestExists{Bool: false},
		},
		args: achievementsByQuestIDLastArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "AchievementsByQuestIDLastID db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsByQuestIDLastForm(),
		db: &mock.DB{
			QuestExistsMock:                 mock.QuestExists{Bool: true},
			AchievementsByQuestIDLastIDMock: mock.AchievementsByQuestIDLastID{Err: mockDbErr},
		},
		args: achievementsByQuestIDLastArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "AchievementsByQuestIDAfter db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsByQuestIDLastForm(),
		db: &mock.DB{
			QuestExistsMock:                mock.QuestExists{Bool: true},
			AchievementsByQuestIDAfterMock: mock.AchievementsByQuestIDAfter{Err: mockDbErr},
		},
		args: achievementsByQuestIDLastArgs,
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsByQuestIDLastForm(),
		db: &mock.DB{
			QuestExistsMock:                mock.QuestExists{Bool: true},
			AchievementsByQuestIDAfterMock: mock.AchievementsByQuestIDAfter{Achs: mock.Achievements(0)},
		},
		args: []string{"0"},
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "4 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsByQuestIDLastForm(),
		db: &mock.DB{
			QuestExistsMock:                mock.QuestExists{Bool: true},
			AchievementsByQuestIDAfterMock: mock.AchievementsByQuestIDAfter{Achs: mock.Achievements(4)},
		},
		args: []string{"4"},
	}),
	constructAchievementsByQuestIDTest(&testInput{
		purpose:            "9 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsByQuestIDLastForm(),
		db: &mock.DB{
			QuestExistsMock:                mock.QuestExists{Bool: true},
			AchievementsByQuestIDAfterMock: mock.AchievementsByQuestIDAfter{Achs: mock.Achievements(9)},
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

	return constructTest(AchievementsByQuestIDLast, testInput, responseResults)
}
