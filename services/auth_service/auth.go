package auth_service

import (
	"github.com/kim118000/game2023/models"
	"github.com/kim118000/game2023/pkg/constant"
	"github.com/kim118000/game2023/pkg/util"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Auth() bool {
	user := models.GetUser(a.Username)
	pwd := util.EncodeMD5(a.Password)
	if user.Password == pwd {
		return true
	}
	return false
}

func (a *Auth) Register() error {
	u := &models.User{
		UserName: a.Username,
		Password: a.Password,
	}

	flag := u.IsExist()
	if flag {
		return constant.USER_EXIST_ERROR
	}

	return u.Register()
}
