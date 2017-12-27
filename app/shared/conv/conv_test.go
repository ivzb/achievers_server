package conv

import (
	"reflect"
	"testing"
	"time"
)

type mock struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func TestSafe_String(t *testing.T) {
	expected := "mock"
	var actual string

	refl := reflect.ValueOf(&actual)
	err := Safe("mock", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Bool(t *testing.T) {
	expected := true
	var actual bool

	refl := reflect.ValueOf(&actual)
	err := Safe("true", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_BoolError(t *testing.T) {
	var expected bool
	var actual bool

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Float32(t *testing.T) {
	var expected float32 = 5.431
	var actual float32

	refl := reflect.ValueOf(&actual)
	err := Safe("5.431", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Float32Error(t *testing.T) {
	var expected float32
	var actual float32

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Float64(t *testing.T) {
	expected := 5.431
	var actual float64

	refl := reflect.ValueOf(&actual)
	err := Safe("5.431", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Float64Error(t *testing.T) {
	var expected float64
	var actual float64

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Int(t *testing.T) {
	expected := 5
	var actual int

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_IntError(t *testing.T) {
	var expected int
	var actual int

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Int8(t *testing.T) {
	var expected int8 = 5
	var actual int8

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Int8Error(t *testing.T) {
	var expected int8
	var actual int8

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Int16(t *testing.T) {
	var expected int16 = 5
	var actual int16

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Int16Error(t *testing.T) {
	var expected int16
	var actual int16

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Int32(t *testing.T) {
	var expected int32 = 5
	var actual int32

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Int32Error(t *testing.T) {
	var expected int32
	var actual int32

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Int64(t *testing.T) {
	var expected int64 = 5
	var actual int64

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Int64Error(t *testing.T) {
	var expected int64
	var actual int64

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Uint(t *testing.T) {
	var expected uint = 5
	var actual uint

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_UintError(t *testing.T) {
	var expected uint
	var actual uint

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Uint8(t *testing.T) {
	var expected uint8 = 5
	var actual uint8

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Uint8Error(t *testing.T) {
	var expected uint8
	var actual uint8

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Uint16(t *testing.T) {
	var expected uint16 = 5
	var actual uint16

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_uint16Error(t *testing.T) {
	var expected uint16
	var actual uint16

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Uint32(t *testing.T) {
	var expected uint32 = 5
	var actual uint32

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Uint32Error(t *testing.T) {
	var expected uint32
	var actual uint32

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_Uint64(t *testing.T) {
	var expected uint64 = 5
	var actual uint64

	refl := reflect.ValueOf(&actual)
	err := Safe("5", refl.Elem())

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}
}

func TestSafe_Uint64Error(t *testing.T) {
	var expected uint64
	var actual uint64

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if expected != actual {
		t.Errorf("Safe returned unexpected value: got %v want %v", actual, expected)
	}

	if err == nil {
		t.Error("Map should have returned convertion error")
	}
}

func TestSafe_UnsupportedError(t *testing.T) {
	var actual time.Time

	refl := reflect.ValueOf(&actual)
	err := Safe("fail", refl.Elem())

	if err == nil {
		t.Error("Map should have returned unsupported error")
	}
}
