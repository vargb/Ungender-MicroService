package validator

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
)

func (v *Validator) Required(field string, value interface{}) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if IsEmpty(value) {
		v.Errors[field] = fmt.Sprintf("%s is required", field)
		logrus.Error("%s is required", field)
		return false
	}

	return true
}

func IsEmpty(value interface{}) bool {
	t := reflect.ValueOf(value)

	switch t.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return t.Len() == 0
	}

	return false
}
