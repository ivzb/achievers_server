package mock

import (
	"net/http"

	"github.com/ivzb/achievers_server/app/shared/form"
)

type Map struct {
	Err error
}

type StringValue struct {
	Str string
	Err error
}

type IntValue struct {
	Int int
	Err error
}

type Former struct {
	MapMock         Map
	StringValueMock StringValue
	IntValueMock    IntValue
}

func (mock *Former) Map(r *http.Request, model interface{}) error {
	r.ParseForm()

	_ = form.Map(r.PostForm, model)

	return mock.MapMock.Err
}

func (mock *Former) StringValue(r *http.Request, key string) (string, error) {
	return mock.StringValueMock.Str, mock.StringValueMock.Err
}

func (mock *Former) IntValue(r *http.Request, key string) (int, error) {
	return mock.IntValueMock.Int, mock.IntValueMock.Err
}
