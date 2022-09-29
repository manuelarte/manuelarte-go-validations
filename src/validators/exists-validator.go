package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type ExistsValidator struct {
	Db      *gorm.DB
	TagName string
}

func NewExistsValidator(db *gorm.DB) ExistsValidator {
	return ExistsValidator{db, "exists"}
}

func NewExistsValidatorWithCustomTag(db *gorm.DB, tagName string) ExistsValidator {
	return ExistsValidator{db, tagName}
}

func (eV ExistsValidator) RegisterExistsValidator(validate *validator.Validate) error {
	return validate.RegisterValidation(eV.TagName, eV.ValidateExists, true)
}

func (eV ExistsValidator) ValidateExists(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.IsZero() {
		return true
	}
	param := fl.Param()
	split := strings.Split(param, ".")
	if len(split) != 2 {
		return false
	}
	tableName := split[0]
	lookupField := split[1]

	value := field.Interface()
	typeOf := reflect.TypeOf(value)
	var whereQuery string
	var expectedCount int64
	switch typeOf.Kind() {
	case reflect.Slice:
		whereQuery = fmt.Sprintf("%s in ?", lookupField)
		expectedCount = int64(field.Len())
	default:
		whereQuery = fmt.Sprintf("%s = ?", lookupField)
		expectedCount = 1
	}

	var actualCount int64
	tx := eV.Db.Table(tableName).Where(whereQuery, value).Count(&actualCount)
	if tx.Error != nil {
		return false
	}
	return actualCount == expectedCount
}
