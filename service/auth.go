package service

import (
	"main/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT(c *gin.Context) {
	jwt, _ := helper.GenerateJWT("6433079093b17af7e4bb8ad8")

	_, err := helper.VerifyJWT("jwt")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": true, "data": ""})
	}

	// for k, v := range jwtdata {
	// 	fmt.Println(k, "value is", v)
	// }

	data := map[string]string{"token": jwt.Token, "refreshToken": jwt.RefreshToken}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func RefreshToken(c *gin.Context) {
	refreshToken := c.Param("refresh-token")

	if refreshToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false})
		return
	}

	if _, err := helper.VerifyJWT(refreshToken); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false})
		return
	}

	jwt, err := helper.GenerateJWT("6433079093b17af7e4bb8ad8")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	data := map[string]string{"token": jwt.Token, "refreshToken": jwt.RefreshToken}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}
