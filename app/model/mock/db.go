package mock

type DB struct {
	UserExistsMock      UserExists
	UserEmailExistsMock UserEmailExists
	UserCreateMock      UserCreate
	UserAuthMock        UserAuth

	AchievementExistsMock AchievementExists
	AchievementSingleMock AchievementSingle
	AchievementsAllMock   AchievementsAll
	AchievementCreateMock AchievementCreate

	EvidenceExistsMock EvidenceExists
	EvidenceSingleMock EvidenceSingle
	EvidencesAllMock   EvidencesAll
	EvidenceCreateMock EvidenceCreate

	RewardExistsMock RewardExists
	RewardSingleMock RewardSingle
	RewardsAllMock   RewardsAll
	RewardCreateMock RewardCreate

	RewardTypeExistsMock RewardTypeExists

	QuestExistsMock QuestExists
	QuestSingleMock QuestSingle
	QuestsAllMock   QuestsAll
	QuestCreateMock QuestCreate

	QuestTypeExistsMock QuestTypeExists

	QuestAchievementExistsMock QuestAchievementExists
	QuestAchievementSingleMock QuestAchievementSingle

	InvolvementExistsMock InvolvementExists

	MultimediaTypeExistsMock MultimediaTypeExists
}
