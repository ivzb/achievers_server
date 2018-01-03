package controller

import (
	"testing"
)

func TestAchievementsIndex(t *testing.T) {
	run(t, achievementsIndexTests)
}

func TestAchievementSingle(t *testing.T) {
	run(t, achievementSingleTests)
}

func TestAchievementCreate(t *testing.T) {
	run(t, achievementCreateTests)
}
