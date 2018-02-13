package reward

import (
	"testing"

	"github.com/ivzb/achievers_server/app/shared/test"
)

func TestRewardsAfter(t *testing.T) {
	test.Run(t, rewardsAfterTests)
}

func TestRewardsLatest(t *testing.T) {
	test.Run(t, rewardsLatestTests)
}

func TestRewardSingle(t *testing.T) {
	test.Run(t, rewardSingleTests)
}

func TestRewardCreate(t *testing.T) {
	test.Run(t, rewardCreateTests)
}
