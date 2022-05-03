package df

import (
	"fmt"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

var allowedTyped = map[string]bool{
	"df_world|historical_events|historical_event":                       true,
	"df_world|historical_event_collections|historical_event_collection": true,
}

func typeName(k string) string {
	if strings.Contains(k, PATH_SEPARATOR) {
		return k[strings.LastIndex(k, PATH_SEPARATOR)+1:]
	}
	return k
}

type Metadata map[string]Object

func CreateMetadata(a *AnalyzeData) (*Metadata, error) {
	fs := filterSubtypes(&a.Fields)

	// unique type names
	names := make(map[string]bool)
	for k := range a.Fields {
		path := strings.Split(k, PATH_SEPARATOR)
		if len(path) >= 2 {
			names[strings.Join(path[:len(path)-1], PATH_SEPARATOR)] = true
		}
	}
	objectTypes := util.Keys(names)

	objects := make(Metadata, 0)
	names = make(map[string]bool)
	double := make(map[string]bool)
	typeNames := make(map[string]string)

	// check object type names
	for _, k := range objectTypes {
		n := typeName(k)
		if _, ok := names[n]; ok {
			double[n] = true
		}
		names[n] = true
	}

	for _, k := range objectTypes {
		typeNames[k] = strcase.ToCamel(typeName(k))
	}
	for _, n := range util.Keys(double) {
		for _, k := range objectTypes {
			if strings.HasSuffix(k, PATH_SEPARATOR+n) {
				path := strings.Split(k, PATH_SEPARATOR)
				for i := len(path) - 1; i > 0; i-- {
					sub := strings.Join(path[:i], PATH_SEPARATOR)
					if ok, _ := isArray(sub, fs); !ok {
						typeNames[k] = strcase.ToCamel(typeName(sub) + "_" + typeName(k))
						break
					}
				}

			}
		}
	}

	// build metadata
	for _, k := range objectTypes {
		if ok, _ := isArray(k, fs); !ok {
			n := typeName(k)

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
							Base:     a.Fields[f].Base,
							Plus:     a.Fields[f].Plus,
							Related:  a.Overwrites.Relations[fmt.Sprintf("%s.%s", typeNames[k], strcase.ToCamel(fn))],
						}
						if ok, elements := isArray(f, fs); ok {
							el := typeNames[elements]
							if _, ok := a.Fields[elements+PATH_SEPARATOR+"id"]; ok {
								field.Type = "map"
							} else {
								field.Type = "array"
							}
							field.ElementType = &(el)
						} else if ok, _ := isObject(f, fs); ok {
							field.Type = "object"
							el := typeNames[f]
							field.ElementType = &el
						} else if !a.Fields[f].NoBool {
							field.Type = "bool"
						} else if a.Fields[f].IsString {
							if !a.Fields[f].Enum {
								field.Type = "string"
							} else {
								var vs []string

								for _, f2 := range fs {
									if strings.HasPrefix(f2, k+PATH_SEPARATOR) {
										fn2 := f2[len(k)+1:]
										if !strings.Contains(fn2, PATH_SEPARATOR) {
											if sn, ok := sameFields[typeNames[k]][strcase.ToCamel(fn2)]; ok {
												if sn == strcase.ToCamel(fn) {
													for k := range a.Fields[f2].Values {
														a.Fields[f].Values[k] = true
													}
												}
											}
										}
									}
								}

								for k := range a.Fields[f].Values {
									vs = append(vs, k)
								}
								sort.Strings(vs)
								field.EnumValues = &vs
								field.Type = "enum"
								el := typeNames[k] + field.Name
								field.ElementType = &el
							}
						}
						objFields[fn] = field
					}
				}
			}

			additional := make(map[string]Field)
			if afs, ok := a.Overwrites.AdditionalFields[typeNames[k]]; ok {
				for _, add := range afs {
					additional[add.Name] = Field{
						Name:   add.Name,
						Type:   add.Type,
						Legend: "add",
					}
				}
			}

			objects[typeNames[k]] = Object{
				Name:       typeNames[k],
				Id:         a.Fields[k+PATH_SEPARATOR+"id"] != nil,
				Named:      a.Fields[k+PATH_SEPARATOR+"name"] != nil,
				Typed:      a.Fields[k+PATH_SEPARATOR+"type"] != nil,
				SubTypes:   getSubtypes(a, k),
				SubTypeOf:  getSubtypeOf(k),
				SubType:    getSubtype(k),
				Fields:     objFields,
				Additional: additional,
			}
		}
	}

	return &objects, nil
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

func getSubtypes(a *AnalyzeData, k string) *[]Subtype {
	if allowedTyped[k] {
		if st, ok := a.SubTypes[k]; ok {
			var list []Subtype
			for _, v := range *st {
				list = append(list, *v)
			}
			return &list
		}
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

func getSubtype(k string) *string {
	if strings.Contains(k, PATH_SEPARATOR) {
		last := k[strings.LastIndex(k, PATH_SEPARATOR)+1:]
		if strings.Contains(last, "+") {
			base := strcase.ToDelimited(last[strings.Index(last, "+")+1:], ' ')
			return &base
		}
	}
	return nil
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

func ListEnumCandidates(a *AnalyzeData) {
	keys := util.Keys(a.Fields)
	sort.Strings(keys)

	for _, k := range keys {
		f := a.Fields[k]
		if !f.Enum {
			continue
		}
		f.Enum = false
		n := k[strings.LastIndex(k, PATH_SEPARATOR)+1:]
		if n == "name" || n == "altname" || strings.Contains(n, "name_") || strings.Contains(n, "spouse") || n == "coords" || n == "rectangle" || n == "interaction_action" || strings.Contains(n, "race") || strings.Contains(n, "caste") || strings.Contains(n, "_mat") {
			continue
		}
		if f.IsString {
			v := util.Keys(f.Values)
			sort.Strings(v)
			if len(v) == 0 {
				continue
			}
			fmt.Println(k, ":", strings.Join(v, ", "))
			f.Enum = true
		}
	}
}
