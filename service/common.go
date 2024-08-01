package service

func GenCommon(projectName, user, password, host, dbName string) {
	CommonConfig := struct {
		Package  string
		Host     string
		Username string
		Password string
		DbName   string
	}{projectName, host, user, password, dbName}
	// config template
	GenGoTemplate("./template/config/dev_conf.go.template", projectName+"/conf", "dev.yml", CommonConfig)
	GenGoTemplate("./template/config/dev_conf.go.template", projectName+"/conf", "prod.yml", CommonConfig)

	// go.mod
	GenGoTemplate("./template/go.mod.template", projectName+"/", "go.mod", CommonConfig)

}
