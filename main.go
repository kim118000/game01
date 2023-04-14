package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kim118000/game2023/models"
	"github.com/kim118000/game2023/pkg/db"
	"github.com/kim118000/game2023/pkg/logging"
	"github.com/kim118000/game2023/pkg/redis"
	"github.com/kim118000/game2023/pkg/setting"
	"github.com/kim118000/game2023/pkg/util"
	"github.com/kim118000/game2023/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	db.Setup()
	logging.Setup()
	redis.Setup()
	util.Setup()

	models.AutoMigrate()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
