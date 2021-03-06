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

func achievementsByQuestIDAfterForm() *map[string]string {
	return &map[string]string{
		consts.QuestID: mockID,
		consts.ID:      mockID,
	}
}

var achievementsByQuestIDAfterArgs = []string{"9"}

var achievementsByQuestIDAfterTests = []*test{
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               achievementsByQuestIDAfterArgs,
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "missing QuestID",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.QuestID),
		form:               mapWithout(achievementsByQuestIDAfterForm(), consts.QuestID),
		db:                 &mock.DB{},
		args:               achievementsByQuestIDAfterArgs,
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "id QuestExists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsByQuestIDAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Err: mockDbErr},
			},
		},
		args: achievementsByQuestIDAfterArgs,
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "id QuestExists does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.QuestID),
		form:               achievementsByQuestIDAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: false},
			},
		},
		args: achievementsByQuestIDAfterArgs,
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "missing AfterID",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(achievementsByQuestIDAfterForm(), consts.ID),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Err: mockDbErr},
			},
		},
		args: achievementsByQuestIDAfterArgs,
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "id AchievementExists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsByQuestIDAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Err: mockDbErr},
			},
		},
		args: achievementsByQuestIDAfterArgs,
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "id AchievementExists does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.ID),
		form:               achievementsByQuestIDAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: false},
			},
		},
		args: achievementsByQuestIDAfterArgs,
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "AchievementsByQuestIDAfterMock db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               achievementsByQuestIDAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock:         mock.AchievementExists{Bool: true},
				AfterByQuestIDMock: mock.AchievementsAfterByQuestID{Err: mockDbErr},
			},
		},
		args: achievementsByQuestIDAfterArgs,
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsByQuestIDAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock:         mock.AchievementExists{Bool: true},
				AfterByQuestIDMock: mock.AchievementsAfterByQuestID{Achs: generate.Achievements(0)},
			},
		},
		args: []string{"0"},
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "4 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsByQuestIDAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock:         mock.AchievementExists{Bool: true},
				AfterByQuestIDMock: mock.AchievementsAfterByQuestID{Achs: generate.Achievements(4)},
			},
		},
		args: []string{"4"},
	}),
	constructAchievementsByQuestIDAfterTest(&testInput{
		purpose:            "9 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Achievements),
		form:               achievementsByQuestIDAfterForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock:         mock.AchievementExists{Bool: true},
				AfterByQuestIDMock: mock.AchievementsAfterByQuestID{Achs: generate.Achievements(9)},
			},
		},
		args: []string{"9"},
	}),
}

func constructAchievementsByQuestIDAfterTest(testInput *testInput) *test {
	achievementsSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(generate.Achievements(achievementsSize))
	}

	return constructTest(AchievementsByQuestIDAfter, testInput, responseResults)
}
