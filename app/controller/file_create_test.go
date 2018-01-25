package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/shared/config"
	"github.com/ivzb/achievers_server/app/shared/consts"
	"github.com/ivzb/achievers_server/app/shared/server"
	uMock "github.com/ivzb/achievers_server/app/shared/uuid/mock"
)

func fileCreateForm() string {
	return `
--` + mockMultipartBoundry + `
Content-Disposition: form-data; name="file"; filename="file.txt"
Content-Type: text/plain

Test text file
--` + mockMultipartBoundry + `--
`
}

var fileCreateTests = []*test{
	constructFileCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
	}),
	constructFileCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "request Content-Type isn't multipart/form-data",
		removeHeaders:      true,
	}),
	constructFileCreateTest(&testInput{
		purpose:            "file create generate err error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		multipartForm:      fileCreateForm(),
		uuider: &uMock.UUIDer{
			GenerateMock: uMock.Generate{
				Err: mockUUIDerErr,
			},
		},
	}),
	constructFileCreateTest(&testInput{
		purpose:            "file create file path not found",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		multipartForm:      fileCreateForm(),
		uuider: &uMock.UUIDer{
			GenerateMock: uMock.Generate{
				UUID: mockUUID,
			},
		},
		config: &config.Config{
			Server: server.Info{
				FileStorage: "/not_existing_folder",
			},
		},
	}),
	constructFileCreateTest(&testInput{
		purpose:            "file create ok",
		requestMethod:      consts.POST,
		responseType:       Retrieve,
		responseStatusCode: http.StatusCreated,
		responseMessage:    fmt.Sprintf(consts.FormatCreated, consts.File),
		multipartForm:      fileCreateForm(),
		uuider: &uMock.UUIDer{
			GenerateMock: uMock.Generate{
				UUID: mockID,
			},
		},
		config: &config.Config{
			Server: server.Info{
				FileStorage: "/tmp",
			},
		},
	}),
}

func constructFileCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)
	return constructTest(FileCreate, testInput, responseResults)
}
