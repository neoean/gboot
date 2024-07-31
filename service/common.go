package service

import (
	"bytes"
	"github.com/neoean/gboot/common"
	"github.com/neoean/gboot/template/config"
	"os"
	"path/filepath"
)

func GenCommon(dir, user, password, host, dbName string) {
	// config template
	configTemplate, err := buildTemplate(config.ConfigTmplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = configTemplate.Execute(&buf, struct {
		Host     string
		Username string
		Password string
		DbName   string
	}{host, user, password, dbName})
	if err != nil {
		panic(err)
	}

	if !common.IsDirExists(dir + "conf") {
		err = os.Mkdir(dir+"conf", 0777)
		if err != nil {
			panic(err)
		}
	}

	err = os.WriteFile(filepath.Join(dir+"conf", "dev.yml"), buf.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filepath.Join(dir+"conf", "prod.yml"), buf.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
}
