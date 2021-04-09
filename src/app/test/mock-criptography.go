package test

import (
	"errors"
	"github.com/douglasdennys/go-mongodb/src/app/protocols/criptography"
)

type criptographySpy struct {}

func NewCriptographySpy() criptography.Hasher {
	return &criptographySpy{}
}

func (c criptographySpy) Hash(plaintext string) (string, error) {
	return "", errors.New("error")
}

type criptographyOkSpy struct {}

func NewCriptographySuccessSpy() criptography.Hasher {
	return &criptographyOkSpy{}
}

func (c criptographyOkSpy) Hash(plaintext string) (string, error) {
	return "123", nil
}