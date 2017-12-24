package model

// EncryptedMock structure used to adjust return values of Encrypt mock
type EncryptedMock struct {
	S string
	E error
}

// DecryptedMock structure used to adjust return values of Decrypt mock
type DecryptedMock struct {
	S string
	E error
}

// TokenMock structure implements Tokener interface
type TokenMock struct {
	EncryptedMock EncryptedMock
	DecryptedMock DecryptedMock
}

// Encrypt mock return mock adjusted values
func (mock *TokenMock) Encrypt(string) (string, error) {
	return mock.EncryptedMock.S, mock.EncryptedMock.E
}

// Decrypt mock return mock adjusted values
func (mock *TokenMock) Decrypt(string) (string, error) {
	return mock.DecryptedMock.S, mock.DecryptedMock.E
}
