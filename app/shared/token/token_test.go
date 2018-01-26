package token

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ivzb/achievers_server/app/shared/file"
)

func TestEnsureExists_NonExisting(t *testing.T) {
	info := &Info{
		Path: "token.pem",
		Bits: 1024,
	}

	if file.Exists(info.Path) {
		err := os.Remove(info.Path)

		if err != nil {
			t.Fatalf("is.Remove returned error: %v", err)
		}
	}

	err := info.EnsureExists()

	if err != nil {
		t.Fatalf("EnsureExists returned error: %v", err)
	}

	if !file.Exists(info.Path) {
		t.Fatalf("Token file does not exist but it should have been created")
	}

	os.Remove(info.Path)
}

func TestEnsureExists_Existing(t *testing.T) {
	info := &Info{
		Path: "token.pem",
		Bits: 1024,
	}

	if !file.Exists(info.Path) {
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

	if !file.Exists(info.Path) {
		t.Fatalf("Token file does not exist but it should have been created")
	}

	os.Remove(info.Path)
}

func TestEnsureExists_InvalidFilePath(t *testing.T) {
	info := &Info{Bits: 1024}

	err := info.EnsureExists()

	if err == nil {
		t.Fatalf("EnsureExists should return error")
	}

	os.Remove(info.Path)
}
