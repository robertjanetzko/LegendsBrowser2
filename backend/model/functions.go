package model

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

var LinkHf = func(w *DfWorld, id int) template.HTML { return template.HTML((&context{world: w}).hf(id)) }
var LinkEntity = func(w *DfWorld, id int) template.HTML { return template.HTML((&context{world: w}).entity(id)) }
var LinkSite = func(w *DfWorld, id int) template.HTML { return template.HTML((&context{world: w}).site(id, "")) }
var LinkRegion = func(w *DfWorld, id int) template.HTML { return template.HTML((&context{world: w}).region(id)) }

func andList(list []string) string {
	if len(list) > 1 {
		return strings.Join(list[:len(list)-1], ", ") + " and " + list[len(list)-1]
	}
	return strings.Join(list, ", ")
}

func articled(s string) string {
	if ok, _ := regexp.MatchString("^([aeio]|un|ul).*", s); ok {
		return "an " + s
	}
	return "a " + s
}

func ShortTime(year, seconds int) string {
	if year == -1 {
		return "a time before time"
	}
	return fmt.Sprintf("%d", year)
}

func Time(year, seconds int) string {
	if year == -1 {
		return "a time before time"
	}
	if seconds == -1 {
		return fmt.Sprintf("%d", year)
	}
	return fmt.Sprintf("%s of %d", Season(seconds), year)
}

func Season(seconds int) string {
	r := ""
	month := seconds % 100800
	if month <= 33600 {
		r += "early "
	} else if month <= 67200 {
		r += "mid"
	} else if month <= 100800 {
		r += "late "
	}

	season := seconds % 403200
	if season < 100800 {
		r += "spring"
	} else if season < 201600 {
		r += "summer"
	} else if season < 302400 {
		r += "autumn"
	} else if season < 403200 {
		r += "winter"
	}

	return r
}

func equipmentLevel(level int) string {
	switch level {
	case 1:
		return "well-crafted"
	case 2:
		return "finely-crafted"
	case 3:
		return "superior quality"
	case 4:
		return "exceptional"
	case 5:
		return "masterwork"
	}
	return ""
}

func containsInt(list []int, id int) bool {
	for _, v := range list {
		if v == id {
			return true
		}
	}
	return false
}
