package validation

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

type IsValid struct {
	Email string `valid:"required,email" faker:"email"`
}

func TestNewStructValidatorAdapter(t *testing.T) {
	valid := NewStructValidatorAdapter()
	assert.NotNil(t, valid)
}

func TestValidators_Validate(t *testing.T) {
	valid := NewStructValidatorAdapter()
	var mock IsValid
	err := faker.FakeData(&mock)
	if err != nil {
		panic(err)
	}

	err = valid.Validate(mock)
	assert.Nil(t, err)
}

func TestValidators_InValidate(t *testing.T) {
	valid := NewStructValidatorAdapter()
	var mock *IsValid

	err := valid.Validate(mock)
	assert.NotNil(t, err)
}
