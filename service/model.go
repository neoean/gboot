package service

import (
	"github.com/neoean/gboot/common/db"
	"gorm.io/gen"
)

// Querier Dynamic SQL
type Querier interface {
	// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func GenModel(dir string) {
	g := gen.NewGenerator(gen.Config{
		OutPath:      dir + "/dao",
		ModelPkgPath: dir + "/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		WithUnitTest: false,
	})

	// db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db.DB) // reuse your gorm db

	// models
	models := g.GenerateAllTable()

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(models...)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, models...)

	// Generate the code
	g.Execute()
}
