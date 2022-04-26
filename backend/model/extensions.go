package model

import (
	"fmt"
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

func (w *WrittenContent) Name() string {
	return w.Title
}
