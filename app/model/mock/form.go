package mock

import (
	"net/http"

	"github.com/ivzb/achievers_server/app/shared/form"
)

type Map struct {
	Err error
}

type Former struct {
	MapMock Map
}

func (mock *Former) Map(r *http.Request, model interface{}) error {
	r.ParseForm()

	_ = form.Map(r.PostForm, model)

	return mock.MapMock.Err
}
