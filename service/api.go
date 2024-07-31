package service

import (
	"github.com/neoean/gboot/common/db"
)

type Model struct {
	Package        string
	TableName      string
	ModelName      string
	LowerModelName string
}

func GenApi(projectName string) {
	tables, err := db.DB.Migrator().GetTables()
	if err != nil {
		panic(err)
	}

	var models []*Model
	for _, t := range tables {
		modelName := db.DB.Config.NamingStrategy.SchemaName(t)
		lowerModelName := CamelToLowerCamel(modelName)
		model := &Model{
			Package:        projectName,
			TableName:      t,
			ModelName:      modelName,
			LowerModelName: lowerModelName,
		}
		models = append(models, model)

		// handler tmpl
		GenGoTemplate("../template/api/handler.go.template", projectName+"/router/"+lowerModelName, "handler.go", model)
	}

	// root router tmpl
	GenGoTemplate("../template/api/root.go.template", projectName+"/router", "root.go", models)

	// base
	GenGoTemplate("../template/api/base.go.template", projectName+"/router", "root.go", models)

}
