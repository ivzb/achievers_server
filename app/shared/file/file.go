package file

import (
	"io"
	"io/ioutil"
	"os"
)

// Read file bytes
func Read(filePath string) ([]byte, error) {
	var err error
	input := io.ReadCloser(os.Stdin)

	if input, err = os.Open(filePath); err != nil {
		return nil, err
	}

	// Read the file
	bytes, err := ioutil.ReadAll(input)
	input.Close()

	return bytes, err
}
