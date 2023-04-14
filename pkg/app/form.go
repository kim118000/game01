package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/kim118000/game2023/pkg/constant"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, constant.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, constant.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, constant.INVALID_PARAMS
	}

	return http.StatusOK, constant.SUCCESS
}
