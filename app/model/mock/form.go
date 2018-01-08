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

type Form struct {
	HttpRequest     *http.Request
	MapMock         Map
	StringValueMock StringValue
	IntValueMock    IntValue
}

func (mock Form) Map(model interface{}) error {
	mock.HttpRequest.ParseForm()

	_ = form.Map(mock.HttpRequest.PostForm, model)

	return mock.MapMock.Err
}

func (mock Form) StringValue(key string) (string, error) {
	return mock.StringValueMock.Str, mock.StringValueMock.Err
}

func (mock Form) IntValue(key string) (int, error) {
	return mock.IntValueMock.Int, mock.IntValueMock.Err
}
