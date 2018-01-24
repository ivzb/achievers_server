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

var evidencesLatestArgs = []string{"9"}

var evidencesLatestTests = []*test{
	constructEvidencesLatestTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		form:               &map[string]string{},
		args:               evidencesLatestArgs,
	}),
	constructEvidencesLatestTest(&testInput{
		purpose:            "id EvidencesLastID db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				LastIDMock: mock.EvidencesLastID{Err: mockDbErr},
			},
		},
		args: evidencesLatestArgs,
	}),
	constructEvidencesLatestTest(&testInput{
		purpose:            "EvidencesAfterMock db error",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusInternalServerError,
		responseMessage:    consts.FriendlyErrorMessage,
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				LastIDMock: mock.EvidencesLastID{ID: mockID},
				AfterMock:  mock.EvidencesAfter{Err: mockDbErr},
			},
		},
		args: evidencesLatestArgs,
	}),
	constructEvidencesLatestTest(&testInput{
		purpose:            "no results",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Evidences),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				LastIDMock: mock.EvidencesLastID{ID: mockID},
				AfterMock:  mock.EvidencesAfter{Evds: generate.Evidences(0)},
			},
		},
		args: []string{"0"},
	}),
	constructEvidencesLatestTest(&testInput{
		purpose:            "4 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Evidences),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				LastIDMock: mock.EvidencesLastID{ID: mockID},
				AfterMock:  mock.EvidencesAfter{Evds: generate.Evidences(4)},
			},
		},
		args: []string{"4"},
	}),
	constructEvidencesLatestTest(&testInput{
		purpose:            "9 results",
		requestMethod:      consts.GET,
		responseType:       Retrieve,
		responseStatusCode: http.StatusOK,
		responseMessage:    fmt.Sprintf(consts.FormatFound, consts.Evidences),
		db: &mock.DB{
			EvidenceMock: mock.Evidence{
				LastIDMock: mock.EvidencesLastID{ID: mockID},
				AfterMock:  mock.EvidencesAfter{Evds: generate.Evidences(9)},
			},
		},
		args: []string{"9"},
	}),
}

func constructEvidencesLatestTest(testInput *testInput) *test {
	evidencesSize, err := strconv.Atoi(testInput.args[0])

	var responseResults []byte

	if err == nil {
		responseResults, _ = json.Marshal(generate.Evidences(evidencesSize))
	}

	return constructTest(EvidencesLast, testInput, responseResults)
}
