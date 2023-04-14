package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kim118000/game2023/pkg/constant"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  constant.GetMsg(errCode),
		Data: data,
	})
	return
}
