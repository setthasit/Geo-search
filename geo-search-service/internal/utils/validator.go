package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type APIValidator struct {
	validator *validator.Validate
}

func NewAPIValidator(validator *validator.Validate) *APIValidator {
	return &APIValidator{
		validator: validator,
	}
}

func (cv *APIValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func ApiValidateError(err error) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	validationErrors := err.(validator.ValidationErrors)

	message := make([]string, len(validationErrors))
	for idx, itErr := range validationErrors {
		message[idx] = fmt.Sprintf("%s is %s", itErr.StructField(), itErr.Tag())
	}

	return errors.New(strings.Join(message, ", "))
}
