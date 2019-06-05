package validator

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func UserIdValid(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if str, ok := field.Interface().(string); ok {
		if !IsIdStr(str) {
			return false
		}
	}
	return true
}
