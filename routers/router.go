package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kim118000/game2023/middleware/jwt"
	"github.com/kim118000/game2023/routers/api"
	v1 "github.com/kim118000/game2023/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth", api.Auth)
	r.POST("/register", api.Register)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/add_bullet", v1.AddBullet)
		apiv1.GET("/get_top_bullet", v1.GetTopBullet)
	}

	return r
}
