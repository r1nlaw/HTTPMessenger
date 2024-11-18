package models

type Register struct {
	UserName string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"min=8"`
}
