package generate

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

func QuestAchievements(size int) []interface{} {
	qstAchs := make([]interface{}, size)

	for i := 0; i < size; i++ {
		qstAchs[i] = QuestAchievement()
	}

	return qstAchs
}

func QuestAchievement() interface{} {
	qstAch := &model.QuestAchievement{
		"mock_id",
		"fd9d70c9-f3aa-8f53-0946-b889a57ef22d",
		"a247c0d9-53bb-e6bf-4828-5ad5804570c4",
		"4e69c9ba-719c-ca7c-fb66-80ab235c8e39",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return qstAch
}
