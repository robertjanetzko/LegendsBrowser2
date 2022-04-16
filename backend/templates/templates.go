package templates

import (
	"embed"
	"fmt"
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
	ts := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
	return &Template{
		funcMap:   funcMap,
		templates: ts,
	}
}

var DebugTemplates = true

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	if DebugTemplates {
		fmt.Println("RENDER", name)
		tmpl := NewDebug(t.funcMap).templates
		tmpl = template.Must(tmpl.ParseFiles("templates/" + name))
		return tmpl.ExecuteTemplate(w, name, data)
	}
	tmpl := template.Must(t.templates.Clone())
	tmpl = template.Must(tmpl.ParseFS(templateFS, name))
	return tmpl.ExecuteTemplate(w, name, data)
}
