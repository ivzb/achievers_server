package mock

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
