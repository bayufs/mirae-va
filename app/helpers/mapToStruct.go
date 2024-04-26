package helpers

import (
	"fmt"
	"reflect"
)

func MapToStruct(data map[string]interface{}, model interface{}) error {
	modelType := reflect.TypeOf(model)

	if modelType.Kind() != reflect.Ptr || modelType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("model must be a pointer to a struct")
	}

	modelValue := reflect.ValueOf(model).Elem()

	for key, value := range data {
		field := modelValue.FieldByName(key)

		if !field.IsValid() {
			continue // Skip if the field does not exist
		}
		if !field.CanSet() {
			continue // Skip if the field cannot be set (unexported field)
		}

		fieldValue := reflect.ValueOf(value)
		if fieldValue.Type().ConvertibleTo(field.Type()) {
			field.Set(fieldValue.Convert(field.Type()))
		}
	}

	return nil
}
