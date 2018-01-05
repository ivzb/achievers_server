package mock

type DB struct {
	UserExistsMock      UserExists
	UserEmailExistsMock UserEmailExists
	UserCreateMock      UserCreate
	UserAuthMock        UserAuth

	AchievementExistsMock     AchievementExists
	AchievementSingleMock     AchievementSingle
	AchievementsAllMock       AchievementsAll
	AchievementsByQuestIDMock AchievementsByQuestID
	AchievementCreateMock     AchievementCreate

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
	QuestAchievementCreateMock QuestAchievementCreate

	InvolvementExistsMock InvolvementExists

	MultimediaTypeExistsMock MultimediaTypeExists
}
