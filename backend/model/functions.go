package model

import (
	"fmt"
	"html/template"
	"regexp"
	"strconv"
	"strings"

	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

var LinkHf = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).hf(id)) }
var LinkHfShort = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).hfShort(id)) }
var LinkHfList = func(w *DfWorld, id []int) template.HTML { return template.HTML((&Context{World: w}).hfList(id)) }
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
var LinkMountain = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).mountain(id)) }
var LinkLandmass = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).landmass(id)) }
var LinkRiver = func(w *DfWorld, id int) template.HTML { return template.HTML((&Context{World: w}).river(id)) }

var AddMapLandmass = func(w *DfWorld, id int) template.HTML {
	if x, ok := w.Landmasses[id]; ok {
		c1 := strings.Split(x.Coord1, ",")
		x1, _ := strconv.Atoi(c1[0])
		y1, _ := strconv.Atoi(c1[1])
		c2 := strings.Split(x.Coord2, ",")
		x2, _ := strconv.Atoi(c2[0])
		y2, _ := strconv.Atoi(c2[1])
		return template.HTML(fmt.Sprintf(`<script>addLandmass(%d, %d, %d, %d, %d, '#FFF')</script>`, x.Id_, x1, y1, x2, y2))
	}
	return ""
}

var AddMapRegion = func(w *DfWorld, id int) template.HTML {
	if x, ok := w.Regions[id]; ok {
		coords := strings.Join(util.Map(x.Outline(), func(c Coord) string { return fmt.Sprintf(`coord(%d,%d)`, c.X, c.Y-1) }), ",")
		fillColor := "transparent"
		switch x.Evilness {
		case RegionEvilness_Evil:
			fillColor = "fuchsia"
		case RegionEvilness_Good:
			fillColor = "aqua"
		}
		return template.HTML(fmt.Sprintf(`<script>addRegion(%d, [%s], '%s')</script>`, x.Id_, coords, fillColor))
	}
	return ""
}

var AddMapSite = func(w *DfWorld, id int) template.HTML {
	if site, ok := w.Sites[id]; ok {
		coords := strings.Split(site.Rectangle, ":")
		c1 := strings.Split(coords[0], ",")
		x1, _ := strconv.ParseFloat(c1[0], 32)
		y1, _ := strconv.ParseFloat(c1[1], 32)
		c2 := strings.Split(coords[1], ",")
		x2, _ := strconv.ParseFloat(c2[0], 32)
		y2, _ := strconv.ParseFloat(c2[1], 32)
		c := "#ff0"
		if e, ok := w.Entities[site.Owner]; ok {
			c = e.Color()
		}
		if site.Ruin {
			c = "#aaa"
		}
		return template.HTML(fmt.Sprintf(`<script>addSite(%d, %f, %f, %f, %f, "%s", "")</script>`, site.Id_, x1/16.0, y1/16.0-1, x2/16.0, y2/16.0-1, c))
	} else {
		return ""
	}
}

var AddMapMountain = func(w *DfWorld, id int) template.HTML {
	if m, ok := w.MountainPeaks[id]; ok {
		c1 := strings.Split(m.Coords, ",")
		x, _ := strconv.Atoi(c1[0])
		y, _ := strconv.Atoi(c1[1])
		return template.HTML(fmt.Sprintf(`<script>addMountain(%d, %d, %d, '#666')</script>`, m.Id_, x, y))
	}
	return ""
}

var AddMapWorldConstruction = func(w *DfWorld, id int) template.HTML {
	if x, ok := w.WorldConstructions[id]; ok {
		color := util.If(x.Type_ == WorldConstructionType_Tunnel, "#000", "#fff")
		line := x.Line()
		if len(line) == 1 {
			return template.HTML(fmt.Sprintf(`<script>addWc(%d, %d, %d, '%s')</script>`, x.Id_, line[0].X, line[0].Y, color))
		} else {
			r := "<script>"
			r += "var polyline = L.polyline(["
			r += strings.Join(util.Map(x.Line(), func(c Coord) string { return fmt.Sprintf(`coord(%d+0.5,%d-0.5)`, c.X, c.Y) }), ",")
			r += "], {color: '" + color + "', opacity: 1, weight: 3}).addTo(constructionsLayer);\n"
			r += fmt.Sprintf(`attachTooltip(polyline, urlToolTip('worldconstruction', %d));`, x.Id_)
			r += "polyline.on('mouseover', function (e) { this.setStyle({weight: 10}); });\n"
			r += "polyline.on('mouseout', function (e) { this.setStyle({ weight: 3}); });\n"
			r += "</script>"
			return template.HTML(r)
		}
	}
	return ""
}

var AddMapRiver = func(w *DfWorld, id int) template.HTML {
	return ""
}

var AndList = func(s []string) template.HTML { return template.HTML(andList(s)) }

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

func LinkDescription(w *DfWorld, desc string) template.HTML {
	c := &Context{World: w}
	desc = replaceNameDescription(desc, "originating in ", `\.`, w.Entities, c.entity)
	desc = replaceNameDescription(desc, "grew out of the performances of ", `\.`, w.Entities, c.entity)
	desc = replaceNameDescription(desc, "accompanied by(?: any composition of)? ", `(?: as |\.)`, w.MusicalForms, c.musicalForm)
	desc = replaceNameDescription(desc, "(?:recites?|acts? out)(?: any composition of)? ", `(?: while |\.)`, w.PoeticForms, c.poeticForm)
	desc = replacHfDescription(desc, "devised by ", `\.`, w.HistoricalFigures, c.hf)
	desc = replacHfDescription(desc, "the story of ", `\.`, w.HistoricalFigures, c.hf)
	desc = replaceNameDescription(desc, "the words of ", `(?: while |\.)`, w.WrittenContents, c.writtenContent)
	desc = replacHfDescription(desc, "express pleasure with ", " originally", w.HistoricalFigures, c.hf)
	s := strings.Split(desc, "[B]")
	if len(s) > 1 {
		desc = s[0] + "<ul><li>" + strings.Join(s[1:], "</li><li>") + "</li></ul>"
	}
	return template.HTML(desc)
}

type NamedIdentifiable interface {
	Id() int
	Name() string
}

func replaceNameDescription[T NamedIdentifiable](s, prefix, suffix string, input map[int]T, mapper func(int) string) string {
	return replaceDescription(s, prefix, suffix, input, func(t T) string { return t.Name() }, mapper)
}

func replacHfDescription(s, prefix, suffix string, input map[int]*HistoricalFigure, mapper func(int) string) string {
	return replaceDescription(s, prefix, suffix, input,
		func(hf *HistoricalFigure) string {
			if hf.Race != "" && !hf.Deity && !hf.Force {
				return fmt.Sprintf("the %s %s", hf.Race, hf.Name())
			} else {
				return hf.Name()
			}
		}, mapper)
}

func replaceDescription[T NamedIdentifiable](s, prefix, suffix string, input map[int]T, namer func(T) string, mapper func(int) string) string {
	r := "(" + prefix + `)([^.]+?)(` + suffix + ")"
	reg := regexp.MustCompile(r)
	res := reg.FindStringSubmatch(s)
	if res == nil {
		return s
	}

	name := strings.ToLower(res[2])
	for id, v := range input {
		if strings.ToLower(namer(v)) == name {
			return reg.ReplaceAllString(s, res[1]+mapper(id)+res[3])
		}
	}
	return s
}
