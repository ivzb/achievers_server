package mock

// Encrypted structure used to adjust return values of Encrypt mock
type Encrypted struct {
	S string
	E error
}

// Decrypted structure used to adjust return values of Decrypt mock
type Decrypted struct {
	S string
	E error
}

// Token structure implements Tokener interface
type Token struct {
	EncryptedMock Encrypted
	DecryptedMock Decrypted
}

// Encrypt mock return mock adjusted values
func (mock *Token) Encrypt(string) (string, error) {
	return mock.EncryptedMock.S, mock.EncryptedMock.E
}

// Decrypt mock return mock adjusted values
func (mock *Token) Decrypt(string) (string, error) {
	return mock.DecryptedMock.S, mock.DecryptedMock.E
}
