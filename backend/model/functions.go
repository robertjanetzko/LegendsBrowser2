package model

import (
	"fmt"
	"html/template"
	"regexp"
	"strconv"
	"strings"
)

var LinkHf = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).hf(id)) }
var LinkEntity = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).entity(id)) }
var LinkSite = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).site(id, "")) }
var LinkStructure = func(w *DfWorld, siteId, id int) template.HTML {
	return template.HTML((&Context{World: w}).structure(siteId, id))
}
var LinkRegion = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).region(id)) }
var LinkWorldConstruction = func(w *DfWorld, id int) template.HTML {
	return template.HTML((&Context{World: w}).worldConstruction(id))
}
var LinkArtifact = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).artifact(id)) }
var LinkDanceForm = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).danceForm(id)) }
var LinkMusicalForm = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).musicalForm(id)) }
var LinkPoeticForm = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).poeticForm(id)) }
var LinkWrittenContent = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).writtenContent(id)) }
var LinkCollection = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).collection(id)) }

var AddMapSite = func(w *DfWorld, id int) template.HTML {
	if site, ok := w.Sites[id]; ok {
		coords := strings.Split(site.Rectangle, ":")
		c1 := strings.Split(coords[0], ",")
		x1, _ := strconv.ParseFloat(c1[0], 32)
		y1, _ := strconv.ParseFloat(c1[1], 32)
		c2 := strings.Split(coords[1], ",")
		x2, _ := strconv.ParseFloat(c2[0], 32)
		y2, _ := strconv.ParseFloat(c2[1], 32)
		return template.HTML(fmt.Sprintf(`<script>addSite("%s", %f, %f, %f, %f, "#FF0", "")</script>`, site.Name(), x1/16.0, y1/16.0-1, x2/16.0, y2/16.0-1))
	} else {
		return ""
	}
}

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
