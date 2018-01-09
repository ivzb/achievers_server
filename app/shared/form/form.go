package form

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/ivzb/achievers_server/app/shared/conv"
)

const (
	json = "json"
)

var (
	errNotStruct        = errors.New("model is not a struct")
	errWrongContentType = errors.New("content-type of request is incorrect")

	formatMissing = "missing %s"
	formatInvalid = "invalid %s"
)

// Map form values to model and returns error if model is not struct, wrong content type or parse error
func ModelValue(r *http.Request, model interface{}) error {
	if !isStructPtr(model) {
		return errNotStruct
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		return errWrongContentType
	}

	if err := r.ParseForm(); err != nil {
		return err
	}

	return mapModel(r.PostForm, model)
}

// StringValue returns string value by key and error if not found
func StringValue(r *http.Request, key string) (string, error) {
	value := r.FormValue(key)

	if value == "" {
		return "", errors.New(fmt.Sprintf(formatMissing, key))
	}

	return value, nil
}

// IntValue returns int value by key and error if not found
func IntValue(r *http.Request, key string) (int, error) {
	value, err := StringValue(r, key)

	if err != nil {
		return 0, err
	}

	castedValue, err := strconv.Atoi(value)

	if err != nil {
		return 0, errors.New(fmt.Sprintf(formatInvalid, key))
	}

	return castedValue, nil
}

// prevent running on types other than struct
func isStructPtr(i interface{}) bool {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		return false
	} else if reflect.TypeOf(i).Elem().Kind() != reflect.Struct {
		return false
	}

	return true
}

// map recives http.Request and maps form values to target model
func mapModel(form url.Values, model interface{}) error {
	// get the struct type
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	// enumerate model fields
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		// get form value by model's tag (`json`)
		key := field.Tag.Get(json)
		value := form.Get(key)

		fieldValue := modelValue.FieldByName(field.Name)

		if len(value) > 0 {
			err := conv.Safe(value, fieldValue)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
