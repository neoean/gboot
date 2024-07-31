package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"step2money-server/common/config"
	"step2money-server/common/db"
	"step2money-server/common/logs"
	"step2money-server/common/qiniu"
	"step2money-server/common/redis"
	"step2money-server/common/weApp"
	"step2money-server/router"
)

func main() {
	// config init
	config.InitConfig()

	// log init
	logs.LogInit()

	// db init
	db.Init()

	// redis init
	redis.Init()

	// weApp init
	weApp.InitWeApp()

	// qiniu init
	qiniu.Init()

	// gin init
	engine := gin.Default()

	// route init
	router.Init(engine)

	// listen init
	port := fmt.Sprintf("%v", config.Config.Port)
	err := engine.Run(":" + port)
	if err != nil {
		logs.Error("run webserver error %v", err)
		panic(err)
	}
}
