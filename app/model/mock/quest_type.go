package mock

type QuestTypeExists struct {
	Bool bool
	Err  error
}

func (mock *DB) QuestTypeExists(uint8) (bool, error) {
	return mock.QuestTypeExistsMock.Bool, mock.QuestTypeExistsMock.Err
}
