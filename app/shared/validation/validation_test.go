package validate

import (
	"testing"
)

type mock struct {
	ID string `validate:"required()"`

	Title       string `validate:"required()&length(3,5)"`
	Description string `validate:"length(3,5)"`

	InvolvementID int `validate:"min(1)"`
	AchievementID int `validate:"min(1)"`
	UserID        int `validate:"required()&min(1)&max(5)"`
}

func TestValidate_NilModel(t *testing.T) {
	err := Validate(nil)

	if err == nil {
		t.Errorf("Validate should return error, but didn't")
	}
}

func TestValidate_ModelNotStruct(t *testing.T) {
	err := Validate(5)

	if err == nil {
		t.Errorf("Validate should return error, but didn't")
	}
}

func TestValidate_(t *testing.T) {
	model := &mock{
		ID:            "mock_id",
		Title:         "mock_title",
		Description:   "mock_description",
		InvolvementID: 1,
		AchievementID: 2,
		UserID:        3,
	}

	err := Validate(model)

	if err != nil {
		t.Errorf("Validate should not return error, but got \"%v\"", err)
	}
}
