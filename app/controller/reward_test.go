package controller

import (
	"testing"
)

func TestRewardsIndex(t *testing.T) {
	run(t, rewardsIndexTests)
}

func TestRewardSingle(t *testing.T) {
	run(t, rewardSingleTests)
}
