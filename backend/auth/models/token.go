package models

import "github.com/golang-jwt/jwt"

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// Структура Claims в токенах
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
