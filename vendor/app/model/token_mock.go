package model

type EncryptedMock struct {
	S string
	E error
}

type DecryptedMock struct {
	S string
	E error
}

type TokenerMock struct {
	EncryptedMock EncryptedMock
	DecryptedMock DecryptedMock
}

func (mock *TokenerMock) Encrypt(string) (string, error) {
	return mock.EncryptedMock.S, mock.EncryptedMock.E
}

func (mock *TokenerMock) Decrypt(string) (string, error) {
	return mock.DecryptedMock.S, mock.DecryptedMock.E
}
