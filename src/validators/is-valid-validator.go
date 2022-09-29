package validators

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type IsValidValidator struct {
	TagName string
}

func NewIsValidValidator() IsValidValidator {
	return IsValidValidator{"is-valid"}
}

func NewIsValidValidatorWithCustomTag(tagName string) IsValidValidator {
	return IsValidValidator{tagName}
}

func (iV IsValidValidator) RegisterValidator(validate *validator.Validate) error {
	return validate.RegisterValidation(iV.TagName, iV.ValidateIsValid)
}

func (iV IsValidValidator) ValidateIsValid(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.IsZero() || (field.Kind() == reflect.Pointer && field.IsNil()) {
		return true
	}
	if isValidField, ok := fl.Field().Interface().(IsValid); ok {
		return isValidField.IsValid()
	}
	return false
}

type IsValid interface {
	IsValid() bool
}
