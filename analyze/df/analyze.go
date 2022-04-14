package df

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func Analyze(filex string) error {
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

func Generate() error {
	a, err := LoadAnalyzeData()
	if err != nil {
		return err
	}
	return createMetadata(a)
}

var allowedTyped = map[string]bool{
	"df_world|historical_events|historical_event":                       true,
	"df_world|historical_event_collections|historical_event_collection": true,
}

func filterSubtypes(data *map[string]*FieldData) []string {
	filtered := make(map[string]*FieldData)
	for k, v := range *data {
		path := strings.Split(k, PATH_SEPARATOR)
		for index, seg := range path {
			if strings.Contains(seg, "+") {
				base := seg[:strings.Index(seg, "+")]
				basePath := strings.Join(append(path[:index], base), PATH_SEPARATOR)
				if allowedTyped[basePath] {
					path[index] = seg
				}
			}
		}
		filtered[strings.Join(path, PATH_SEPARATOR)] = v
	}
	*data = filtered
	list := util.Keys(filtered)
	sort.Strings(list)
	return list
}

func getSubtypes(objectTypes []string, k string) *[]string {
	subtypes := make(map[string]bool)

	for _, t := range objectTypes {
		if strings.HasPrefix(t, k+"+") && !strings.Contains(t[len(k):], PATH_SEPARATOR) {
			subtypes[t[strings.LastIndex(t, "+")+1:]] = true
		}
	}

	keys := util.Keys(subtypes)
	sort.Strings(keys)

	if len(keys) > 0 {
		return &keys
	}

	return nil
}

func getSubtypeOf(k string) *string {
	if strings.Contains(k, PATH_SEPARATOR) {
		last := k[strings.LastIndex(k, PATH_SEPARATOR)+1:]
		if strings.Contains(last, "+") {
			base := strcase.ToCamel(last[:strings.Index(last, "+")])
			return &base
		}
	}
	return nil
}

func createMetadata(a *AnalyzeData) error {
	fs := filterSubtypes(&a.Fields)

	var objectTypes []string
	for k := range a.Fields {
		path := strings.Split(k, PATH_SEPARATOR)
		if len(path) >= 2 {
			objectTypes = append(objectTypes, strings.Join(path[:len(path)-1], PATH_SEPARATOR))
		}
	}

	objects := make(map[string]Object, 0)

	for _, k := range objectTypes {
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
						if a.Fields[f].Base && a.Fields[f].Plus {
							legend = "both"
						} else if a.Fields[f].Base {
							legend = "base"
						} else if a.Fields[f].Plus {
							legend = "plus"
						}

						field := Field{
							Name:     strcase.ToCamel(fn),
							Type:     "int",
							Multiple: a.Fields[f].Multiple,
							Legend:   legend,
						}
						if ok, elements := isArray(f, fs); ok {
							el := elements[strings.LastIndex(elements, PATH_SEPARATOR)+1:]
							if _, ok := a.Fields[elements+PATH_SEPARATOR+"id"]; ok {
								field.Type = "map"
							} else {
								field.Type = "array"
							}
							field.ElementType = &(el)
						} else if ok, _ := isObject(f, fs); ok {
							field.Type = "object"
						} else if a.Fields[f].IsString {
							field.Type = "string"
						}
						objFields[fn] = field
					}
				}
			}

			objects[n] = Object{
				Name:      strcase.ToCamel(n),
				Id:        a.Fields[k+PATH_SEPARATOR+"id"] != nil,
				Named:     a.Fields[k+PATH_SEPARATOR+"name"] != nil,
				Typed:     a.Fields[k+PATH_SEPARATOR+"type"] != nil,
				SubTypes:  getSubtypes(objectTypes, k),
				SubTypeOf: getSubtypeOf(k),
				Fields:    objFields,
			}
		}
	}

	return generateCode(&objects)
}

func generateCode(objects *map[string]Object) error {
	file, _ := json.MarshalIndent(objects, "", "  ")
	_ = ioutil.WriteFile("model.json", file, 0644)

	f, err := os.Create("../backend/model/model.go")
	if err != nil {
		return err
	}
	defer f.Close()

	var buf bytes.Buffer
	err = packageTemplate.Execute(&buf, struct {
		Objects *map[string]Object
	}{
		Objects: objects,
	})
	if err != nil {
		return err
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	_, err = f.Write(p)
	return err
}

func isArray(typ string, types []string) (bool, string) {
	fc := 0
	elements := ""

	if !strings.Contains(typ, PATH_SEPARATOR) || strings.Contains(typ[strings.LastIndex(typ, PATH_SEPARATOR):], "+") {
		return false, ""
	}

	for _, t := range types {
		if !strings.HasPrefix(t, typ+PATH_SEPARATOR) {
			continue
		}
		if strings.Contains(t[len(typ)+1:], PATH_SEPARATOR) {
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
