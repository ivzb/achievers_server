package controller

import (
	"testing"
)

func TestAchievementsIndex(t *testing.T) {
	for _, test := range achievementsIndexTests {
		rec := constructRequest(t, test)
		expect(t, rec, test)
	}
}

func TestAchievementSingle(t *testing.T) {
	for _, test := range achievementSingleTests {
		rec := constructRequest(t, test)
		expect(t, rec, test)
	}
}

func TestAchievementCreate(t *testing.T) {
	for _, test := range achievementCreateTests {
		rec := constructRequest(t, test)
		expect(t, rec, test)
	}
}
