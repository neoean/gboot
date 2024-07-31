package main

import (
	"step2money-server/common/config"
	"step2money-server/common/db"
	"step2money-server/common/logs"
)

func main() {
	// config init
	config.InitConfig()

	// log init
	logs.LogInit()

	// db init
	db.Init()

	// gen
	db.InitGen()
}
