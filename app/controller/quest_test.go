package controller

import "testing"

func TestQuestsAfter(t *testing.T) {
	run(t, questsAfterTests)
}

func TestQuestsLatest(t *testing.T) {
	run(t, questsLatestTests)
}

func TestQuestSingle(t *testing.T) {
	run(t, questSingleTests)
}

func TestQuestCreate(t *testing.T) {
	run(t, questCreateTests)
}
