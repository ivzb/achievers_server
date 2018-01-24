package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/config"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/server"
)

func fileSingleForm() *map[string]string {
	return &map[string]string{
		consts.ID: mockFileID,
	}
}

var fileSingleTests = []*test{
	constructFileSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructFileSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(fileSingleForm(), consts.ID),
		db:                 &mock.DB{},
	}),
	constructFileSingleTest(&testInput{
		purpose:            "file not found",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.ID),
		form:               fileSingleForm(),
		config: &config.Config{
			Server: server.Info{
				FileStorage: "no_such_folder",
			},
		},
	}),
	constructFileSingleTest(&testInput{
		purpose:            "file single OK",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    "/" + mockFileID,
		form:               fileSingleForm(),
		config: &config.Config{
			Server: server.Info{
				FileStorage: "",
			},
		},
	}),
}

func constructFileSingleTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockFileID)

	return constructTest(FileSingle, testInput, responseResults)
}
