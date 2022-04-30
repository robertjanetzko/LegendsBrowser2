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

	"github.com/cheggaaa/pb/v3"
	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func AnalyzeStructure(filex string, a *AnalyzeData) error {
	fmt.Println("Search...", filex)
	files, err := filepath.Glob(filex + "/*-legends.xml")
	if err != nil {
		return err
	}
	fmt.Println(files)

	if a == nil {
		a = NewAnalyzeData()
	}

	for _, file := range files {
		analyze(file, a)
	}

	return a.Save()
}

type FieldData struct {
	IsString bool
	NoBool   bool
	Multiple bool
	Base     bool
	Plus     bool
	Values   map[string]bool
	Enum     bool
}

func NewFieldData() *FieldData {
	return &FieldData{
		Enum:   true,
		Values: make(map[string]bool),
	}
}

type AnalyzeData struct {
	Fields     map[string]*FieldData
	SubTypes   map[string]*map[string]*Subtype
	Overwrites *Overwrites `json:"-"`
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

	data, err = ioutil.ReadFile("overwrites.json")
	if err != nil {
		return nil, err
	}

	overwrites := &Overwrites{}
	json.Unmarshal(data, overwrites)
	a.Overwrites = overwrites

	return a, nil
}

func (a *AnalyzeData) GetField(s string) *FieldData {
	if f, ok := a.Fields[s]; ok {
		return f
	} else {
		f := NewFieldData()
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
	file       string
	plus       bool
	subtypes   map[string]map[int]string
	overwrites *Overwrites
}

type AdditionalField struct {
	Name string
	Type string
}

type Overwrites struct {
	ForceEnum        map[string]bool
	AdditionalFields map[string][]AdditionalField
}

func analyze(file string, a *AnalyzeData) error {
	data, err := ioutil.ReadFile("overwrites.json")
	if err != nil {
		return err
	}

	overwrites := &Overwrites{}
	json.Unmarshal(data, overwrites)

	ctx := AnalyzeContext{
		file:       file,
		plus:       false,
		subtypes:   make(map[string]map[int]string),
		overwrites: overwrites,
	}

	// base file

	fi, err := os.Stat(file)
	if err != nil {
		return err
	}
	size := fi.Size()
	bar := pb.Full.Start64(size)

	xmlFile, err := os.Open(file)
	if err != nil {
		return err
	}

	fmt.Println("\nAnalyzing", file)
	defer xmlFile.Close()

	converter := util.NewConvertReader(xmlFile)
	barReader := bar.NewProxyReader(converter)

	_, err = analyzeElement(xml.NewDecoder(barReader), a, make([]string, 0), &ctx)
	if err != nil {
		return err
	}

	bar.Finish()

	// plus file

	ctx.plus = true
	file = strings.Replace(file, "-legends.xml", "-legends_plus.xml", 1)

	fi, err = os.Stat(file)
	if err != nil {
		return err
	}
	size = fi.Size()
	bar = pb.Full.Start64(size)

	xmlFile, err = os.Open(file)
	if err != nil {
		return err
	}

	fmt.Println("\nAnalyzing", file)
	defer xmlFile.Close()

	converter = util.NewConvertReader(xmlFile)
	barReader = bar.NewProxyReader(converter)

	_, err = analyzeElement(xml.NewDecoder(barReader), a, make([]string, 0), &ctx)

	bar.Finish()

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

					if allowedTyped[strings.Join(path, PATH_SEPARATOR)] {
						path[len(path)-1] = path[len(path)-1] + "+" + strcase.ToCamel(subtype)
					}
				}
			}

		case xml.CharData:
			data = append(data, t...)

		case xml.EndElement:
			if value {
				s := strings.TrimSpace(string(data))
				if _, err := strconv.Atoi(s); err != nil {
					f := a.GetField(strings.Join(path, PATH_SEPARATOR))

					force := ctx.overwrites.ForceEnum[strings.Join(path, PATH_SEPARATOR)]
					if force {
						f.Enum = true
					}

					f.IsString = true
					if s != "" && f.Enum {
						f.Values[s] = true
					}

					if len(f.Values) > 30 && !force {
						f.Values = make(map[string]bool)
						f.Enum = false
					}
				}
				if len(s) > 0 {
					a.GetField(strings.Join(path, PATH_SEPARATOR)).NoBool = true
				}
				return &Value{Name: t.Name.Local, Value: s}, nil
			}

			return nil, err
		}
	}
	return nil, nil
}
