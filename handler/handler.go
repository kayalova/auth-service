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

	tokens, err := utils.GenerateTokens(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate tokens"})
		return
	}

	tokenPair := model.TokenPair{
		AccessToken:  tokens["access"].(string),
		RefreshToken: tokens["refreshClient"].(string),
	}

	err = saveToken(guid, tokens["refreshHash"].([]byte)) //check
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokenPair})
	return
}

// UpdateTokens  generates new tokens' pair for client
func UpdateTokens(c *gin.Context) {
	guid, _ := c.GetQuery("guid")
	refreshClient, _ := c.GetQuery("refreshToken")
	if !utils.IsValidToken(guid) {
		c.JSON(http.StatusLengthRequired, gin.H{"message": "Inavalid token's length"})
		return
	}

	// проверяем существует ли документ с таким ид
	isExists := isDocumentExists(guid)
	if !isExists {
		c.JSON(http.StatusOK, gin.H{"message": "No user with such guid"})
		return
	}

	// проверяем совпадают ли рефреш токен клиента с тем, что в базе
	refreshDB := getDocument(guid).RefreshToken
	err := utils.CompareHashAndToken(refreshDB, refreshClient)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "User with such guid has different refresh token"})
		return
	}

	// генерируем токены
	tokens, err := utils.GenerateTokens(guid)
	tokenPair := model.TokenPair{
		AccessToken:  tokens["access"].(string),
		RefreshToken: tokens["refreshClient"].(string),
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate tokens"})
		return
	}

	r := tokens["refreshHash"]
	b := r.([]byte)
	// ui := r.(uint)
	err = saveToken(guid, b) //check
	if err != nil {
		fmt.Println("I am here")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can't generate tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokenPair})
	return

}

// DeleteRefreshToken deletes resfresh token from db
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

func getDocument(guid string) model.Document {
	db := model.NewDB(settings.OpenDBConnection())
	return db.FindOne(guid)
}
