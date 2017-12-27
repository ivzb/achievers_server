package conv

import (
	"reflect"
	"strconv"
	"testing"
)

type mock struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func TestSafe_Int(t *testing.T) {
	var expected int = 5
	var actual int

	refl := reflect.ValueOf(&actual)
	err := Safe(strconv.Itoa(expected), refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if 5 != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_IntError(t *testing.T) {
	var expected int = 0
	var actual int

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if err == nil {
		t.Error("Map should have returned convertion error")
	}

	if 0 != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}
