package handlers

import "github.com/gin-gonic/gin"

func RegisterUserHandler(ctx *gin.Context) {
	registerUser(ctx)
}

func SignInHandler(ctx *gin.Context) {
	signIn(ctx)
}

func GetUserHandler(ctx *gin.Context) {

}

func RefreshTokenHandler(ctx *gin.Context) {

}
