package utils

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncodeToBase64(data uint32) string {
	s := fmt.Sprint(data)
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func EncodeToBcryptHash(data string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
}
