package main

import (
	"flag"
	"github.com/neoean/gboot/common/db"
	"github.com/neoean/gboot/service"
)

var (
	bindAddr = flag.String("bindAddr", ":8090", "http listen address")
	dir      = flag.String("dir", "project-demo/", "project directory")

	dbType   = flag.String("dbtype", "mysql", "database type")
	dsn      = flag.String("dsn", "", "database schema name")
	userName = flag.String("username", "root", "database username")
	password = flag.String("password", "123456", "database password")
	host     = flag.String("host", "127.0.0.1:3306", "database host")
	dbName   = flag.String("dbname", "test", "database name")
)

func main() {
	// flog
	flag.Parse()

	// init db
	db.InitMysql(*dsn, *userName, *password, *host, *dbName)

	// gorm gen
	service.GenModel(*dir)

	// api gen
	service.GenApi(*dir)

	// common gen
	service.GenCommon(*dir, *userName, *password, *host, *dbName)
}
