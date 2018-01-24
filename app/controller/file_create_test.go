package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ivzb/achievers_server/app/shared/consts"
)

//func fileCreateForm() *map[string]string {
//return &map[string]string{
//consts.File: "file",
//}
//}

func fileCreateForm() string {
	//return strings.NewReader(strings.Replace(message, "\n", "\r\n", -1))
	return `
--xxx

Content-Disposition: form-data; name="file"; filename="file"

Content-Type: application/octet-stream

Content-Transfer-Encoding: binary


binary data

--xxx--`
}

var fileCreateTests = []*test{
	constructFileCreateTest(&testInput{
		purpose:            "invalid request method",
		requestMethod:      consts.GET,
		responseType:       Core,
		responseStatusCode: http.StatusMethodNotAllowed,
		responseMessage:    consts.MethodNotAllowed,
		//form:               &map[string]string{},
	}),
	constructFileCreateTest(&testInput{
		purpose:            "former error",
		requestMethod:      consts.POST,
		responseType:       Core,
		responseStatusCode: http.StatusBadRequest,
		responseMessage:    "request Content-Type isn't multipart/form-data",
		//form:               &map[string]string{},
		removeHeaders: true,
	}),
	//constructFileCreateTest(&testInput{
	//purpose:            "file create db error",
	//requestMethod:      consts.POST,
	//responseType:       Core,
	//responseStatusCode: http.StatusInternalServerError,
	//responseMessage:    consts.FriendlyErrorMessage,
	////form:               fileCreateForm(),
	////multipartForm: fileCreateForm(),
	////isMultipart:        true,
	////db: &mock.DB{
	////InvolvementMock: mock.Involvement{
	////ExistsMock: mock.InvolvementExists{Bool: true},
	////},
	////FileMock: mock.File{
	////CreateMock: mock.FileCreate{Err: mockDbErr},
	////},
	////},
	//}),
	//constructFileCreateTest(&testInput{
	//purpose:            "file create ok",
	//requestMethod:      consts.POST,
	//responseType:       Retrieve,
	//responseStatusCode: http.StatusCreated,
	//responseMessage:    fmt.Sprintf(consts.FormatCreated, consts.File),
	//form:               fileCreateForm(),
	//db: &mock.DB{
	//InvolvementMock: mock.Involvement{
	//ExistsMock: mock.InvolvementExists{Bool: true},
	//},
	//FileMock: mock.File{
	//CreateMock: mock.FileCreate{ID: mockID},
	//},
	//},
	//}),
}

func constructFileCreateTest(testInput *testInput) *test {
	responseResults, _ := json.Marshal(mockID)
	return constructTest(FileCreate, testInput, responseResults)
}
