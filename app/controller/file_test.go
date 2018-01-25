package controller

import (
	"testing"
)

func TestFileSingle(t *testing.T) {
	run(t, fileSingleTests)
}

func TestFileCreate(t *testing.T) {
	run(t, fileCreateTests)
}
