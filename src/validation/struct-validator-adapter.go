package validation

import (
	"github.com/douglasdennys/go-mongodb/src/presentation/protocols"
	"github.com/go-playground/validator/v10"
)

type StructValidatorAdapter interface {
	protocols.Validation
}

type validators struct {
}

func NewStructValidatorAdapter() StructValidatorAdapter {
	return &validators{}
}

func (v *validators) Validate(inputs interface{}) error {
	validate := validator.New()
	return validate.Struct(inputs)
}