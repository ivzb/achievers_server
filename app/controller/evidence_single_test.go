package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

type evidenceSingleTest struct {
	purpose            string
	requestMethod      string
	responseType       int
	responseStatusCode int
	responseMessage    string
	formID             string
	dbEvidenceExists   mock.EvidenceExists
	dbEvidenceSingle   mock.EvidenceSingle
}

var evidenceSingleTests = []*test{
	constructEvidenceSingleTest(&evidenceSingleTest{
		purpose:            "invalid request method",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		formID:             "",
		dbEvidenceExists:   mock.EvidenceExists{},
		dbEvidenceSingle:   mock.EvidenceSingle{},
	}),
	constructEvidenceSingleTest(&evidenceSingleTest{
		purpose:            "missing id",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, id),
		formID:             "",
		dbEvidenceExists:   mock.EvidenceExists{},
		dbEvidenceSingle:   mock.EvidenceSingle{},
	}),
	constructEvidenceSingleTest(&evidenceSingleTest{
		purpose:            "evidence exists db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formID:             mockID,
		dbEvidenceExists:   mock.EvidenceExists{Err: mockDbErr},
		dbEvidenceSingle:   mock.EvidenceSingle{},
	}),
	constructEvidenceSingleTest(&evidenceSingleTest{
		purpose:            "evidence does not exist",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, evidence),
		formID:             mockID,
		dbEvidenceExists:   mock.EvidenceExists{Bool: false},
		dbEvidenceSingle:   mock.EvidenceSingle{},
	}),
	constructEvidenceSingleTest(&evidenceSingleTest{
		purpose:            "evidence single db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		formID:             mockID,
		dbEvidenceExists:   mock.EvidenceExists{Bool: true},
		dbEvidenceSingle:   mock.EvidenceSingle{Err: mockDbErr},
	}),
	constructEvidenceSingleTest(&evidenceSingleTest{
		purpose:            "evidence single OK",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, evidence),
		formID:             mockID,
		dbEvidenceExists:   mock.EvidenceExists{Bool: true},
		dbEvidenceSingle:   mock.EvidenceSingle{Evd: mock.Evidence()},
	}),
}

func constructEvidenceSingleTest(testInput *evidenceSingleTest) *test {
	responseResults, _ := json.Marshal(mock.Evidence())

	db := &mock.DB{
		EvidenceExistsMock: testInput.dbEvidenceExists,
		EvidenceSingleMock: testInput.dbEvidenceSingle,
	}

	logger := &mock.Logger{}

	return &test{
		purpose: testInput.purpose,
		handle:  EvidenceSingle,
		request: constructTestRequest(
			testInput.requestMethod,
			constructForm(map[string]string{
				id: testInput.formID,
			}),
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
