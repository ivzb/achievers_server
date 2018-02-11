package validator

import (
	v "github.com/ivzb/validator"
)

func Validate(model interface{}) error {
	errs := v.ValidateStruct(model)

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}
