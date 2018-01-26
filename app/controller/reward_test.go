package controller

import "testing"

func TestRewardsAfter(t *testing.T) {
	run(t, rewardsAfterTests)
}

func TestRewardsLatest(t *testing.T) {
	run(t, rewardsLatestTests)
}

func TestRewardSingle(t *testing.T) {
	run(t, rewardSingleTests)
}

func TestRewardCreate(t *testing.T) {
	run(t, rewardCreateTests)
}
