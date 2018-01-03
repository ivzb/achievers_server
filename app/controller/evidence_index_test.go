package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type evidencesIndexTest struct {
	purpose            string
	requestMethod      string
	responseType       int
	responseStatusCode int
	responseMessage    string
	mockPageSize       int
	formPage           string
	dbEvidencesAll     mock.EvidencesAll
}

var evidencesIndexTests = []*test{
	constructEvidencesIndexTest(&evidencesIndexTest{
		purpose:            "invalid request method",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
	}),
	constructEvidencesIndexTest(&evidencesIndexTest{
		purpose:            "9 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, evidences),
		mockPageSize:       9,
		formPage:           "0",
		dbEvidencesAll:     mock.EvidencesAll{Evds: mock.Evidences(9)},
	}),
	constructEvidencesIndexTest(&evidencesIndexTest{
		purpose:            "4 results on page",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, evidences),
		mockPageSize:       4,
		formPage:           "1",
		dbEvidencesAll:     mock.EvidencesAll{Evds: mock.Evidences(4)},
	}),
	constructEvidencesIndexTest(&evidencesIndexTest{
		purpose:            "no results on page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, page),
		mockPageSize:       0,
		formPage:           "2",
		dbEvidencesAll:     mock.EvidencesAll{Evds: mock.Evidences(0)},
	}),
	constructEvidencesIndexTest(&evidencesIndexTest{
		purpose:            "missing page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, page),
		formPage:           "",
	}),
	constructEvidencesIndexTest(&evidencesIndexTest{
		purpose:            "invalid page",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatInvalid, page),
		formPage:           "-1",
	}),
	constructEvidencesIndexTest(&evidencesIndexTest{
		purpose:            "db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formPage:           mockPage,
		dbEvidencesAll:     mock.EvidencesAll{Err: mockDbErr},
	}),
}

func constructEvidencesIndexTest(testInput *evidencesIndexTest) *test {
	responseResults, _ := json.Marshal(mock.Evidences(testInput.mockPageSize))

	db := &mock.DB{EvidencesAllMock: testInput.dbEvidencesAll}

	logger := &mock.Logger{}

	return &test{
		purpose: testInput.purpose,
		handle:  EvidencesIndex,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(map[string]string{page: testInput.formPage}),
			constructEnv(db, logger, nil, nil),
		),
		response: constructTestResponse(
			testInput.responseType,
			testInput.responseStatusCode,
			testInput.responseMessage,
			responseResults,
		),
	}
}
