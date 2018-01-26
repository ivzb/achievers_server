package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func evidenceSingleForm() *map[string]string {
	return &map[string]string{
		consts.ID: mockID,
	}
}

var evidenceSingleTests = []*test{
	constructEvidenceSingleTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.ID),
		form:               mapWithout(evidenceSingleForm(), consts.ID),
		db:                 &mock.DB{},
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "evidence exists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               evidenceSingleForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Err: mockDbErr},
			},
		},
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "evidence does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.ID),
		form:               evidenceSingleForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Bool: false},
			},
		},
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "evidence single db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               evidenceSingleForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Bool: true},
				SingleMock: mock.EvidenceSingle{Err: mockDbErr},
			},
		},
	}),
	constructEvidenceSingleTest(&testInput{
		purpose:            "evidence single OK",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Evidence),
		form:               evidenceSingleForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Bool: true},
				SingleMock: mock.EvidenceSingle{Evd: generate.Evidence()},
			},
		},
	}),
}

func constructEvidenceSingleTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(generate.Evidence())

	return constructTest(EvidenceSingle, testInput, responseResults)
}
