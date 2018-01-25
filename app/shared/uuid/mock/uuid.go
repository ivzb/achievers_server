package mock

// Generate structure used to adjust return values of Generate mock
type Generate struct {
	UUID string
	Err  error
}

// UUIDer structure implements UUIDer interface
type UUIDer struct {
	GenerateMock Generate
}

// Generate mock return mock adjusted values
func (mock *UUIDer) Generate() (string, error) {
	return mock.GenerateMock.UUID, mock.GenerateMock.Err
}
