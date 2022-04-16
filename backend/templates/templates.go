package templates

import (
	"embed"
	"html/template"
	"io"
)

//go:embed *.html
var templateFS embed.FS

type Template struct {
	funcMap   template.FuncMap
	templates *template.Template
}

func New(funcMap template.FuncMap) *Template {
	templates := template.Must(template.New("").Funcs(funcMap).ParseFS(templateFS, "*.html"))
	return &Template{
		funcMap:   funcMap,
		templates: templates,
	}
}

func NewDebug(funcMap template.FuncMap) *Template {
	templates := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
	return &Template{
		funcMap:   funcMap,
		templates: templates,
	}
}

var DebugTemplates = true

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	if DebugTemplates {
		return NewDebug(t.funcMap).templates.ExecuteTemplate(w, name, data)
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
