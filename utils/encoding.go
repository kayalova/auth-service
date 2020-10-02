package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func EncodeToBase64(data string) string {
	// s := fmt.Sprint(data)
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func EncodeToBcryptHash(data string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
}

func CompareHashAndToken(hashed []byte, token string) error {
	return bcrypt.CompareHashAndPassword(hashed, []byte(token))
}
