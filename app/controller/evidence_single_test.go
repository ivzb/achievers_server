package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/model/mock"
)

func evidenceSingleForm() *map[string]string {
	return &map[string]string{
		id: mockID,
	}
}

var evidenceSingleTests = []*test{
	constructEvidenceSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      post,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    methodNotAllowed,
		form:               &map[string]string{},
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(formatMissing, id),
		form:               mapWithout(evidenceSingleForm(), id),
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "evidence exists db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               evidenceSingleForm(),
		db: &mock.DB{
			EvidenceExistsMock: mock.EvidenceExists{Err: mockDbErr},
		},
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "evidence does not exist",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(formatNotFound, evidence),
		form:               evidenceSingleForm(),
		db: &mock.DB{
			EvidenceExistsMock: mock.EvidenceExists{Bool: false},
		},
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "evidence single db error",
		requestMethod:      get,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    friendlyErrorMessage,
		form:               evidenceSingleForm(),
		db: &mock.DB{
			EvidenceExistsMock: mock.EvidenceExists{Bool: true},
			EvidenceSingleMock: mock.EvidenceSingle{Err: mockDbErr},
		},
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "evidence single OK",
		requestMethod:      get,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(formatFound, evidence),
		form:               evidenceSingleForm(),
		db: &mock.DB{
			EvidenceExistsMock: mock.EvidenceExists{Bool: true},
			EvidenceSingleMock: mock.EvidenceSingle{Evd: mock.Evidence()},
		},
	}),
}

func constructEvidenceSingleTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mock.Evidence())

	return constructTest(EvidenceSingle, testInput, responseResults)
}
