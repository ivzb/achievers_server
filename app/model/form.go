package model

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/ivzb/achievers_server/app/shared/form"
)

var (
	errNotStruct        = errors.New("model is not a struct")
	errWrongContentType = errors.New("content-type of request is incorrect")
)

type Former interface {
	Map(r *http.Request, model interface{}) error
}

type Form struct {
}

func NewFormer() *Form {
	return &Form{}
}

func (f *Form) Map(r *http.Request, model interface{}) error {
	if !isStructPtr(model) {
		return errNotStruct
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		return errWrongContentType
	}

	if err := r.ParseForm(); err != nil {
		return err
	}

	return form.Map(r.PostForm, model)
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
