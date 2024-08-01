package main

import (
	"{{.Package}}/common/config"
	"{{.Package}}/common/db"
	"{{.Package}}/common/logs"
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
