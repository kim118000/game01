package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kim118000/game2023/pkg/app"
	"github.com/kim118000/game2023/pkg/constant"
	"github.com/kim118000/game2023/services/bullet_service"
	"github.com/unknwon/com"
	"net/http"
	"time"
)

type AddBulletForm struct {
	Content string `valid:"Required;MinSize(1);MaxSize(40)"`
}

// @Summary Add bullet
// @Produce  json
// @Param content body string true "Content"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/add_bullet [post]
func AddBullet(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddBulletForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	bulletService := bullet_service.Bullet{
		Ts:       time.Now().UnixMilli(),
		Content:  form.Content,
		Username: c.GetString("username"),
	}

	err := bulletService.AddBullet()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Summary Get Top bullet
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/get_top_bullet [get]
func GetTopBullet(c *gin.Context) {
	appG := app.Gin{C: c}
	ts := c.Query("ts")
	list := bullet_service.GetTopBullet()

	if ts == "" {
		appG.Response(http.StatusOK, constant.SUCCESS, list)
		return
	}

	min, _ := com.StrTo(ts).Int64()

	bullets := make([]*bullet_service.Bullet, 0, len(list))

	for _, b := range list {
		if b.Ts > min {
			bullets = append(bullets, b)
		}
	}

	appG.Response(http.StatusOK, constant.SUCCESS, bullets)
}
