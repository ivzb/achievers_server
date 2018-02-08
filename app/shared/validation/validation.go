package validate

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/ivzb/achievers_server/app/shared/ptrs"
)

const tag = "validate"

func Validate(model interface{}) error {
	if model == nil {
		return errors.New("model is nil")
	}

	if !ptrs.IsStructPtr(model) {
		return errors.New("model should be pointer to struct")
	}

	// get the struct type
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	// enumerate struct fields
	for i := 0; i < modelType.NumField(); i++ {
		// get struct's value by tag (`validate`)
		fieldType := modelType.Field(i)
		validations := fieldType.Tag.Get(tag)

		if len(validations) == 0 {
			continue
		}

		rules := strings.Split(validations, "&")

		// load corresponding validation rule
		for _, rule := range rules {
			re, err := regexp.Compile(`^(.+?)(\(.*?\))$`)

			if err != nil {
				return err
			}

			result := re.FindStringSubmatch(rule)

			if len(result) != 3 {
				return errors.New(fmt.Sprintf("Unknown rule format: \"%v\""))
			}

			function := result[1]
			args := strings.Split(result[2], ",")
			fieldValue := modelValue.Field(i).Interface()
			value := reflect.ValueOf(fieldValue)

			err = invoke(function, args, fieldType.Type, value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func invoke(
	function string,
	args []string,
	rtype reflect.Type,
	value interface{}) error {

	log.Println(fmt.Sprintf("%v => %v", function, args))

	// todo: create map with all rule functions

	return nil
}

// todo: sample rule function

func required(field string, rtype reflect.Type, value interface{}) error {
	switch rtype.Kind() {
	case reflect.Int:
		if value == 0 {
			return errors.New(fmt.Sprintf("%v is required"))
		}

		return nil
	default:
		return errors.New(fmt.Sprintf("Required rule unsupported type: %v"))
	}

	return nil
}
