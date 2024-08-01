package service

import (
	"os"
	"path/filepath"
	"strings"
)

type CommonConfig struct {
	Package  string
	Host     string
	Username string
	Password string
	DbName   string
}

func GenCommon(projectName, user, password, host, dbName string) {
	commonConfig := &CommonConfig{projectName, host, user, password, dbName}

	// config template
	GenGoTemplate("./template/config/dev_conf.go.template", projectName+"/conf", "dev.yml", commonConfig)
	GenGoTemplate("./template/config/dev_conf.go.template", projectName+"/conf", "prod.yml", commonConfig)

	// go.mod
	GenGoTemplate("./template/go.mod.template", projectName+"/", "go.mod", commonConfig)

	// common utils
	genCommonTemplate(projectName, "./template/common", commonConfig)
	genCommonTemplate(projectName, "./template/cmd", commonConfig)
}

func genCommonTemplate(projectName, templatePath string, CommonConfig *CommonConfig) {
	err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileName := info.Name()

			filePath := strings.Replace(path, "template", "", 1)
			filePath = strings.Replace(filePath, fileName, "", 1)

			GenGoTemplate(path, projectName+filePath, fileName, CommonConfig)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
