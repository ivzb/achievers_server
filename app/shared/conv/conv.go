package conv

import (
	"fmt"
	"reflect"
	"strconv"
)

// Safe convert a string to concrete value type and assigns the value
func Safe(s string, v reflect.Value) error {
	var err error

	// convert to correct type
	switch v.Kind() {
	case reflect.String:
		v.SetString(s)
	case reflect.Bool:
		var parsed bool
		parsed, err = strconv.ParseBool(s)
		v.SetBool(parsed)
	case reflect.Float32:
		var parsed float64
		parsed, err = strconv.ParseFloat(s, 32)
		v.SetFloat(parsed)
	case reflect.Float64:
		var parsed float64
		parsed, err = strconv.ParseFloat(s, 64)
		v.SetFloat(parsed)
	case reflect.Int:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 10, 0)
		v.SetInt(parsed)
	case reflect.Int8:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 0, 8)
		v.SetInt(parsed)
	case reflect.Int16:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 0, 16)
		v.SetInt(parsed)
	case reflect.Int32:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 0, 32)
		v.SetInt(parsed)
	case reflect.Int64:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 0, 64)
		v.SetInt(parsed)
	case reflect.Uint:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 10, 0)
		v.SetUint(parsed)
	case reflect.Uint8:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 0, 8)
		v.SetUint(parsed)
	case reflect.Uint16:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 0, 16)
		v.SetUint(parsed)
	case reflect.Uint32:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 0, 32)
		v.SetUint(parsed)
	case reflect.Uint64:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 0, 64)
		v.SetUint(parsed)
	default:
		return fmt.Errorf("Type conversion is not supported for type: %v", v.Type())
	}

	return err
}
