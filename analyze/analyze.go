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
	fmt.Println("Search...", filex)
	files, err := filepath.Glob(filex + "/*.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)

	a := NewAnalyzeData()

	for _, file := range files {
		analyze(file, a)
	}

	file, _ := json.MarshalIndent(a, "", "  ")
	_ = ioutil.WriteFile("analyze.json", file, 0644)

	createMetadata(a)
}

func Generate() {
	data, err := ioutil.ReadFile("analyze.json")
	if err != nil {
		return
	}

	a := NewAnalyzeData()
	json.Unmarshal(data, a)
	createMetadata(a)
}

type FieldData struct {
	IsString bool
	Multiple bool
	Base     bool
	Plus     bool
}

type AnalyzeData struct {
	// Types  map[string]bool
	Fields map[string]*FieldData
}

func NewAnalyzeData() *AnalyzeData {
	return &AnalyzeData{
		// Types:  make(map[string]bool, 0),
		Fields: make(map[string]*FieldData, 0),
	}
}

func (a *AnalyzeData) GetField(s string) *FieldData {
	if f, ok := a.Fields[s]; ok {
		return f
	} else {
		f := &FieldData{}
		a.Fields[s] = f
		return f
	}
}

func analyze(file string, a *AnalyzeData) error {
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	plus := strings.HasSuffix(file, "_plus.xml")

	fmt.Println("Successfully Opened", file)
	defer xmlFile.Close()

	converter := util.NewConvertReader(xmlFile)

	return analyzeElement(xml.NewDecoder(converter), a, make([]string, 0), plus)
}

const PATH_SEPARATOR = "|"

func analyzeElement(d *xml.Decoder, a *AnalyzeData, path []string, plus bool) error {
	if len(path) > 1 {
		s := strings.Join(path, PATH_SEPARATOR)
		a.Fields[s] = true
		if plus {
			a.Plus[s] = true
		} else {
			a.Base[s] = true
		}
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

			a.Types[strings.Join(path, PATH_SEPARATOR)] = true

			newPath := append(path, t.Name.Local)

			if _, ok := fields[t.Name.Local]; ok {
				a.Multiple[strings.Join(newPath, PATH_SEPARATOR)] = true
			}
			fields[t.Name.Local] = true

			analyzeElement(d, a, newPath, plus)

		case xml.CharData:
			data = append(data, t...)

		case xml.EndElement:
			if value {
				if _, err := strconv.Atoi(string(data)); err != nil {
					a.IsString[strings.Join(path, PATH_SEPARATOR)] = true
				}
			}

			if t.Name.Local == "type" {
				path[len(path)-2] = path[len(path)-2] + "+" + strcase.ToCamel(string(data))
			}

			return nil
		}
	}
	return nil
}

func filterSubtypes(data map[string]bool) []string {
	allowed := map[string]bool{
		"df_world|historical_events|historical_event":                       true,
		"df_world|historical_event_collections|historical_event_collection": true,
	}

	filtered := make(map[string]bool)
	for k, v := range data {
		if !v {
			continue
		}
		path := strings.Split(k, PATH_SEPARATOR)
		for index, seg := range path {
			if strings.Contains(seg, "+") {
				base := seg[:strings.Index(seg, "+")]
				basePath := strings.Join(append(path[:index], base), PATH_SEPARATOR)
				if allowed[basePath] {
					path[index] = seg
				}
			}
		}
		filtered[strings.Join(path, PATH_SEPARATOR)] = true
	}
	list := util.Keys(filtered)
	sort.Strings(list)
	return list
}

func createMetadata(a *AnalyzeData) {
	ts := filterSubtypes(a.Types)
	fs := filterSubtypes(a.Fields)

	// for _, s := range fs {
	// 	fmt.Println(s)
	// }

	objects := make(map[string]Object, 0)

	for _, k := range ts {
		if ok, _ := isArray(k, fs); !ok {
			n := k
			if strings.Contains(k, PATH_SEPARATOR) {
				n = k[strings.LastIndex(k, PATH_SEPARATOR)+1:]
			}

			if n == "" {
				continue
			}

			objFields := make(map[string]Field, 0)

			for _, f := range fs {
				if strings.HasPrefix(f, k+PATH_SEPARATOR) {
					fn := f[len(k)+1:]
					if !strings.Contains(fn, PATH_SEPARATOR) {
						legend := ""
						if a.Base[f] && a.Plus[f] {
							legend = "both"
						} else if a.Base[f] {
							legend = "base"
						} else if a.Plus[f] {
							legend = "plus"
						}

						field := Field{
							Name:     strcase.ToCamel(fn),
							Type:     "int",
							Multiple: a.Multiple[f],
							Legend:   legend,
						}
						if ok, elements := isArray(f, fs); ok {
							el := elements[strings.LastIndex(elements, PATH_SEPARATOR)+1:]
							fmt.Println(f + PATH_SEPARATOR + elements + PATH_SEPARATOR + "id")
							if a.Fields[elements+PATH_SEPARATOR+"id"] {
								field.Type = "map"
							} else {
								field.Type = "array"
							}
							field.ElementType = &(el)
						} else if ok, _ := isObject(f, fs); ok {
							field.Type = "object"
						} else if a.IsString[f] {
							field.Type = "string"
						}
						objFields[fn] = field
					}
				}
			}

			objects[n] = Object{
				Name:   strcase.ToCamel(n),
				Id:     a.Fields[k+PATH_SEPARATOR+"id"],
				Named:  a.Fields[k+PATH_SEPARATOR+"name"],
				Typed:  a.Fields[k+PATH_SEPARATOR+"type"],
				Fields: objFields,
			}
		}
	}

	file, _ := json.MarshalIndent(objects, "", "  ")
	_ = ioutil.WriteFile("model.json", file, 0644)

	f, err := os.Create("df/model.go")
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
		if !strings.HasPrefix(t, typ+PATH_SEPARATOR) {
			continue
		}
		f := t[len(typ)+1:]
		if strings.Contains(f, PATH_SEPARATOR) {
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
		if !strings.HasPrefix(t, typ+PATH_SEPARATOR) {
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
	Legend      string  `json:"legend"`
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
		t = "[]*" + objects[*f.ElementType].Name
	}
	if f.Type == "map" {
		t = "map[int]*" + objects[*f.ElementType].Name
	}
	if f.Type == "object" {
		t = f.Name
	}
	j := fmt.Sprintf("`json:\"%s\" legend:\"%s\"`", strcase.ToLowerCamel(f.Name), f.Legend)
	return fmt.Sprintf("%s %s%s %s", n, m, t, j)
}

func (f Field) StartAction() string {
	n := f.Name

	if n == "Id" || n == "Name" {
		n = n + "_"
	}

	if f.Type == "object" {
		p := fmt.Sprintf("v := %s{}\nv.Parse(d, &t)", f.Name)
		if !f.Multiple {
			return fmt.Sprintf("%s\nobj.%s = v", p, n)
		} else {
			return fmt.Sprintf("%s\nobj.%s = append(obj.%s, v)", p, n, n)
		}
	}

	if f.Type == "array" || f.Type == "map" {
		el := strcase.ToCamel(*f.ElementType)
		gen := fmt.Sprintf("New%s", el)

		if f.Type == "array" {
			return fmt.Sprintf("parseArray(d, &obj.%s, %s)", f.Name, gen)
		}

		if f.Type == "map" {
			return fmt.Sprintf("obj.%s = make(map[int]*%s)\nparseMap(d, &obj.%s, %s)", f.Name, el, f.Name, gen)
		}
	}

	if f.Type == "int" || f.Type == "string" {
		return "data = nil"
	}

	return ""
}

func (f Field) EndAction() string {
	n := f.Name

	if n == "Id" || n == "Name" {
		n = n + "_"
	}

	if !f.Multiple {
		if f.Type == "int" {
			return fmt.Sprintf("obj.%s = n(data)", n)
		} else if f.Type == "string" {
			return fmt.Sprintf("obj.%s = string(data)", n)
		}
	} else {
		if f.Type == "int" {
			return fmt.Sprintf("obj.%s = append(obj.%s, n(data))", n, n)
		} else if f.Type == "string" {
			return fmt.Sprintf("obj.%s = append(obj.%s, string(data))", n, n)
		}
	}

	return ""
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package df

import (
	"encoding/xml"
	"strconv"
)

{{- range $name, $obj := .Objects }}
type {{ $obj.Name }} struct {
	{{- range $fname, $field := $obj.Fields }}
	{{ $field.TypeLine $.Objects }}
	{{- end }}
}

func New{{ $obj.Name }}() *{{ $obj.Name }} { return &{{ $obj.Name }}{} }
{{- if $obj.Id }}
func (x *{{ $obj.Name }}) Id() int { return x.Id_ }
{{- end }}
{{- if $obj.Named }}
func (x *{{ $obj.Name }}) Name() string { return x.Name_ }
{{- end }}



{{- end }}

// Parser

func n(d []byte) int {
	v, _ := strconv.Atoi(string(d))
	return v
}

{{- range $name, $obj := .Objects }}
func (obj *{{ $obj.Name }}) Parse(d *xml.Decoder, start *xml.StartElement) error {
	var data []byte

	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			{{- range $fname, $field := $obj.Fields }}
			case "{{ $fname }}":
				{{ $field.StartAction }}
			{{- end }}
			default:
				// fmt.Println("unknown field", t.Name.Local)
				d.Skip()
			}

		case xml.CharData:
			data = append(data, t...)

		case xml.EndElement:
			if t.Name.Local == start.Name.Local {
				return nil
			}

			switch t.Name.Local {
			{{- range $fname, $field := $obj.Fields }}
			case "{{ $fname }}":
				{{ $field.EndAction }}
			{{- end }}
			default:
				// fmt.Println("unknown field", t.Name.Local)
			}
		}
	}
}

{{- end }}
`))
