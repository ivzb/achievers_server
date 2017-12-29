package mock

// Encrypted structure used to adjust return values of Encrypt mock
type Encrypted struct {
	Enc string
	Err error
}

// Decrypted structure used to adjust return values of Decrypt mock
type Decrypted struct {
	Dec string
	Err error
}

// Token structure implements Tokener interface
type Token struct {
	EncryptedMock Encrypted
	DecryptedMock Decrypted
}

// Encrypt mock return mock adjusted values
func (mock *Token) Encrypt(string) (string, error) {
	return mock.EncryptedMock.Enc, mock.EncryptedMock.Err
}

// Decrypt mock return mock adjusted values
func (mock *Token) Decrypt(string) (string, error) {
	return mock.DecryptedMock.Dec, mock.DecryptedMock.Err
}
