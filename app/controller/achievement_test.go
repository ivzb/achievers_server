package controller

import (
	"testing"
)

func TestAchievementsAfter(t *testing.T) {
	run(t, achievementsAfterTests)
}

func TestAchievementsLatest(t *testing.T) {
	run(t, achievementsLatestTests)
}

func TestAchievementsByQuestIDAfter(t *testing.T) {
	run(t, achievementsByQuestIDAfterTests)
}

func TestAchievementsByQuestIDLast(t *testing.T) {
	run(t, achievementsByQuestIDLastTests)
}

func TestAchievementSingle(t *testing.T) {
	run(t, achievementSingleTests)
}

func TestAchievementCreate(t *testing.T) {
	run(t, achievementCreateTests)
}
