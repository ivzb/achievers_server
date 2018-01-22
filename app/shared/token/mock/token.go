package mock

// Encrypt structure used to adjust return values of Encrypt mock
type Encrypt struct {
	Enc string
	Err error
}

// Decrypt structure used to adjust return values of Decrypt mock
type Decrypt struct {
	Dec string
	Err error
}

// Tokener structure implements Tokener interface
type Tokener struct {
	EncryptMock Encrypt
	DecryptMock Decrypt
}

// Encrypt mock return mock adjusted values
func (mock *Tokener) Encrypt(string) (string, error) {
	return mock.EncryptMock.Enc, mock.EncryptMock.Err
}

// Decrypt mock return mock adjusted values
func (mock *Tokener) Decrypt(string) (string, error) {
	return mock.DecryptMock.Dec, mock.DecryptMock.Err
}
