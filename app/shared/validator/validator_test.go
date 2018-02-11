package validator

import (
	"testing"
)

func TestValidate_Fail(t *testing.T) {
	type mock struct {
		ID string `validate:"string.(2,5)"`
	}

	model := mock{
		ID: "too long id",
	}

	actual := Validate(model)

	if actual == nil {
		t.Errorf("Validator return unexpected result:\nwant: string too long error\ngot: %v", actual)
	}
}

func TestValidate(t *testing.T) {
	type mock struct {
		ID string `validate:"-"`
	}

	model := mock{
		ID: "",
	}

	actual := Validate(model)

	if actual != nil {
		t.Errorf("Validator return unexpected result:\nwant: empty errors slice\ngot: %v", actual)
	}
}
