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

func TestExist_Existing(t *testing.T) {
	exist := Exists("file.mock")

	if exist == false {
		t.Fatalf("Exist should return true")
	}
}

func TestExist_NonExisting(t *testing.T) {
	exist := Exists("non_existing_file.mock")

	if exist == true {
		t.Fatalf("Exist should return false")
	}
}
