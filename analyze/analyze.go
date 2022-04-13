package analyze

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"legendsbrowser/util"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func Analyze(filex string) {
	files, err := filepath.Glob("*.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)

	files = []string{filex}

	a := NewAnalyzeData()

	for _, file := range files {
		xmlFile, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Successfully Opened", file)
		defer xmlFile.Close()

		converter := util.NewConvertReader(xmlFile)
		analyze(converter, a)
	}

	createMetadata(a)
}

type analyzeData struct {
	path     []string
	types    *map[string]bool
	fields   *map[string]bool
	isString *map[string]bool
	multiple *map[string]bool
}

func NewAnalyzeData() analyzeData {
	path := make([]string, 0)
	types := make(map[string]bool, 0)
	fields := make(map[string]bool, 0)
	isString := make(map[string]bool, 0)
	multiple := make(map[string]bool, 0)

	return analyzeData{
		path:     path,
		types:    &types,
		fields:   &fields,
		isString: &isString,
		multiple: &multiple,
	}
}

func analyzeElement(d *xml.Decoder, a analyzeData) error {
	if len(a.path) > 1 {
		(*a.fields)[strings.Join(a.path, ">")] = true
	}

	var (
		data []byte
	)
	value := true

	fields := make(map[string]bool)

Loop:
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break Loop
		} else if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			value = false

			(*a.types)[strings.Join(a.path, ">")] = true

			a2 := a
			a2.path = append(a.path, t.Name.Local)

			if _, ok := fields[t.Name.Local]; ok {
				(*a.multiple)[strings.Join(a2.path, ">")] = true
			}
			fields[t.Name.Local] = true

			analyzeElement(d, a2)

		case xml.CharData:
			data = append(data, t...)

		case xml.EndElement:
			if value {
				if _, err := strconv.Atoi(string(data)); err != nil {
					(*a.isString)[strings.Join(a.path, ">")] = true
				}
			}

			if t.Name.Local == "type" {
				a.path[len(a.path)-2] = a.path[len(a.path)-2] + strcase.ToCamel(string(data))
				fmt.Println(a.path)
			}

			return nil
		}
	}
	return nil
}

func analyze(r io.Reader, a analyzeData) error {
	d := xml.NewDecoder(r)
	return analyzeElement(d, a)
}

func createMetadata(a analyzeData) {
	ts := util.Keys(*a.types)
	sort.Strings(ts)

	fs := util.Keys(*a.fields)
	sort.Strings(fs)

	objects := make(map[string]Object, 0)

	for _, k := range ts {
		if ok, _ := isArray(k, fs); !ok {
			n := k
			if strings.Contains(k, ">") {
				n = k[strings.LastIndex(k, ">")+1:]
			}

			if n == "" {
				continue
			}

			objFields := make(map[string]Field, 0)

			fmt.Println("\n", n)
			for _, f := range fs {
				if strings.HasPrefix(f, k+">") {
					fn := f[len(k)+1:]
					if !strings.Contains(fn, ">") {
						fmt.Println("     ", fn)

						if ok, elements := isArray(f, fs); ok {
							el := elements[strings.LastIndex(elements, ">")+1:]
							objFields[fn] = Field{
								Name:        strcase.ToCamel(fn),
								Type:        "array",
								ElementType: &(el),
							}
						} else if ok, _ := isObject(f, fs); ok {
							objFields[fn] = Field{
								Name:     strcase.ToCamel(fn),
								Type:     "object",
								Multiple: (*a.multiple)[f],
							}
						} else if (*a.isString)[f] {
							objFields[fn] = Field{
								Name:     strcase.ToCamel(fn),
								Type:     "string",
								Multiple: (*a.multiple)[f],
							}
						} else {
							objFields[fn] = Field{
								Name:     strcase.ToCamel(fn),
								Type:     "int",
								Multiple: (*a.multiple)[f],
							}
						}
					}
				}
			}

			objects[n] = Object{
				Name:   strcase.ToCamel(n),
				Id:     (*a.fields)[k+">id"],
				Named:  (*a.fields)[k+">name"],
				Typed:  (*a.fields)[k+">type"],
				Fields: objFields,
			}
		}
	}

	file, _ := json.MarshalIndent(objects, "", "  ")
	_ = ioutil.WriteFile("model.json", file, 0644)

	f, err := os.Create("contributors.go")
	defer f.Close()

	err = packageTemplate.Execute(f, struct {
		Objects map[string]Object
	}{
		Objects: objects,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func isArray(typ string, types []string) (bool, string) {
	fc := 0
	elements := ""

	for _, t := range types {
		if !strings.HasPrefix(t, typ+">") {
			continue
		}
		f := t[len(typ)+1:]
		if strings.Contains(f, ">") {
			continue
		}
		fc++
		elements = t
	}
	return fc == 1, elements
}

func isObject(typ string, types []string) (bool, string) {
	fc := 0

	for _, t := range types {
		if !strings.HasPrefix(t, typ+">") {
			continue
		}
		fc++
	}
	return fc > 0, typ
}

type Object struct {
	Name   string           `json:"name"`
	Id     bool             `json:"id,omitempty"`
	Named  bool             `json:"named,omitempty"`
	Typed  bool             `json:"typed,omitempty"`
	Fields map[string]Field `json:"fields"`
}

type Field struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Multiple    bool    `json:"multiple,omitempty"`
	ElementType *string `json:"elements,omitempty"`
}

func (f Field) TypeLine(objects map[string]Object) string {
	n := f.Name

	if n == "Id" || n == "Name" {
		n = n + "_"
	}

	m := ""
	if f.Multiple {
		m = "[]"
	}
	t := f.Type
	if f.Type == "array" {
		t = "map[int]*" + objects[*f.ElementType].Name
	}
	if f.Type == "object" {
		t = f.Name
	}
	j := "`json:\"" + strcase.ToLowerCamel(f.Name) + "\"`"
	return fmt.Sprintf("%s %s%s %s", n, m, t, j)
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package generate

{{- range $name, $obj := .Objects }}
type {{ $obj.Name }} struct {
	{{- range $fname, $field := $obj.Fields }}
	{{ $field.TypeLine $.Objects }}
	{{- end }}
}

{{- if $obj.Id }}
func (x *{{ $obj.Name }}) Id() int { return x.Id_ }
{{- end }}
{{- if $obj.Named }}
func (x *{{ $obj.Name }}) Name() string { return x.Name_ }
{{- end }}

{{- end }}
`))
