package df

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func AnalyzeStructure(filex string) error {
	fmt.Println("Search...", filex)
	files, err := filepath.Glob(filex + "/*-legends.xml")
	if err != nil {
		return err
	}
	fmt.Println(files)

	a := NewAnalyzeData()

	for _, file := range files {
		analyze(file, a)
	}

	return a.Save()
}

type FieldData struct {
	IsString bool
	Multiple bool
	Base     bool
	Plus     bool
}

func NewFieldData() *FieldData {
	return &FieldData{}
}

type AnalyzeData struct {
	Fields   map[string]*FieldData
	SubTypes map[string]*map[string]*Subtype
}

func NewAnalyzeData() *AnalyzeData {
	return &AnalyzeData{
		Fields:   make(map[string]*FieldData, 0),
		SubTypes: make(map[string]*map[string]*Subtype),
	}
}

func (a *AnalyzeData) Save() error {
	file, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("analyze.json", file, 0644)
}

func LoadAnalyzeData() (*AnalyzeData, error) {
	data, err := ioutil.ReadFile("analyze.json")
	if err != nil {
		return nil, err
	}

	a := NewAnalyzeData()
	json.Unmarshal(data, a)
	return a, nil
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

func (a *AnalyzeData) GetSubType(s string, t string) *Subtype {
	var (
		st *map[string]*Subtype
		su *Subtype
		ok bool
	)

	if st, ok = a.SubTypes[s]; !ok {
		x := make(map[string]*Subtype)
		a.SubTypes[s] = &x
		st = &x
	}

	if su, ok = (*st)[t]; !ok {
		x := Subtype{Name: t}
		(*st)[t] = &x
		su = &x
	}

	return su
}

type AnalyzeContext struct {
	file     string
	plus     bool
	subtypes map[string]map[int]string
}

func analyze(file string, a *AnalyzeData) error {

	ctx := AnalyzeContext{
		file:     file,
		plus:     false,
		subtypes: make(map[string]map[int]string),
	}

	// base file

	xmlFile, err := os.Open(file)
	if err != nil {
		return err
	}

	fmt.Println("Successfully Opened", file)
	defer xmlFile.Close()

	_, err = analyzeElement(xml.NewDecoder(util.NewConvertReader(xmlFile)), a, make([]string, 0), &ctx)
	if err != nil {
		return err
	}

	// plus file

	ctx.plus = true
	file = strings.Replace(file, "-legends.xml", "-legends_plus.xml", 1)
	xmlFile, err = os.Open(file)
	if err != nil {
		return err
	}

	fmt.Println("Successfully Opened", file)
	defer xmlFile.Close()

	_, err = analyzeElement(xml.NewDecoder(util.NewConvertReader(xmlFile)), a, make([]string, 0), &ctx)

	return err
}

const PATH_SEPARATOR = "|"

type Value struct {
	Name  string
	Value string
}

func analyzeElement(d *xml.Decoder, a *AnalyzeData, path []string, ctx *AnalyzeContext) (*Value, error) {
	if len(path) > 1 {
		s := strings.Join(path, PATH_SEPARATOR)
		fd := a.GetField(s)
		if ctx.plus {
			fd.Plus = true
		} else {
			fd.Base = true
		}
	}

	var (
		data         []byte
		id           int
		idFound      bool
		subtype      string
		subtypeFound bool
	)
	value := true

	fields := make(map[string]bool)

Loop:
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break Loop
		} else if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			value = false

			newPath := append(path, t.Name.Local)

			if _, ok := fields[t.Name.Local]; ok {
				a.GetField(strings.Join(newPath, PATH_SEPARATOR)).Multiple = true
			}
			fields[t.Name.Local] = true

			v, err := analyzeElement(d, a, newPath, ctx)
			if err != nil {
				return nil, err
			}
			if v != nil {
				if v.Name == "id" {
					idFound = true
					id, _ = strconv.Atoi(v.Value)
				}
				if v.Name == "type" {
					subtypeFound = true
					subtype = v.Value

					if idFound && subtypeFound {
						p := strings.Join(path, PATH_SEPARATOR)
						if strings.Contains(p, "+") {
							p = p[:strings.LastIndex(p, "+")]
						}
						typeMap, ok := ctx.subtypes[p]
						if !ok {
							typeMap = make(map[int]string)
							ctx.subtypes[p] = typeMap
						}
						if !ctx.plus {
							typeMap[id] = subtype
							a.GetSubType(p, strcase.ToCamel(subtype)).BaseType = subtype
						} else {
							if typeMap[id] != subtype {
								if typeMap[id] != "" {
									a.GetSubType(p, strcase.ToCamel(typeMap[id])).PlusType = subtype
								} else {
									a.GetSubType(p, strcase.ToCamel(subtype)).PlusType = subtype
								}
								subtype = typeMap[id]
							} else {
								a.GetSubType(p, strcase.ToCamel(subtype)).PlusType = subtype
							}
						}
					}

					path[len(path)-1] = path[len(path)-1] + "+" + strcase.ToCamel(subtype)
				}
			}

		case xml.CharData:
			data = append(data, t...)

		case xml.EndElement:
			if value {
				s := string(data)
				if _, err := strconv.Atoi(s); err != nil {
					a.GetField(strings.Join(path, PATH_SEPARATOR)).IsString = true
				}
				return &Value{Name: t.Name.Local, Value: s}, nil
			}

			return nil, err
		}
	}
	return nil, nil
}
