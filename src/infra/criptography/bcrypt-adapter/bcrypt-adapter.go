package bcrypt_adapter

import (
	"github.com/douglasdennys/go-mongodb/src/app/protocols/criptography"
	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter interface {
	criptography.Hasher
}

type crypt struct{}

func NewBcryptAdapter() BcryptAdapter {
	return &crypt{}
}

func (b *crypt) Hash(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), 14)
	return string(bytes), err
}