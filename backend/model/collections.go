package model

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func (x *HistoricalEventCollection) Link(s string) string {
	return fmt.Sprintf(`<a class="collection %s" href="/collection/%d">%s</a>`, strcase.ToKebab(x.Details.Type()), x.Id_, util.Title(s))
}

func (x *HistoricalEventCollection) ParentId() int {
	switch t := x.Details.(type) {
	case *HistoricalEventCollectionAbduction:
		return t.ParentEventcol
	case *HistoricalEventCollectionBattle:
		return t.WarEventcol
	case *HistoricalEventCollectionBeastAttack:
		return t.ParentEventcol
	case *HistoricalEventCollectionDuel:
		return t.ParentEventcol
	case *HistoricalEventCollectionRaid:
		return t.ParentEventcol
	case *HistoricalEventCollectionSiteConquered:
		return t.WarEventcol
	case *HistoricalEventCollectionTheft:
		return t.ParentEventcol
	}
	return -1
}

func ord(ordinal int) string {
	switch ordinal {
	case 1:
		return ""
	case 2:
		return "second" + " "
	case 3:
		return "third" + " "
	}
	return humanize.Ordinal(ordinal) + " "
}

func (x *HistoricalEventCollectionAbduction) Html(e *HistoricalEventCollection, c *Context) string {
	loc := c.location(x.SiteId, " at", x.SubregionId, " at")
	switch l := len(x.TargetHfids); {
	case l == 0:
		return "the " + e.Link(ord(x.Ordinal)+"attempted abduction") + loc
	case l == 1:
		return "the " + e.Link(ord(x.Ordinal)+"abduction") + " of " + c.hf(x.TargetHfids[0]) + loc
	}
	return "the " + e.Link(ord(x.Ordinal)+"abduction") + loc
}

func (x *HistoricalEventCollectionBattle) Html(e *HistoricalEventCollection, c *Context) string {
	return e.Link(util.Title(x.Name_))
}

func (x *HistoricalEventCollectionBeastAttack) Html(e *HistoricalEventCollection, c *Context) string {
	r := "the "
	switch l := len(x.AttackerHfIds); {
	case l == 1:
		r += e.Link(ord(x.Ordinal)+"rampage") + " of " + c.hf(x.AttackerHfIds[0])
	case l > 1:
		if hf, ok := c.World.HistoricalFigures[x.AttackerHfIds[0]]; ok {
			r += e.Link(ord(x.Ordinal) + hf.Race + " " + "rampage")
		}
	default:
		r += e.Link(ord(x.Ordinal) + "rampage")
	}
	r += c.location(x.SiteId, " in", x.SubregionId, " in")
	return r
}

func (x *HistoricalEventCollectionCeremony) Html(e *HistoricalEventCollection, c *Context) string {
	r := "ceremony"
	if len(e.Event) > 0 {
		if event, ok := c.World.HistoricalEvents[e.Event[0]]; ok {
			if d, ok := event.Details.(*HistoricalEventCeremony); ok {
				if entity, ok := c.World.Entities[d.CivId]; ok {
					if d.OccasionId < len(entity.Occasion) {
						occ := entity.Occasion[d.OccasionId]
						if len(occ.Schedule) > 1 {
							switch d.ScheduleId {
							case 0:
								r = "opening ceremony"
							case len(occ.Schedule) - 1:
								r = "closing ceremony"
							default:
								r = "main ceremony"
							}
						}
					}
				}
			}
		}
	}
	return "the " + e.Link(ord(x.Ordinal)+r) + " of " + c.collection(x.OccasionEventcol)
}

func (x *HistoricalEventCollectionCompetition) Html(e *HistoricalEventCollection, c *Context) string {
	r := "competition"
	if len(e.Event) > 0 {
		if event, ok := c.World.HistoricalEvents[e.Event[0]]; ok {
			if d, ok := event.Details.(*HistoricalEventCompetition); ok {
				if entity, ok := c.World.Entities[d.CivId]; ok {
					if d.OccasionId < len(entity.Occasion) {
						occ := entity.Occasion[d.OccasionId]
						if d.ScheduleId < len(occ.Schedule) {
							r = occ.Schedule[d.ScheduleId].Type_.String()
						}
					}
				}
			}
		}
	}
	return "the " + e.Link(ord(x.Ordinal)+r) + " of " + c.collection(x.OccasionEventcol)
}

func (x *HistoricalEventCollectionDuel) Html(e *HistoricalEventCollection, c *Context) string {
	r := "the "
	r += e.Link(ord(x.Ordinal)+"duel") + " of " + c.hf(x.AttackingHfid) + " and " + c.hfRelated(x.DefendingHfid, x.AttackingHfid)
	r += c.location(x.SiteId, " in", x.SubregionId, " in")
	return r
}

func (x *HistoricalEventCollectionEntityOverthrown) Html(e *HistoricalEventCollection, c *Context) string {
	return "the " + e.Link(ord(x.Ordinal)+"overthrow") + " of " + c.entity(x.TargetEntityId) + c.site(x.SiteId, " in")
}

func (x *HistoricalEventCollectionInsurrection) Html(e *HistoricalEventCollection, c *Context) string {
	return "the " + e.Link(ord(x.Ordinal)+"insurrection") + c.site(x.SiteId, " at")
}

func (x *HistoricalEventCollectionJourney) Html(e *HistoricalEventCollection, c *Context) string {
	r := "the "
	r += e.Link(ord(x.Ordinal)+"journey") + " of " + c.hfList(x.TravellerHfIds)
	return r
}

func (x *HistoricalEventCollectionOccasion) Html(e *HistoricalEventCollection, c *Context) string {
	if civ, ok := c.World.Entities[x.CivId]; ok {
		if x.OccasionId < len(civ.Occasion) {
			occ := civ.Occasion[x.OccasionId]
			return util.If(x.Ordinal > 1, "the "+ord(x.Ordinal)+"occasion of ", "") + e.Link(occ.Name_)
		}
	}
	return util.If(x.Ordinal > 1, "the "+ord(x.Ordinal)+"occasion of ", "") + e.Link("UNKNOWN OCCASION")
}

func (x *HistoricalEventCollectionPerformance) Html(e *HistoricalEventCollection, c *Context) string {
	r := "performance"
	if len(e.Event) > 0 {
		if event, ok := c.World.HistoricalEvents[e.Event[0]]; ok {
			if d, ok := event.Details.(*HistoricalEventPerformance); ok {
				if entity, ok := c.World.Entities[d.CivId]; ok {
					if d.OccasionId < len(entity.Occasion) {
						occ := entity.Occasion[d.OccasionId]
						if d.ScheduleId < len(occ.Schedule) {
							r = occ.Schedule[d.ScheduleId].Type_.String()
						}
					}
				}
			}
		}
	}
	return "the " + e.Link(ord(x.Ordinal)+r) + " of " + c.collection(x.OccasionEventcol)
}

func (x *HistoricalEventCollectionPersecution) Html(e *HistoricalEventCollection, c *Context) string {
	return "the " + e.Link(ord(x.Ordinal)+"persecution") + " of " + c.entity(x.TargetEntityId) + c.site(x.SiteId, " in")
}

func (x *HistoricalEventCollectionProcession) Html(e *HistoricalEventCollection, c *Context) string {
	return "the " + e.Link(ord(x.Ordinal)+"procession") + " of " + c.collection(x.OccasionEventcol)
}

func (x *HistoricalEventCollectionPurge) Html(e *HistoricalEventCollection, c *Context) string {
	return "the " + e.Link(ord(x.Ordinal)+x.Adjective.String()+" purge") + c.site(x.SiteId, " in")
}

func (x *HistoricalEventCollectionRaid) Html(e *HistoricalEventCollection, c *Context) string {
	return "the " + e.Link(ord(x.Ordinal)+"raid") + c.site(x.SiteId, " at")
}

func (x *HistoricalEventCollectionSiteConquered) Html(e *HistoricalEventCollection, c *Context) string {
	return "the " + e.Link(ord(x.Ordinal)+"pillaging") + c.site(x.SiteId, " of")
}

func (x *HistoricalEventCollectionTheft) Html(e *HistoricalEventCollection, c *Context) string {
	return "the " + e.Link(ord(x.Ordinal)+"theft") + c.site(x.SiteId, " at")
}

func (x *HistoricalEventCollectionWar) Html(e *HistoricalEventCollection, c *Context) string {
	return e.Link(util.Title(x.Name_))
}
