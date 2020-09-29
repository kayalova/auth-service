package utils

import (
	"hash/fnv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kayalova/auth-service/settings"
)

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
