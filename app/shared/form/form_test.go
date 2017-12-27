package form

import (
	"net/url"
	"strconv"
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

func TestMap_ValidMap(t *testing.T) {
	expectedID := 5
	expectedTitle := "Mock"

	form := url.Values{}
	form.Add("id", strconv.Itoa(expectedID))
	form.Add("title", expectedTitle)

	actual := &mockValid{}

	err := Map(form, actual)

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

func TestMap_MissingFormValues(t *testing.T) {
	expectedID := 0
	expectedTitle := ""

	form := url.Values{}
	form.Add("id", "")
	form.Add("title", "")

	actual := &mockValid{}

	err := Map(form, actual)

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

func TestMap_UnsupportedConvert(t *testing.T) {
	form := url.Values{}
	form.Add("created_at", "mock_time")

	actual := &mockUnsupported{}

	err := Map(form, actual)

	if err == nil {
		t.Error("Map should have returned unsupported error")
	}
}
