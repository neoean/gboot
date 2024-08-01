package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"{{.Package}}/common/config"
	"{{.Package}}/common/db"
	"{{.Package}}/common/logs"
	"{{.Package}}/common/qiniu"
	"{{.Package}}/common/redis"
	"{{.Package}}/common/weApp"
	"{{.Package}}/router"
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
