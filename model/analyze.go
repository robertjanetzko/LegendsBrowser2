package model

import (
	"fmt"
	"legendsbrowser/util"
	"sort"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

func ListOtherElements[T TypedOthers](items *[]T) {
	m := make(map[string]map[string]bool)
	cantInt := make(map[string]bool)
	isObj := make(map[string]bool)
	isMultiple := make(map[string]bool)
	for _, item := range *items {
		found := make(map[string]bool)
		for _, el := range item.Others() {
			if !m[el.XMLName.Local][item.Type()] {
				if m[el.XMLName.Local] == nil {
					m[el.XMLName.Local] = map[string]bool{}
				}
				m[el.XMLName.Local][item.Type()] = true
			}
			_, err := strconv.Atoi(el.Value)
			if err != nil {
				cantInt[el.XMLName.Local] = true
			}
			if strings.Contains(el.Value, "<") {
				isObj[el.XMLName.Local] = true
			}
			if found[el.XMLName.Local] {
				isMultiple[el.XMLName.Local] = true
			}
			found[el.XMLName.Local] = true
		}
	}
	ks := util.Keys(m)
	sort.Strings(ks)
	for _, k := range ks {
		events := util.Keys(m[k])
		sort.Strings(events)
		// fmt.Println(strconv.FormatBool(cantInt[k]) + " - " + k + ": " + strings.Join(events, ", "))
		var mult string
		if isMultiple[k] {
			mult = "[]"
		} else {
			mult = ""
		}

		if isObj[k] {
			fmt.Printf("// %s object\n", k)
		} else if cantInt[k] {
			fmt.Printf("%s *%sstring `xml:\"%s\" json:\"%s,omitempty\"`\n", strcase.ToCamel(k), mult, k, strcase.ToLowerCamel(k))
		} else {
			var types []string
			if util.ContainsAny(k, "entity_id", "enid", "civ_id", "entity_1", "entity_2") {
				types = append(types, "entity")
			}
			if util.ContainsAny(k, "site_id") {
				types = append(types, "site")
			}
			if util.ContainsAny(k, "structure_id") {
				types = append(types, "structure")
			}
			if util.ContainsAny(k, "hfid", "hist_figure_id", "hist_fig_id") {
				types = append(types, "hf")
			}
			if util.ContainsAny(k, "wcid", "wc_id") {
				types = append(types, "wc")
			}
			if util.ContainsAny(k, "artifact_id") {
				types = append(types, "artifact")
			}
			typestr := strings.Join(types, ",")
			if typestr != "" {
				typestr = fmt.Sprintf(" legend:\"%s\"", typestr)
			}
			fmt.Printf("%s *%sint `xml:\"%s\" json:\"%s,omitempty\"%s`\n", strcase.ToCamel(k), mult, k, strcase.ToLowerCamel(k), typestr)
		}
	}
}
