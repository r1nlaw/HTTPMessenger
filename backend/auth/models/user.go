package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"name"`
	Email    string `gorm:"not nul;unique"`
	Hash     string `gorm:"hash" json:"-"`
}
