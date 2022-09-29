package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Name string
type testRequestStruct struct {
	Name Name `validate:"is-valid"`
}

func (n Name) IsValid() bool {
	return len(n) > 3 && len(n) < 40
}

func TestExistsValidator_EntityExists(t *testing.T) {

	tests := map[string]struct {
		Entity testRequestStruct
	}{
		"testing finding an existing entity": {
			Entity: testRequestStruct{Name: "Manuel"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			validate := validator.New()
			isValidValidator := NewIsValidValidator()
			err := isValidValidator.RegisterValidator(validate)
			if err != nil {
				assert.Failf(t, "Can't migrate entity", "Can't auto migrate entity: %s", err.Error())
			}
			err = validate.Struct(tc.Entity)
			assert.Nil(t, err)
		})
	}
}
