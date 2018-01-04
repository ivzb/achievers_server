package mock

type InvolvementExists struct {
	Bool bool
	Err  error
}

func (mock *DB) InvolvementExists(string) (bool, error) {
	return mock.InvolvementExistsMock.Bool, mock.InvolvementExistsMock.Err
}
