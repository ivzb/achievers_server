package file

import (
	"io"
	"io/ioutil"
	"os"
)

// Read file bytes
func Read(path string) ([]byte, error) {
	var err error
	input := io.ReadCloser(os.Stdin)

	if input, err = os.Open(path); err != nil {
		return nil, err
	}

	// Read the file
	bytes, err := ioutil.ReadAll(input)
	input.Close()

	return bytes, err
}

// Exist file
func Exist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
