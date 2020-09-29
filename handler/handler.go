package handler

import (
	"net/http"

	"github.com/kayalova/auth-service/model"
	"github.com/kayalova/auth-service/settings"
	"github.com/kayalova/auth-service/utils"

	"github.com/gin-gonic/gin"
)

// GenerateTokens using guid
func GenerateTokens(c *gin.Context) {
	guid, _ := c.GetQuery("guid")
	if !utils.IsValidToken(guid) {
		c.JSON(http.StatusLengthRequired, gin.H{"message": "Inavalid token's length"})
		return
	}

	access, err := utils.GenerateAccessToken(guid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate tokens"})
		return
	}

	refresh := utils.GenerateRefreshToken(guid)
	refreshbase64 := utils.EncodeToBase64(refresh)
	refreshBcrypted, err := utils.EncodeToBcryptHash(string(refresh))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate tokens"})
		return
	}

	tokens := model.TokenPair{
		access,
		refreshbase64,
	}

	err = saveToken(guid, refreshBcrypted)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate tokens"})
		return
	}
	c.JSON(http.StatusOK, tokens)
	return
}

// RefreshTokens ...
func RefreshTokens(c *gin.Context) {

}

// DeleteRefreshToken ...
func DeleteRefreshToken(c *gin.Context) {

}

func saveToken(guid string, token []byte) error {
	db := model.NewDB(settings.OpenDBConnection())
	err := db.Save(guid, token)

	return err
}
