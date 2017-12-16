package model

type LogMock struct {
	E error
}

type LoggerMock struct {
	LogMock LogMock
}

func (mock *LoggerMock) Log(string) error {
	return mock.LogMock.E
}
