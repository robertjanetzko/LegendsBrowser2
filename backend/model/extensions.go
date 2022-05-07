package model

import (
	"fmt"
	"html/template"
	"sort"
	"strings"

	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func (w *DfWorld) AllEventTypes() []string {
	types := make(map[string]bool)
	for _, e := range w.HistoricalEvents {
		types[e.Details.Type()] = true
	}
	var list = util.Keys(types)
	sort.Strings(list)
	return list
}

func (w *DfWorld) EventsOfType(t string) any {
	var list []*HistoricalEvent
	for _, e := range w.HistoricalEvents {
		if e.Details.Type() == t {
			list = append(list, e)
		}
	}

	sort.Slice(list, func(i, j int) bool { return list[i].Id_ < list[j].Id_ })

	return struct {
		Type   string
		Events []*HistoricalEvent
	}{
		Type:   t,
		Events: list,
	}
}

func (w *DfWorld) EventsMatching(f func(HistoricalEventDetails) bool) []*HistoricalEvent {
	var list []*HistoricalEvent
	for _, e := range w.HistoricalEvents {
		if f(e.Details) {
			list = append(list, e)
		}
	}
	sort.Slice(list, func(a, b int) bool { return list[a].Id_ < list[b].Id_ })
	return list
}

func (w *DfWorld) SiteHistory(siteId int) []*HistoricalEvent {
	var list []*HistoricalEvent
	for _, e := range w.HistoricalEvents {
		if e.Details.RelatedToSite(siteId) {
			switch e.Details.(type) {
			case *HistoricalEventCreatedSite, *HistoricalEventDestroyedSite, *HistoricalEventSiteTakenOver, *HistoricalEventHfDestroyedSite, *HistoricalEventReclaimSite:
				list = append(list, e)
			}
		}
	}
	sort.Slice(list, func(a, b int) bool { return list[a].Id_ < list[b].Id_ })
	return list
}

func (c *HistoricalEventCollection) Type() string {
	if c.Details == nil {
		return "unk"
	}
	return c.Details.Type()
}

func (e *HistoricalEventCollection) Html(c *Context) string {
	if e.Details == nil {
		return "unk"
	}
	return e.Details.Html(e, c)
}

func (e *Artifact) Type() string {
	switch e.ItemSubtype {
	case "scroll":
		return "scroll"
	}
	switch e.ItemType {
	case "weapon", "tool", "book", "slab":
		return e.ItemType
	case "armor", "shoe", "gloves", "helm", "pants":
		return "armor"
	default:
		return "item"
	}
}

func (e *Entity) Type() string {
	return e.Type_.String()
}

func (x *Entity) Color() string {
	c := ""
	switch x.Race {
	case "dwarf":
		c = `#FFCC33`
	case "elf":
		c = `#99FF00`
	case "human":
		c = `#0000CC`
	case "kobold":
		c = `#333`
	case "goblin":
		c = `#CC0000`
	}
	if x.Necromancer {
		c = `#A0A`
	}
	return c
}

func (e *Entity) Weapons() []string {
	return util.Map(e.Weapon, func(w EntityWeapon) string { return w.String() })
}

func (e *Entity) Position(id int) *EntityPosition {
	for _, p := range e.EntityPosition {
		if p.Id_ == id {
			return p
		}
	}
	return &EntityPosition{Name_: "UNKNOWN POSITION"}
}

func (e *Entity) PositionByIndex(index int) *EntityPosition {
	return e.EntityPosition[len(e.EntityPosition)-1-index]
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

func (hf *HistoricalFigure) Goals() string {
	return andList(util.Map(hf.Goal, func(g HistoricalFigureGoal) string { return g.String() }))
}

func (hf *HistoricalFigure) Pets() string {
	return andList(util.Map(hf.JourneyPet, func(g string) string { return articled(strings.ToLower(g)) }))
}

func (x *Honor) Requirement() string {
	var list []string
	if x.RequiresAnyMeleeOrRangedSkill {
		list = append(list, "attaining sufficent skill with a weapon or technique")
	}
	if x.RequiredSkill != HonorRequiredSkill_Unknown {
		list = append(list, "attaining enough skill with the "+x.RequiredSkill.String())
	}
	if x.RequiredBattles == 1 {
		list = append(list, "serving in combat")
	}
	if x.RequiredBattles > 1 {
		list = append(list, fmt.Sprintf("participating in %d battles", x.RequiredBattles))
	}
	if x.RequiredYears >= 1 {
		list = append(list, fmt.Sprintf("%d years of membership", x.RequiredYears))
	}
	if x.RequiredKills >= 1 {
		list = append(list, fmt.Sprintf("slaying %d enemies", x.RequiredKills))
	}

	return " after " + andList(list)
}

func (r *Region) Type() string {
	return r.Type_.String()
}

func (s *Site) Type() string {
	return s.Type_.String()
}

func (s *Structure) Type() string {
	return s.Type_.String()
}

func (w *WorldConstruction) Type() string {
	return w.Type_.String()
}

func (w *WrittenContent) Name() string {
	return w.Title
}

func (w *WrittenContent) Type() string {
	return w.Form.String()
}

func (w *DanceForm) Type() string {
	return "dance form"
}

func (w *MusicalForm) Type() string {
	return "musical form"
}

func (w *PoeticForm) Type() string {
	return "poetic form"
}

func (r *Reference) Html(c *Context) template.HTML {
	switch r.Type_ {
	case ReferenceType_ABSTRACTBUILDING:
		return template.HTML("a building")
	case ReferenceType_ARTIFACT:
		return template.HTML(c.artifact(r.Id_))
	case ReferenceType_DANCEFORM:
		return template.HTML(c.danceForm(r.Id_))
	case ReferenceType_ENTITY:
		return template.HTML(c.entity(r.Id_))
	case ReferenceType_HISTORICALEVENT:
		if e, ok := c.World.HistoricalEvents[r.Id_]; ok {
			return template.HTML("how in " + Time(e.Year, e.Seconds72) + " " + e.Details.Html(c))
		}
	case ReferenceType_HISTORICALFIGURE:
		return template.HTML(c.hf(r.Id_))
	case ReferenceType_INTERACTION:
		return template.HTML("an interaction")
	case ReferenceType_KNOWLEDGESCHOLARFLAG:
		return template.HTML("specific knowledge")
	case ReferenceType_LANGUAGE:
		return template.HTML("a language")
	case ReferenceType_MUSICALFORM:
		return template.HTML(c.musicalForm(r.Id_))
	case ReferenceType_POETICFORM:
		return template.HTML(c.poeticForm(r.Id_))
	case ReferenceType_SITE:
		return template.HTML(c.site(r.Id_, ""))
	case ReferenceType_SUBREGION:
		return template.HTML(c.region(r.Id_))
	case ReferenceType_VALUELEVEL:
		return template.HTML("a value")
	case ReferenceType_WRITTENCONTENT:
		return template.HTML(c.writtenContent(r.Id_))
	}
	return template.HTML(r.Type_.String())
}
