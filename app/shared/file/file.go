package file

import (
	"io"
	"io/ioutil"
	"mime/multipart"
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

// Exists file
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// Create a file
func Create(path string, file multipart.File) error {
	defer file.Close()

	out, err := os.Create(path)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	return err
}
