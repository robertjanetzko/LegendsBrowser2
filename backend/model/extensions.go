package model

import (
	"fmt"
	"html/template"
	"regexp"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func (e *Entity) Position(id int) *EntityPosition {
	for _, p := range e.EntityPosition {
		if p.Id_ == id {
			return p
		}
	}
	return &EntityPosition{Name_: "UNKNOWN POSITION"}
}

func (p *EntityPosition) GenderName(hf *HistoricalFigure) string {
	if hf.Female() && p.NameFemale != "" {
		return p.NameFemale
	} else if hf.Male() && p.NameMale != "" {
		return p.NameMale
	} else {
		return p.Name_
	}
}

func (hf *HistoricalFigure) Female() bool {
	return hf.Sex == 0 || hf.Caste == "FEMALE"
}

func (hf *HistoricalFigure) Male() bool {
	return hf.Sex == 1 || hf.Caste == "MALE"
}

func (hf *HistoricalFigure) Pronoun() string {
	if hf.Female() {
		return "she"
	}
	return "he"
}

func (hf *HistoricalFigure) PossesivePronoun() string {
	if hf.Female() {
		return "her"
	}
	return "his"
}

func (hf *HistoricalFigure) FirstName() string {
	return strings.Split(hf.Name_, " ")[0]
}

type HistoricalEventDetails interface {
	RelatedToEntity(int) bool
	RelatedToHf(int) bool
	RelatedToArtifact(int) bool
	RelatedToSite(int) bool
	RelatedToRegion(int) bool
	Html(*context) string
	Type() string
}

type HistoricalEventCollectionDetails interface {
}

type EventList struct {
	Events  []*HistoricalEvent
	Context *context
}

func filter(f func(HistoricalEventDetails) bool) []*HistoricalEvent {
	var list []*HistoricalEvent
	for _, e := range world.HistoricalEvents {
		if f(e.Details) {
			list = append(list, e)
		}
	}
	sort.Slice(list, func(a, b int) bool { return list[a].Id_ < list[b].Id_ })
	return list
}

func NewEventList(obj any) *EventList {
	el := EventList{
		Context: &context{HfId: -1},
	}

	switch x := obj.(type) {
	case *Entity:
		el.Events = filter(func(d HistoricalEventDetails) bool { return d.RelatedToEntity(x.Id()) })
	case *HistoricalFigure:
		el.Context.HfId = x.Id()
		el.Events = filter(func(d HistoricalEventDetails) bool { return d.RelatedToHf(x.Id()) })
	case *Artifact:
		el.Events = filter(func(d HistoricalEventDetails) bool { return d.RelatedToArtifact(x.Id()) })
	case *Site:
		el.Events = filter(func(d HistoricalEventDetails) bool { return d.RelatedToSite(x.Id()) })
	case *Region:
		el.Events = filter(func(d HistoricalEventDetails) bool { return d.RelatedToRegion(x.Id()) })
	case []*HistoricalEvent:
		el.Events = x
	default:
		fmt.Printf("unknown type %T\n", obj)
	}

	return &el
}

type context struct {
	HfId int
}

func (c *context) hf(id int) string {
	if c.HfId != -1 {
		if c.HfId == id {
			return hfShort(id)
		} else {
			return hfRelated(id, c.HfId)
		}
	}
	return hf(id)
}

func (c *context) hfShort(id int) string {
	return hfShort(id)
}

func (c *context) hfRelated(id, to int) string {
	if c.HfId != -1 {
		if c.HfId == id {
			return hfShort(id)
		} else {
			return hfRelated(id, c.HfId)
		}
	}
	return hfRelated(id, to)
}

func (c *context) hfList(ids []int) string {
	return andList(util.Map(ids, func(id int) string { return c.hf(id) }))
}

func (c *context) hfListRelated(ids []int, to int) string {
	return andList(util.Map(ids, func(id int) string { return c.hfRelated(id, to) }))
}

func containsInt(list []int, id int) bool {
	for _, v := range list {
		if v == id {
			return true
		}
	}
	return false
}

var world *DfWorld

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

func artifact(id int) string {
	if x, ok := world.Artifacts[id]; ok {
		return fmt.Sprintf(`<a class="artifact" href="/artifact/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN ARTIFACT"
}

func entity(id int) string {
	if x, ok := world.Entities[id]; ok {
		return fmt.Sprintf(`<a class="entity" href="/entity/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN ENTITY"
}

func entityList(ids []int) string {
	return andList(util.Map(ids, entity))
}

func position(entityId, positionId, hfId int) string {
	if e, ok := world.Entities[entityId]; ok {
		if h, ok := world.HistoricalFigures[hfId]; ok {
			return e.Position(positionId).GenderName(h)
		}
	}
	return "UNKNOWN POSITION"
}

func siteCiv(siteCivId, civId int) string {
	if siteCivId == civId {
		return entity(civId)
	}
	return util.If(siteCivId != -1, entity(siteCivId), "") + util.If(civId != -1 && siteCivId != -1, " of ", "") + util.If(civId != -1, entity(civId), "")
}

func siteStructure(siteId, structureId int, prefix string) string {
	if siteId == -1 {
		return ""
	}
	return " " + prefix + " " + util.If(structureId != -1, structure(siteId, structureId)+" in ", "") + site(siteId, "")
}

func hf(id int) string {
	if x, ok := world.HistoricalFigures[id]; ok {
		return fmt.Sprintf(`the %s <a class="hf" href="/hf/%d">%s</a>`, x.Race+util.If(x.Deity, " deity", "")+util.If(x.Force, " force", ""), x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func hfShort(id int) string {
	if x, ok := world.HistoricalFigures[id]; ok {
		return fmt.Sprintf(`<a class="hf" href="/hf/%d">%s</a>`, x.Id(), util.Title(x.FirstName()))
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func hfRelated(id, to int) string {
	if x, ok := world.HistoricalFigures[id]; ok {
		if t, ok := world.HistoricalFigures[to]; ok {
			if y, ok := util.Find(t.HfLink, func(l *HfLink) bool { return l.Hfid == id }); ok {
				return fmt.Sprintf(`%s %s <a class="hf" href="/hf/%d">%s</a>`, t.PossesivePronoun(), y.LinkType, x.Id(), util.Title(x.Name()))
			}
		}
		return fmt.Sprintf(`the %s <a class="hf" href="/hf/%d">%s</a>`, x.Race+util.If(x.Deity, " deity", "")+util.If(x.Force, " force", ""), x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func pronoun(id int) string {
	if x, ok := world.HistoricalFigures[id]; ok {
		return x.Pronoun()
	}
	return "he"
}

func posessivePronoun(id int) string {
	if x, ok := world.HistoricalFigures[id]; ok {
		return x.PossesivePronoun()
	}
	return "his"
}

func site(id int, prefix string) string {
	if x, ok := world.Sites[id]; ok {
		return fmt.Sprintf(`%s <a class="site" href="/site/%d">%s</a>`, prefix, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN SITE"
}

func structure(siteId, structureId int) string {
	if x, ok := world.Sites[siteId]; ok {
		if y, ok := x.Structures[structureId]; ok {
			return fmt.Sprintf(`<a class="structure" href="/site/%d/structure/%d">%s</a>`, siteId, structureId, util.Title(y.Name()))
		}
	}
	return "UNKNOWN STRUCTURE"
}

func property(siteId, propertyId int) string {
	if x, ok := world.Sites[siteId]; ok {
		if y, ok := x.SiteProperties[propertyId]; ok {
			if y.StructureId != -1 {
				return structure(siteId, y.StructureId)
			}
			return articled(y.Type.String())
		}
	}
	return "UNKNOWN PROPERTY"
}

func region(id int) string {
	if x, ok := world.Regions[id]; ok {
		return fmt.Sprintf(`<a class="region" href="/region/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN REGION"
}

func location(siteId int, sitePrefix string, regionId int, regionPrefix string) string {
	if siteId != -1 {
		return site(siteId, sitePrefix)
	}
	if regionId != -1 {
		return regionPrefix + " " + region(regionId)
	}
	return ""
}

func place(structureId, siteId int, sitePrefix string, regionId int, regionPrefix string) string {
	if siteId != -1 {
		return siteStructure(siteId, structureId, sitePrefix)
	}
	if regionId != -1 {
		return regionPrefix + " " + region(regionId)
	}
	return ""
}

func identity(id int) string {
	if x, ok := world.Identities[id]; ok {
		return fmt.Sprintf(`<a class="identity" href="/region/%d">%s</a>`, x.Id(), util.Title(x.Name()))
	}
	return "UNKNOWN IDENTITY"
}

func fullIdentity(id int) string {
	if x, ok := world.Identities[id]; ok {
		return fmt.Sprintf(`&quot;the %s <a class="identity" href="/region/%d">%s</a> of %s&quot;`, x.Profession.String(), x.Id(), util.Title(x.Name()), entity(x.EntityId))
	}
	return "UNKNOWN IDENTITY"
}

func feature(x *Feature) string {
	switch x.Type {
	case FeatureType_DancePerformance:
		return "a perfomance of " + danceForm(x.Reference)
	case FeatureType_Images:
		if x.Reference != -1 {
			return "images of " + hf(x.Reference)
		}
		return "images"
	case FeatureType_MusicalPerformance:
		return "a perfomance of " + musicalForm(x.Reference)
	case FeatureType_PoetryRecital:
		return "a recital of " + poeticForm(x.Reference)
	case FeatureType_Storytelling:
		if x.Reference != -1 {
			return "a telling of the story of " + hf(x.Reference)
		}
		return "a story recital"
	default:
		return strcase.ToDelimited(x.Type.String(), ' ')
	}
}

func danceForm(id int) string {
	if x, ok := world.DanceForms[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/danceForm/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN DANCE FORM"
}

func musicalForm(id int) string {
	if x, ok := world.MusicalForms[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/musicalForm/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN MUSICAL FORM"
}

func poeticForm(id int) string {
	if x, ok := world.PoeticForms[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/poeticForm/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN POETIC FORM"
}

func worldConstruction(id int) string {
	if x, ok := world.WorldConstructions[id]; ok {
		return fmt.Sprintf(`<a class="artform" href="/wc/%d">%s</a>`, id, util.Title(x.Name()))
	}
	return "UNKNOWN WORLD CONSTRUCTION"
}

var LinkHf = func(id int) template.HTML { return template.HTML(hf(id)) }
var LinkEntity = func(id int) template.HTML { return template.HTML(entity(id)) }
var LinkSite = func(id int) template.HTML { return template.HTML(site(id, "")) }
var LinkRegion = func(id int) template.HTML { return template.HTML(region(id)) }

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
