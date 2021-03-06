package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func questAchievementCreateForm() *map[string]string {
	return &map[string]string{
		consts.QuestID:       mockID,
		consts.AchievementID: mockID,
	}
}

var questAchievementCreateTests = []*test{
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "content-type of request is incorrect",
		form:               &map[string]string{},
		removeHeaders:      true,
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "missing form quest_id",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatBlank, consts.QuestID),
		form:               mapWithout(questAchievementCreateForm(), consts.QuestID),
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "missing form achievement_id",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatBlank, consts.AchievementID),
		form:               mapWithout(questAchievementCreateForm(), consts.AchievementID),
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questAchievementCreateForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Err: mockDbErr},
			},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_id does not exist",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.QuestID),
		form:               questAchievementCreateForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: false},
			},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "achievement exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questAchievementCreateForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Err: mockDbErr},
			},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "achievement_id does not exist",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.AchievementID),
		form:               questAchievementCreateForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: false},
			},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_achievement exists db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questAchievementCreateForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			QuestAchievementMock: mock.QuestAchievement{
				ExistsMock: mock.QuestAchievementExists{Err: mockDbErr},
			},
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
			},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_achievement already exists",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatAlreadyExists, consts.QuestAchievement),
		form:               questAchievementCreateForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
			},
			QuestAchievementMock: mock.QuestAchievement{
				ExistsMock: mock.QuestAchievementExists{Bool: true},
			},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_achievement create db error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               questAchievementCreateForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
			},
			QuestAchievementMock: mock.QuestAchievement{
				ExistsMock: mock.QuestAchievementExists{Bool: false},
				CreateMock: mock.QuestAchievementCreate{Err: mockDbErr},
			},
		},
	}),
	constructQuestAchievementCreateTest(&testInput{
		purpose:            "quest_achievement create ok",
		requestMethod:      consts.POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(consts.FormatCreated, consts.QuestAchievement),
		form:               questAchievementCreateForm(),
		db: &mock.DB{
			QuestMock: mock.Quest{
				ExistsMock: mock.QuestExists{Bool: true},
			},
			AchievementMock: mock.Achievement{
				ExistsMock: mock.AchievementExists{Bool: true},
			},
			QuestAchievementMock: mock.QuestAchievement{
				ExistsMock: mock.QuestAchievementExists{Bool: false},
				CreateMock: mock.QuestAchievementCreate{ID: mockID},
			},
		},
	}),
}

func constructQuestAchievementCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)
	return constructTest(QuestAchievementCreate, testInput, responseResults)
}
