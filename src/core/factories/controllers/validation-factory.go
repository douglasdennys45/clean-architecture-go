package controllers

import (
	"github.com/douglasdennys/go-mongodb/src/presentation/protocols"
	"github.com/douglasdennys/go-mongodb/src/validation"
)

func MakeValidation() protocols.Validation {
	return validation.NewStructValidatorAdapter()
}