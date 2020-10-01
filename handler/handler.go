package handler

import (
	"fmt"
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
	refreshBase64 := utils.EncodeToBase64(refresh)
	refreshBcrypted, err := utils.EncodeToBcryptHash(fmt.Sprint(refresh))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate tokens"})
		return
	}

	err = saveToken(guid, refreshBcrypted)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate tokens"})
		return
	}

	tokens := model.TokenPair{
		AccessToken:  access,
		RefreshToken: refreshBase64,
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
	return
}

// RefreshTokens ...
func RefreshTokens(c *gin.Context) {

}

// DeleteRefreshToken ...
func DeleteRefreshToken(c *gin.Context) {
	guid, _ := c.GetQuery("guid")
	if !utils.IsValidToken(guid) {
		c.JSON(http.StatusLengthRequired, gin.H{"message": "Inavalid token's length"})
		return
	}

	isExists := isDocumentExists(guid)
	if !isExists {
		c.JSON(http.StatusOK, gin.H{"message": "No user with such guid"})
		return
	}

	err := deleteToken(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't delete token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
	return
}

/* --------------- db handlers ----------- */
func isDocumentExists(guid string) bool {
	db := model.NewDB(settings.OpenDBConnection())
	return db.IsExists(guid)
}

func deleteToken(guid string) error {
	db := model.NewDB(settings.OpenDBConnection())
	err := db.RemoveTokenFromDocument(guid)
	return err
}

func saveToken(guid string, token []byte) error {
	db := model.NewDB(settings.OpenDBConnection())
	err := db.Save(guid, token)

	return err
}
