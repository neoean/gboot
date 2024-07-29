package service

import "github.com/neoean/gboot/common/db"

func GenApi(dir string) {
	tables, err := db.DB.Migrator().GetTables()
	if err != nil {
		panic(err)
	}
	//
	//for _, t := range tables {
	//	schemaName := db.DB.Config.NamingStrategy.SchemaName(t)
	//}
}
