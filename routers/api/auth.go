package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kim118000/game2023/pkg/app"
	"github.com/kim118000/game2023/pkg/constant"
	"github.com/kim118000/game2023/pkg/setting"
	"github.com/kim118000/game2023/pkg/util"
	"github.com/kim118000/game2023/services/auth_service"
	"net/http"
)

type authForm struct {
	Username string `valid:"Required; AlphaNumeric; MinSize(6);MaxSize(26)"`
	Password string `valid:"Required; AlphaNumeric; MinSize(6)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /Auth [post]
func Auth(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form authForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := auth_service.Auth{Username: form.Username, Password: form.Password}
	isLogin := authService.Auth()
	if !isLogin {
		appG.Response(http.StatusUnauthorized, constant.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(form.Username, form.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, map[string]string{
		"token": token,
	})
}

// @Summary Get Register
// @Produce  json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /Register [post]
func Register(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form authForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := auth_service.Auth{Username: form.Username, Password: form.Password}
	err := authService.Register()
	if err == constant.USER_EXIST_ERROR {
		appG.Response(http.StatusBadRequest, constant.USER_EXIST, nil)
		return
	}

	if err != nil {
		appG.Response(http.StatusBadRequest, constant.INVALID_PARAMS, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

func UploadFile(c *gin.Context) {
	var appG = app.Gin{C: c}
	var fp = setting.AppSetting.FilePath
	form, err := c.MultipartForm()
	if err != nil {
		appG.Response(http.StatusBadRequest, constant.INVALID_PARAMS, nil)
	}

	//通过字段名映射
	f := form.File["file"]
	//for range遍历文件
	for _, file := range f {
		fmt.Println(file.Filename)
		c.SaveUploadedFile(file, fp+file.Filename)
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
