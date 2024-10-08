package generator

var handlerTemplate = `// Code generated by "templplaygroundgenerator"; DO NOT EDIT.
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/jfbus/templ-components-playground/templplayground/views/components"
{{- range . }}
	"{{.ImportPath}}"
{{- end }}
)

var binder = &echo.DefaultBinder{}
func ViewComponent(c echo.Context) error {
	switch c.Param("component") {
{{- range . }}
		case "{{.Name}}":
			def := {{.Package}}.D{}
			err := Bind(c.Request(), &def)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			components.{{.Name}}Viewer(def).Render(c.Request().Context(), c.Response().Writer)
{{- end }}
	}
	return nil 
}
`

var mainTemplate = `// Code generated by "templplaygroundgenerator"; DO NOT EDIT.
package views

import (
	"github.com/jfbus/templ-components/components/accordion"
	"github.com/jfbus/templ-components/components/accordion/element"
	"github.com/jfbus/templ-components-playground/templplayground/views/components"
)

templ Main() {
	@accordion.C(accordion.D{
		ID: "templPlayground",
		Children: []element.D{
{{- range . }}
			{
				Title: "{{.Name}}",
				Content: components.{{.Name}}Section(),
			},
{{- end}}
		},
	})
}
`

var componentTemplate = `// Code generated by "templplaygroundgenerator"; DO NOT EDIT.
package components

import (
{{- range .ImportPaths }}
	"{{.}}"
{{- end }}
)

templ {{.Name}}Viewer(def {{.Package}}.D) {
	@{{.Package}}.C(def)
}

templ {{.Name}}Form() {
{{range .Fields}}
{{- if eq .InputType "select" }}
	@selectfield.C(selectfield.D{
		Name: "{{.Name}}",
		Label: "{{.Name}}",
		Options: []option.D{
			{
				Label: "Select a value",
			},
{{- range .Values}}
			{
				Label: "{{.Package}}.{{.Name}}",
				Value: {{.Value}},
			},
{{- end }}
		},
	})
{{- end }}
{{- if eq .InputType "input" }}
	@input.C(input.D{
		Name: "{{.Name}}",
		Label: "{{.Name}}",
{{- if .Default}}
		Value: {{.Default}},
{{- end }}
	})
{{- end }}
{{- if eq .InputType "checkbox" }}
	@checkbox.C(checkbox.D{
		Name: "{{.Name}}",
		Label: "{{.Name}}",
		Value: "true",
	})
{{- end }}
{{- end }}
}

templ {{.Name}}Section() {
	@ComponentViewer(Component{
		ID:     "{{.Name}}",
		Form:   {{.Name}}Form(),
		Viewer: {{.Name}}Viewer({{.Package}}.D{
{{- range .Fields}}{{if .Default}}
			{{.Name}}:{{.Default}},
{{- end}}{{end}}
		}),
	})
}
`
