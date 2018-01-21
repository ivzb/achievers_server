package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func evidencesIndexForm() *map[string]string {
	return &map[string]string{
		consts.Page: mockPage,
	}
}

var evidencesIndexArgs = []string{"9"}

var evidencesIndexTests = []*test{
	constructEvidencesIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               evidencesIndexArgs,
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "missing consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.Page),
		form:               mapWithout(evidencesIndexForm(), consts.Page),
		args:               evidencesIndexArgs,
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "invalid consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatInvalid, consts.Page),
		form: &map[string]string{
			consts.Page: "-1",
		},
		args: evidencesIndexArgs,
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               evidencesIndexForm(),
		db: &mock.DB{
			EvidencesAllMock: mock.EvidencesAll{Err: mockDbErr},
		},
		args: evidencesIndexArgs,
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "no results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.Page),
		form:               evidencesIndexForm(),
		db: &mock.DB{
			EvidencesAllMock: mock.EvidencesAll{Evds: mock.Evidences(0)},
		},
		args: []string{"0"},
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "4 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Evidences),
		form:               evidencesIndexForm(),
		db: &mock.DB{
			EvidencesAllMock: mock.EvidencesAll{Evds: mock.Evidences(4)},
		},
		args: []string{"4"},
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "9 results on consts.Page",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Evidences),
		form:               evidencesIndexForm(),
		db: &mock.DB{
			EvidencesAllMock: mock.EvidencesAll{Evds: mock.Evidences(9)},
		},
		args: []string{"9"},
	}),
}

func constructEvidencesIndexTest(testInput *testInput) *test {
	evidencesSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(mock.Evidences(evidencesSize))
	}

	return constructTest(EvidencesIndex, testInput, responseResults)
}
