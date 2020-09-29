package utils

func IsValidToken(token string) bool {
	return !(len(token) < 1)
}
