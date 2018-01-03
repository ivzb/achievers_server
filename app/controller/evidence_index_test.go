package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func evidencesIndexForm() *map[string]string {
	return &map[string]string{
		page: mockPage,
	}
}

var evidencesIndexArgs = []string{"9"}

var evidencesIndexTests = []*test{
	constructEvidencesIndexTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
		args:               evidencesIndexArgs,
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "missing page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, page),
		form:               mapWithout(evidencesIndexForm(), page),
		args:               evidencesIndexArgs,
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "invalid page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatInvalid, page),
		form: &map[string]string{
			page: "-1",
		},
		args: evidencesIndexArgs,
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               evidencesIndexForm(),
		db: &mock.DB{
			EvidencesAllMock: mock.EvidencesAll{Err: mockDbErr},
		},
		args: evidencesIndexArgs,
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "no results on page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, page),
		form:               evidencesIndexForm(),
		db: &mock.DB{
			EvidencesAllMock: mock.EvidencesAll{Evds: mock.Evidences(0)},
		},
		args: []string{"0"},
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "4 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, evidences),
		form:               evidencesIndexForm(),
		db: &mock.DB{
			EvidencesAllMock: mock.EvidencesAll{Evds: mock.Evidences(4)},
		},
		args: []string{"4"},
	}),
	constructEvidencesIndexTest(&testInput{
		purpose:            "9 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, evidences),
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
