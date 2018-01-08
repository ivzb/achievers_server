package form

import (
	"net/url"
	"reflect"

	"github.com/ivzb/achievers_server/app/shared/conv"
)

const (
	json = "json"
)

// Map recives http.Request and maps form values to target model
func Map(form url.Values, model interface{}) error {
	// get the struct type
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	// enumerate model fields
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		// get form value by model's tag (`json`)
		key := field.Tag.Get(json)
		value := form.Get(key)

		fieldValue := modelValue.FieldByName(field.Name)

		if len(value) > 0 {
			err := conv.Safe(value, fieldValue)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
