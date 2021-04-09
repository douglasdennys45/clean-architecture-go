package criptography

type Hasher interface {
	Hash(plaintext string) (string, error)
}