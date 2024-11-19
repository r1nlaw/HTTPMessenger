package handlers

import (
	"HTTPMessenger/auth/database"
	"HTTPMessenger/auth/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateAccessTokens(username string) (string, error) {
	claims := models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString(database.SecretKey)
}
func generateRefreshTokens(username string) (string, error) {
	claims := models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString(database.SecretKey)
}
