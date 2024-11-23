package handlers

import (
	"HTTPMessenger/auth/database"
	"HTTPMessenger/auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil { // Пытаемся десериализовать данные с запроса в структуру
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось десериализовать в структуру"})
	}

	if _, exite := database.Users[user.Email]; !exite {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Пользователя нету в базе данных"})
		return
	} else {
		result := database.Users[user.UserName]

		userInfo := struct {
			Email    string      `json:"email"`
			UserName models.User `json:"name"`
		}{
			Email:    user.Email,
			UserName: result,
		}
		ctx.JSON(http.StatusOK, gin.H{"user": userInfo})
	}

}
