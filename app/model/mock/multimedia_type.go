package mock

type MultimediaTypeExists struct {
	Bool bool
	Err  error
}

func (mock *DB) MultimediaTypeExists(uint8) (bool, error) {
	return mock.MultimediaTypeExistsMock.Bool, mock.MultimediaTypeExistsMock.Err
}
