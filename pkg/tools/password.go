package tools

import (
	"crypto/rand"

	"github.com/sethvargo/go-password/password"
)

type PasswordOptions struct {
	Length     int
	NumDigits  int
	NumSymbols int
}

func GeneratePassword(options PasswordOptions) (string, error) {
	randomBytes := make([]byte, options.Length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	pass, err := password.Generate(options.Length, options.NumDigits, options.NumSymbols, false, false)
	if err != nil {
		return "", err
	}

	return pass, nil
}
