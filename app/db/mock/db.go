package mock

import "github.com/ivzb/achievers_server/app/db"

type DB struct {
	UUIDMock UUID

	UserMock             User
	ProfileMock          Profile
	AchievementMock      Achievement
	EvidenceMock         Evidence
	RewardMock           Reward
	RewardTypeMock       RewardType
	QuestMock            Quest
	QuestTypeMock        QuestType
	QuestAchievementMock QuestAchievement
	InvolvementMock      Involvement
	MultimediaTypeMock   MultimediaType
}

type UUID struct {
	UUID string
	Err  error
}

func (mock *DB) UUID() (string, error) {
	return mock.UUIDMock.UUID, mock.UUIDMock.Err
}

func (mock *DB) Achievement() db.Achievementer {
	return &mock.AchievementMock
}

func (mock *DB) Evidence() db.Evidencer {
	return &mock.EvidenceMock
}

func (mock *DB) Involvement() db.Involvementer {
	return &mock.InvolvementMock
}

func (mock *DB) MultimediaType() db.MultimediaTyper {
	return &mock.MultimediaTypeMock
}

func (mock *DB) Profile() db.Profiler {
	return &mock.ProfileMock
}

func (mock *DB) Quest() db.Quester {
	return &mock.QuestMock
}

func (mock *DB) QuestAchievement() db.QuestAchievementer {
	return &mock.QuestAchievementMock
}

func (mock *DB) QuestType() db.QuestTyper {
	return &mock.QuestTypeMock
}

func (mock *DB) Reward() db.Rewarder {
	return &mock.RewardMock
}

func (mock *DB) RewardType() db.RewardTyper {
	return &mock.RewardTypeMock
}

func (mock *DB) User() db.Userer {
	return &mock.UserMock
}
