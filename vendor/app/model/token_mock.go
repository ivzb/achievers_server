package model

type EncryptedMock struct {
	S string
	E error
}

type DecryptedMock struct {
	S string
	E error
}

type TokenMock struct {
	EncryptedMock EncryptedMock
	DecryptedMock DecryptedMock
}

func (mock *TokenMock) Encrypt(string) (string, error) {
	return mock.EncryptedMock.S, mock.EncryptedMock.E
}

func (mock *TokenMock) Decrypt(string) (string, error) {
	return mock.DecryptedMock.S, mock.DecryptedMock.E
}
