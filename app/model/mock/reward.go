package mock

import (
	"time"

	"github.com/ivzb/achievers_server/app/model"
)

type RewardExists struct {
	Bool bool
	Err  error
}

type RewardSingle struct {
	Rwd *model.Reward
	Err error
}

type RewardsAll struct {
	Rwds []*model.Reward
	Err  error
}

type RewardCreate struct {
	ID  string
	Err error
}

func Rewards(size int) []*model.Reward {
	rwds := make([]*model.Reward, size)

	for i := 0; i < size; i++ {
		rwds[i] = Reward()
	}

	return rwds
}

func Reward() *model.Reward {
	rwd := &model.Reward{
		"fb7691eb-ea1d-b20f-edee-9cadcf23181f",
		"name",
		"desc",
		"http://picture.jpg",
		3,
		"user_id",
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(2017, 12, 9, 15, 4, 23, 0, time.UTC),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	return rwd
}

func (mock *DB) RewardExists(string) (bool, error) {
	return mock.RewardExistsMock.Bool, mock.RewardExistsMock.Err
}

func (mock *DB) RewardSingle(id string) (*model.Reward, error) {
	return mock.RewardSingleMock.Rwd, mock.RewardSingleMock.Err
}

func (mock *DB) RewardsAll(page int) ([]*model.Reward, error) {
	return mock.RewardsAllMock.Rwds, mock.RewardsAllMock.Err
}

func (mock *DB) RewardCreate(reward *model.Reward) (string, error) {
	return mock.RewardCreateMock.ID, mock.RewardCreateMock.Err
}