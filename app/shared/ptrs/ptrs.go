package ptrs

import "reflect"

// prevent running on types other than struct
func IsStructPtr(i interface{}) bool {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		return false
	} else if reflect.TypeOf(i).Elem().Kind() != reflect.Struct {
		return false
	}

	return true
}
