package service

import (
	"bytes"
	"fmt"
	"github.com/neoean/gboot/common"
	"github.com/neoean/gboot/common/db"
	"github.com/neoean/gboot/template/api"
	"go/format"
	"os"
	"path/filepath"
)

type Model struct {
	TableName string
	ModelName string
}

func GenApi(dir string) {
	tables, err := db.DB.Migrator().GetTables()
	if err != nil {
		panic(err)
	}

	if !common.IsDirExists(dir + "router") {
		err = os.Mkdir(dir+"router", 0777)
		if err != nil {
			fmt.Printf("ERROR: mkdir error: %v", err)
			panic(err)
		}
	}

	var models []*Model
	for _, t := range tables {
		modelName := db.DB.Config.NamingStrategy.SchemaName(t)
		models = append(models, &Model{
			TableName: t,
			ModelName: modelName,
		})
	}

	routerTemplate, err := buildTemplate(api.RouterTmpl)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = routerTemplate.Execute(&buf, models)
	if err != nil {
		panic(err)
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println("Error in format source: " + err.Error())
		panic(err)
	}
	err = os.WriteFile(filepath.Join(dir+"router", "root.go"), data, 0777)
	if err != nil {
		panic(err)
	}
}
