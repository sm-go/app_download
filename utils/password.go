package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPW, nil
}

func ComparePasswords(oldPassword string, newPassword string) (bool, error) {
	pwsalt := strings.Split(oldPassword, ".")
	if len(pwsalt) != 2 {
		return false, errors.New("wrong password")
	}

	salt, err := hex.DecodeString(pwsalt[1])
	if err != nil {
		return false, errors.New("unable to verify user password")
	}

	shash, err := scrypt.Key([]byte(newPassword), salt, 32768, 8, 1, 32)
	if err != nil {
		return false, err
	}

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
