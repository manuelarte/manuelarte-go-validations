package validators

import "github.com/go-playground/validator/v10"

type IsValidValidator struct {
	TagName string
}

func NewIsValidValidator() IsValidValidator {
	return IsValidValidator{"is-valid"}
}

func NewIsValidValidatorWithCustomTag(tagName string) IsValidValidator {
	return IsValidValidator{tagName}
}

func (iV IsValidValidator) RegisterIsValidValidator(validate *validator.Validate) error {
	return validate.RegisterValidation(iV.TagName, iV.ValidateIsValid)
}

func (iV IsValidValidator) ValidateIsValid(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.IsNil() || field.IsZero() {
		return true
	}
	if isValidField, ok := fl.Field().Interface().(IsValid); ok {
		return isValidField.isValid()
	}
	return false
}

type IsValid interface {
	isValid() bool
}
