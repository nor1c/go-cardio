package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct[T any](payload T) []*IError {
	var errors []*IError

	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elem IError

			elem.Error = true
			elem.FailedField = err.StructNamespace()
			elem.Tag = err.Tag()
			elem.Value = err.Value()

			errors = append(errors, &elem)
		}
	}

	return errors
}
