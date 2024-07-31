package service

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/inflection"
	"github.com/neoean/gboot/common"
	"github.com/serenize/snaker"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenGoTemplate(tmpl, dir, name string, models any) {
	if !common.IsDirExists(dir) {
		err := os.Mkdir(dir, 0777)
		if err != nil {
			fmt.Printf("ERROR: mkdir error: %v", err)
			panic(err)
		}
	}

	routerTemplate, err := buildTemplate(tmpl)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = routerTemplate.Execute(&buf, models)
	if err != nil {
		panic(err)
	}

	data := buf.Bytes()
	if strings.HasSuffix(name, ".go") {
		data, err = format.Source(buf.Bytes())
		if err != nil {
			fmt.Println("Error in format source: " + err.Error())
			panic(err)
		}
	}
	err = os.WriteFile(filepath.Join(dir, name), data, 0777)
	if err != nil {
		panic(err)
	}
}

func buildTemplate(t string) (*template.Template, error) {
	var funcMap = template.FuncMap{
		"pluralize":        inflection.Plural,
		"title":            strings.Title,
		"toLower":          strings.ToLower,
		"toLowerCamelCase": CamelToLowerCamel,
		"toSnakeCase":      snaker.CamelToSnake,
	}

	tmpl, err := template.New("models").Funcs(funcMap).ParseFiles(t)

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func CamelToLowerCamel(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])

	return strings.Join(ss, "")
}
