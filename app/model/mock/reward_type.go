package mock

type RewardTypeExists struct {
	Bool bool
	Err  error
}

func (mock *DB) RewardTypeExists(uint8) (bool, error) {
	return mock.RewardTypeExistsMock.Bool, mock.RewardTypeExistsMock.Err
}
