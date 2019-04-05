package server

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	mathRand "math/rand"

	"github.com/mjarkk/mini-exec/src/utils"
)

// Key is the access key to the server
var Key = ""

// Validation is the validation string
var Validation = ""

// GenValidation generates a new key and prints it to stdout
func GenValidation() error {
	key, err := randomString(15)
	if err != nil {
		return err
	}

	Validation = key

	fmt.Println("Validation:", key)
	return nil
}

// GenerateKey generates a new key
func GenerateKey(validation string) error {
	key, err := randomString(15)
	if err != nil {
		return err
	}

	if Validation != validation {
		return errors.New("Wrong validation key")
	}

	Key = key

	utils.Println("New login key:", key)
	return nil
}

// randomString generates a purly random string with the lenght of n
func randomString(length int) (string, error) {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(2147483647))
	if err != nil {
		return "", err
	}
	r := mathRand.New(mathRand.NewSource(randomNumber.Int64()))
	possibleLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	l := int64(len(possibleLetters))

	toReturn := ""
	for i := 0; i < length; i++ {
		toReturn = toReturn + string(possibleLetters[r.Int63n(l)])
	}

	return toReturn, nil
}
