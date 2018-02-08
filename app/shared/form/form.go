package form

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/conv"
	"github.com/ivzb/achievers_server/app/shared/ptrs"
)

const (
	json = "json"
)

var (
	errNotStruct        = errors.New("model is not a struct")
	errWrongContentType = errors.New("content-type of request is incorrect")
)

// Map form values to model and returns error if model is not struct, wrong content type or parse error
func ModelValue(r *http.Request, model interface{}) error {
	if !ptrs.IsStructPtr(model) {
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
		return "", errors.New(fmt.Sprintf(consts.FormatMissing, key))
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
		return 0, errors.New(fmt.Sprintf(consts.FormatInvalid, key))
	}

	return castedValue, nil
}

// MultipartFile returns multipart file by key and error if not found
func MultipartFile(r *http.Request, key string) (multipart.File, *multipart.FileHeader, error) {
	r.ParseMultipartForm(0)
	file, header, err := r.FormFile(key)

	return file, header, err
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
