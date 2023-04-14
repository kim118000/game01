package models

import (
	"github.com/kim118000/game2023/pkg/db"
	"github.com/kim118000/game2023/pkg/util"
)

type User struct {
	Model
	UserName string `gorm:"column:username" json:"userName"`
	Password string `gorm:"column:password" json:"password"`
}

func (u *User) Register() error {
	user := User{
		UserName: u.UserName,
		Password: util.EncodeMD5(u.Password),
	}
	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) IsExist() bool {
	var user User
	db.DB.First(&user, "username = ?", u.UserName)
	if user.ID > 0 {
		return true
	}
	return false
}

func GetUser(username string) *User {
	var user User
	db.DB.First(&user, "username = ?", username)
	return &user
}
