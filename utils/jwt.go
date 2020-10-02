package utils

import (
	"crypto/rand"
	"encoding/hex"
	"hash/fnv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kayalova/auth-service/settings"
)

func GenerateTokens(guid string) (map[string]interface{}, error) {
	var tokens map[string]interface{}
	access, err := GenerateAccessToken(guid)
	if err != nil {
		return tokens, err
	}

	// refresh := GenerateRefreshToken(guid)
	refresh, err := GenerateRefreshToken2()
	if err != nil {
		return tokens, err
	}

	refreshClient := EncodeToBase64(refresh)
	refreshHash, err := EncodeToBcryptHash(refreshClient)

	if err != nil {
		return tokens, err
	}

	tokens = map[string]interface{}{
		"access":        access,
		"refreshClient": refreshClient,
		"refreshHash":   refreshHash,
	}

	return tokens, nil

}

func GenerateAccessToken(guid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = guid
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	key := settings.GetEnvKey("SECRET_KEY", "MY_RESERVED_SECRET_KEY")

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(guid string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(guid))
	return h.Sum32() //hash of int type
}

func GenerateRefreshToken2() (string, error) {
	bytes := make([]byte, 10)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
