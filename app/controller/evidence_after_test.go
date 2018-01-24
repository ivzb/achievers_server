package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivzb/achievers_server/app/db/mock"
	"github.com/ivzb/achievers_server/app/db/mock/generate"
	"github.com/ivzb/achievers_server/app/shared/consts"
)

func evidencesAfterForm() *map[string]string {
	return &map[string]string{
		consts.AfterID: mockID,
	}
}

var evidencesAfterArgs = []string{"9"}

var evidencesAfterTests = []*test{
	constructEvidencesAfterTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               evidencesAfterArgs,
	}),
	constructEvidencesAfterTest(&testInput{
		purpose:            "missing id",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    fmt.Sprintf(consts.FormatMissing, consts.AfterID),
		form:               mapWithout(evidencesAfterForm(), consts.AfterID),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Err: mockDbErr},
			},
		},
		args: evidencesAfterArgs,
	}),
	constructEvidencesAfterTest(&testInput{
		purpose:            "id EvidenceExists db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               evidencesAfterForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Err: mockDbErr},
			},
		},
		args: evidencesAfterArgs,
	}),
	constructEvidencesAfterTest(&testInput{
		purpose:            "id EvidenceExists does not exist",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusNotFound,
		responseMessage:    fmt.Sprintf(consts.FormatNotFound, consts.AfterID),
		form:               evidencesAfterForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Bool: false},
			},
		},
		args: evidencesAfterArgs,
	}),
	constructEvidencesAfterTest(&testInput{
		purpose:            "EvidencesAfterMock db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		form:               evidencesAfterForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Bool: true},
				AfterMock:  mock.EvidencesAfter{Err: mockDbErr},
			},
		},
		args: evidencesAfterArgs,
	}),
	constructEvidencesAfterTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Evidences),
		form:               evidencesAfterForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Bool: true},
				AfterMock:  mock.EvidencesAfter{Evds: generate.Evidences(0)},
			},
		},
		args: []string{"0"},
	}),
	constructEvidencesAfterTest(&testInput{
		purpose:            "4 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Evidences),
		form:               evidencesAfterForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Bool: true},
				AfterMock:  mock.EvidencesAfter{Evds: generate.Evidences(4)},
			},
		},
		args: []string{"4"},
	}),
	constructEvidencesAfterTest(&testInput{
		purpose:            "9 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Evidences),
		form:               evidencesAfterForm(),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				ExistsMock: mock.EvidenceExists{Bool: true},
				AfterMock:  mock.EvidencesAfter{Evds: generate.Evidences(9)},
			},
		},
		args: []string{"9"},
	}),
}

func constructEvidencesAfterTest(testInput *testInput) *test {
	evidencesSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(generate.Evidences(evidencesSize))
	}

	return constructTest(EvidencesAfter, testInput, responseResults)
}
