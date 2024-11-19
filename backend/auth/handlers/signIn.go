package handlers

import (
	"HTTPMessenger/auth/database"
	"HTTPMessenger/auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func signIn(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные данные"})
	}

	searchUserInDatabase, exists := database.Users[user.Email]
	if err := bcrypt.CompareHashAndPassword([]byte(searchUserInDatabase.Hash), []byte(user.Hash)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
		return
	}
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}
	accessToken, err := generateAccessTokens(user.UserName)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}
	refreshToken, err := generateRefreshTokens(user.UserName)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})

}
