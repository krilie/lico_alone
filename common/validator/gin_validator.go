package validator

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func PageValid(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if num, ok := field.Interface().(int); ok {
		if num <= 0 {
			return false
		}
	}
	return true
}
