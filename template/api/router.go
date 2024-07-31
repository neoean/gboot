package api

var RouterTmpl = `package api

import (
	"github.com/gin-gonic/gin"
)

func ConfigRouter(r *gin.RouterGroup) {
	// api root
	apiRoot := r.Group("/api")

	// add jwt
	apiRoot.Use(middlewares.Jwt())

    {{range .}}
	// {{.ModelName}} group
	{{.ModelName}}Group(apiRoot.Group("/{{.TableName}}"))

    {{end}}
}

{{range .}}
func {{ .ModelName}}Group(g *gin.RouterGroup) {
	g.POST("/{{.TableName}}", Add{{.ModelName}})
	g.DELETE("/{{.TableName}}/:id", Delete{{.ModelName}})
	g.PUT("/{{.TableName}}/:id", Update{{.ModelName}})
	g.GET("/{{.TableName}}", GetAll{{.ModelName}})
	g.GET("/{{.TableName}}/:id", Get{{.ModelName}})
}

{{end}}
`
