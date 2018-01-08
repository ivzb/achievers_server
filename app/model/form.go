package model

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/ivzb/achievers_server/app/shared/form"
)

var (
	errNotStruct        = errors.New("model is not a struct")
	errWrongContentType = errors.New("content-type of request is incorrect")

	formatMissing = "missing %s"
	formatInvalid = "invalid %s"
)

type Former interface {
	Map(model interface{}) error
	StringValue(key string) (string, error)
	IntValue(key string) (int, error)
}

type Form struct {
	httpRequest *http.Request
}

func NewForm(r *http.Request) Former {
	return &Form{
		httpRequest: r,
	}
}

// Map form values to model and returns error if model is not struct, wrong content type or parse error
func (f *Form) Map(model interface{}) error {
	if !isStructPtr(model) {
		return errNotStruct
	}

	if f.httpRequest.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		return errWrongContentType
	}

	if err := f.httpRequest.ParseForm(); err != nil {
		return err
	}

	return form.Map(f.httpRequest.PostForm, model)
}

// StringValue returns string value by key and error if not found
func (f *Form) StringValue(key string) (string, error) {
	value := f.httpRequest.FormValue(key)

	if value == "" {
		return "", errors.New(fmt.Sprintf(formatMissing, key))
	}

	return value, nil
}

// IntValue returns int value by key and error if not found
func (f *Form) IntValue(key string) (int, error) {
	value, err := f.StringValue(key)

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
