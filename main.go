package main

import (
	"flag"
	"github.com/neoean/gboot/common/db"
	"github.com/neoean/gboot/service"
)

var (
	name     = flag.String("name", "yunyu-server", "project name")
	dbType   = flag.String("dbtype", "mysql", "database type")
	dsn      = flag.String("dsn", "", "database schema name")
	userName = flag.String("username", "root", "database username")
	password = flag.String("password", "123456", "database password")
	host     = flag.String("host", "127.0.0.1:3306", "database host")
	dbName   = flag.String("dbname", "yunyu_server", "database name")
)

func main() {
	// flog
	flag.Parse()

	// init db
	db.InitMysql(*dsn, *userName, *password, *host, *dbName)

	// gorm gen
	service.GenModel(*name)

	// api gen
	service.GenApi(*name)

	// common gen
	service.GenCommon(*name, *userName, *password, *host, *dbName)
}
