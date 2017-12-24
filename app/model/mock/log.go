package mock

type Log struct {
	E error
}

type Logger struct {
	LogMock Log
}

func (mock *Logger) Log(string) error {
	return mock.LogMock.E
}
