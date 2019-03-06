package validator_util

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.CustomTypeTagMap.Set("password", func(i interface{}, o interface{}) bool {
		return false
	})
	govalidator.CustomTypeTagMap.Set("userName", func(i interface{}, o interface{}) bool {
		return false
	})
}
