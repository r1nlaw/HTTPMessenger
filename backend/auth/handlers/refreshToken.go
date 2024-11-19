package handlers

import (
	"HTTPMessenger/auth/database"
	"HTTPMessenger/auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Refresh(ctx *gin.Context) {
	var token models.Token
	if err := ctx.BindJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
	}

	claims := &models.Claims{}
	parsedToken, err := jwt.Parse(token.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		return database.SecretKey, nil
	})
	if err != nil || !parsedToken.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось спарсить токен"})
		return
	}

	accessToken, err := generateAccessTokens(claims.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось создать новый accessToken"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken})

}
