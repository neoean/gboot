package gboot

import "flag"

var (
	bindAddr = flag.String("bindAddr", ":8090", "http listen address")
	dir      = flag.String("dir", "project-demo/", "project directory")
	dsn      = flag.String("dsn", "", "database schema name")
)

func main() {
	// flog
	flag.Parse()

	//
}
