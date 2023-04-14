package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kim118000/game2023/pkg/constant"
	"github.com/kim118000/game2023/pkg/util"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var claims *util.Claims
		var err error
		code = constant.SUCCESS
		token := c.GetHeader("token")
		if token == "" {
			code = constant.INVALID_PARAMS
		} else {
			claims, err = util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = constant.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = constant.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != constant.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  constant.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Set("username", claims.Username)

		c.Next()
	}
}
