package handlers

import (
	"HTTPMessenger/auth/database"
	"HTTPMessenger/auth/models"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func registerUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil { // Пытаемся десериализовать данные с запроса в структуру
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось десериализовать в структуру"})
	}
	if _, exist := database.Users[user.Email]; exist { // Проверяем есть ли пользователь с таким email
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким email уже есть"})
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Hash), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось хэшировать пароль"})
	}
	user.Hash = string(hashPassword)

	database.Users[user.Email] = user
	fmt.Println(database.Users)
	ctx.JSON(http.StatusOK, gin.H{"messege": "Пользователь успешно зарегистрирован"})

}
