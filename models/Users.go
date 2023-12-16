package models

import (
	"fastprint/helpers"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type UsersRegister struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UsersLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password = helpers.EncryptPassword(user.Password)

	err = nil
	return
}
	