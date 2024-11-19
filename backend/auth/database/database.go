package database

import "HTTPMessenger/auth/models"

var SecretKey = []byte("secret")

var Users = make(map[string]models.User)
