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
	files, err := filepath.Glob(filex + "/*.xml")
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
	Fields map[string]*FieldData
}

func NewAnalyzeData() *AnalyzeData {
	return &AnalyzeData{
		Fields: make(map[string]*FieldData, 0),
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
		fd := a.GetField(s)
		if plus {
			fd.Plus = true
		} else {
			fd.Base = true
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

			newPath := append(path, t.Name.Local)

			if _, ok := fields[t.Name.Local]; ok {
				a.GetField(strings.Join(newPath, PATH_SEPARATOR)).Multiple = true
			}
			fields[t.Name.Local] = true

			analyzeElement(d, a, newPath, plus)

		case xml.CharData:
			data = append(data, t...)

		case xml.EndElement:
			if value {
				if _, err := strconv.Atoi(string(data)); err != nil {
					a.GetField(strings.Join(path, PATH_SEPARATOR)).IsString = true
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
