package form

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

type mockValid struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type mockUnsupported struct {
	CreatedAt time.Time `json:"created_at"`
}

func TestModelValue_ValidMap(t *testing.T) {
	expectedID := 5
	expectedTitle := "Mock"

	form := url.Values{}
	form.Add("id", strconv.Itoa(expectedID))
	form.Add("title", expectedTitle)

	req := httptest.NewRequest("GET", "/auth", strings.NewReader(form.Encode()))

	req.PostForm = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	actual := &mockValid{}

	err := ModelValue(req, actual)

	if err != nil {
		t.Errorf("Map returned unexpected error: %v",
			err)
	}

	formatUnexpected := "Map struct returned unexpected %s: got %v want %v"

	if expectedID != actual.ID {
		t.Errorf(formatUnexpected, "id", expectedID, actual.ID)
	}

	if expectedTitle != actual.Title {
		t.Errorf(formatUnexpected, "title", expectedTitle, actual.Title)
	}
}

func TestModelValue_MissingFormValues(t *testing.T) {
	expectedID := 0
	expectedTitle := ""

	form := url.Values{}
	form.Add("id", "")
	form.Add("title", "")

	req := httptest.NewRequest("GET", "/test", strings.NewReader(form.Encode()))

	req.PostForm = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	actual := &mockValid{}

	err := ModelValue(req, actual)

	if err != nil {
		t.Errorf("Map returned unexpected error: %v", err)
	}

	formatUnexpected := "Map struct returned unexpected %s: got %v want %v"

	if expectedID != actual.ID {
		t.Errorf(formatUnexpected, "id", expectedID, actual.ID)
	}

	if expectedTitle != actual.Title {
		t.Errorf(formatUnexpected, "title", expectedTitle, actual.Title)
	}
}

func TestModelValue_UnsupportedConvert(t *testing.T) {
	form := url.Values{}
	form.Add("created_at", "mock_time")

	req := httptest.NewRequest("GET", "/test", strings.NewReader(form.Encode()))

	req.PostForm = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	actual := &mockUnsupported{}

	err := ModelValue(req, actual)

	if err == nil {
		t.Error("Map should have returned unsupported error")
	}
}

func TestModelValue_NotPointer(t *testing.T) {
	form := url.Values{}
	form.Add("created_at", "mock_time")

	req := httptest.NewRequest("GET", "/test", strings.NewReader(form.Encode()))

	req.PostForm = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	actual := mockUnsupported{}

	err := ModelValue(req, actual)

	if err == nil {
		t.Error("Map should have returned unsupported error")
	}
}

func TestModelValue_NotStruct(t *testing.T) {
	form := url.Values{}
	form.Add("created_at", "mock_time")

	req := httptest.NewRequest("GET", "/test", strings.NewReader(form.Encode()))

	req.PostForm = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	arr := make([]string, 2)
	actual := &arr

	err := ModelValue(req, actual)

	if err == nil {
		t.Error("Map should have returned unsupported error")
	}
}

func TestModelValue_InvalidContentType(t *testing.T) {
	expectedID := 5
	expectedTitle := "Mock"

	form := url.Values{}
	form.Add("id", strconv.Itoa(expectedID))
	form.Add("title", expectedTitle)

	//req := httptest.NewRequest("POST", "/auth", nil)
	header := http.Header{}
	header.Add("Content-Type", "text/plain; boundary=")

	req := &http.Request{
		Method: "POST",
		Header: header,
		Body:   ioutil.NopCloser(strings.NewReader("body")),
	}

	//req.PostForm = form
	//req.Header = http.Header{}
	//req.Header.Add("Content-Type", "text/plain; boundary=")

	actual := &mockValid{}

	err := ModelValue(req, actual)

	if err == nil {
		t.Error("Map should have returned unsupported error")
	}
}

//func TestModelValue_InvalidForm(t *testing.T) {
//req, _ := http.NewRequest(
//"POST",
//"/test",
//nil,
//)

//req.Header = http.Header{}
//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

//actual := &mockValid{}

//err := ModelValue(req, actual)

//if err == nil {
//t.Error("ModelValue should have returned error")
//}
//}

func TestStringValue_ValidString(t *testing.T) {
	expectedTitle := "Mock"

	form := url.Values{}
	form.Add("title", expectedTitle)

	req := httptest.NewRequest("GET", "/test", strings.NewReader(form.Encode()))

	req.Form = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	actualTitle, err := StringValue(req, "title")

	if err != nil {
		t.Errorf("StringValue returned unexpected error: %v",
			err)
	}

	formatUnexpected := "StringValue returned unexpected %s: got %v want %v"

	if expectedTitle != actualTitle {
		t.Errorf(formatUnexpected, "title", expectedTitle, actualTitle)
	}
}

func TestStringValue_MissingString(t *testing.T) {
	expected := ""

	form := url.Values{}

	req := httptest.NewRequest("GET", "/test", strings.NewReader(form.Encode()))

	req.Form = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	actual, err := StringValue(req, "title")

	if err == nil {
		t.Errorf("StringValue should have returned an error, but didn't")
	}

	formatUnexpected := "StringValue returned unexpected %s: got %v want %v"

	if expected != actual {
		t.Errorf(formatUnexpected, "title", expected, actual)
	}
}

func TestIntValue_ValidInt(t *testing.T) {
	expectedID := 5

	form := url.Values{}
	form.Add("id", strconv.Itoa(expectedID))

	req := httptest.NewRequest("GET", "/test", strings.NewReader(form.Encode()))

	req.Form = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	actualID, err := IntValue(req, "id")

	if err != nil {
		t.Errorf("IntValue returned unexpected error: %v",
			err)
	}

	formatUnexpected := "IntValue returned unexpected %s: got %v want %v"

	if expectedID != actualID {
		t.Errorf(formatUnexpected, "id", expectedID, actualID)
	}
}

func TestIntValue_MissingInt(t *testing.T) {
	expectedID := 0

	form := url.Values{}

	req := httptest.NewRequest("GET", "/test", strings.NewReader(form.Encode()))

	req.Form = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	actualID, err := IntValue(req, "id")

	if err == nil {
		t.Errorf("IntValue should have returned an error, but didn't")
	}

	formatUnexpected := "IntValue returned unexpected %s: got %v want %v"

	if expectedID != actualID {
		t.Errorf(formatUnexpected, "id", expectedID, actualID)
	}
}

func TestIntValue_InvalidInt(t *testing.T) {
	expectedID := 0

	form := url.Values{}
	form.Add("id", "five")

	req := httptest.NewRequest("GET", "/test", strings.NewReader(form.Encode()))

	req.Form = form
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	actualID, err := IntValue(req, "id")

	if err == nil {
		t.Errorf("IntValue should have returned an error, but didn't")
	}

	formatUnexpected := "IntValue returned unexpected %s: got %v want %v"

	if expectedID != actualID {
		t.Errorf(formatUnexpected, "id", expectedID, actualID)
	}
}
