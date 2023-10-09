package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func Validate(data interface{}) []string {
	errMsg := make([]string, 0)

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			errMsg = append(errMsg, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.Field(),
				err.Value,
				err.Tag,
			))
		}
	}

	return errMsg
}
