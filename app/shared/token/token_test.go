package token

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ivzb/achievers_server/app/shared/file"
)

func TestEnsureExists_NonExisting(t *testing.T) {
	info := &Info{Path: "token.pem"}

	if file.Exist(info.Path) {
		err := os.Remove(info.Path)

		if err != nil {
			t.Fatalf("is.Remove returned error: %v", err)
		}
	}

	err := info.EnsureExists()

	if err != nil {
		t.Fatalf("EnsureExists returned error: %v", err)
	}

	if !file.Exist(info.Path) {
		t.Fatalf("Token file does not exist but it should have been created")
	}

	os.Remove(info.Path)
}

func TestEnsureExists_Existing(t *testing.T) {
	info := &Info{Path: "token.pem"}

	if !file.Exist(info.Path) {
		content := []byte{1, 2, 3}
		err := ioutil.WriteFile(info.Path, content, 0600)

		if err != nil {
			t.Fatalf("ioutil.WriteFile returned error: %v", err)
		}
	}

	err := info.EnsureExists()

	if err != nil {
		t.Fatalf("EnsureExists returned error: %v", err)
	}

	if !file.Exist(info.Path) {
		t.Fatalf("Token file does not exist but it should have been created")
	}

	os.Remove(info.Path)
}

func TestEnsureExists_InvalidFilePath(t *testing.T) {
	info := &Info{}

	err := info.EnsureExists()

	if err == nil {
		t.Fatalf("EnsureExists should return error")
	}

	os.Remove(info.Path)
}