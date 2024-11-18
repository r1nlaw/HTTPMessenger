package server

import (
	"HTTPMessenger/auth/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	routes := gin.Default()

	routes.POST("/registeruser", handlers.RegisterUserHandler) // Регистрация пользователя
	routes.GET("/signin", handlers.SignInHandler)              // Авторизация пользователя
	routes.GET("/user", handlers.GetUserHandler)               // Получение данных о пользователе
	routes.PUT("/refresh", handlers.RefreshTokenHandler)       // Обновление токена

	if err := routes.Run(":8080"); err != nil {
		log.Fatal("Не удалось запустить сервер", err)
	}
}
