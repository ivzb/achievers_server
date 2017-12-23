package file

import (
	"reflect"
	"testing"
)

func TestRead_ExistingFile(t *testing.T) {
	expected := []byte{102, 105, 108, 101, 32, 109, 111, 99, 107}

	actual, err := Read("file.mock")

	if err != nil {
		t.Fatalf("Read returned error: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Read returned wrong value: expected %v, actual %v",
			expected, actual)
	}
}

func TestRead_NonExistingFile(t *testing.T) {
	_, err := Read("non_existing_file.mock")

	if err == nil {
		t.Fatalf("Read should returned non existing error")
	}
}
