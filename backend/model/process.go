package model

import (
	"sort"
	"strings"

	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func (w *DfWorld) process() {
	w.addRelationshipEvents()

	// set site in structure
	for _, site := range w.Sites {
		for _, structure := range site.Structures {
			structure.SiteId = site.Id_
		}
	}

	w.processEvents()
	w.processCollections()
	w.processHistoricalFigures()

	for _, e := range w.Entities {
		if len(e.Sites) > 0 {
			if site, ok := w.Sites[e.Sites[0]]; ok {
				if site.Type_ == SiteType_Tower {
					e.Necromancer = true
				}
			}
		}

		for _, l := range e.EntityLink {
			if l.Type_ == EntityEntityLinkType_PARENT {
				e.Parent = l.Target
			}
		}

		idx := slices.Index(e.Child, e.Id_)
		if idx != -1 {
			e.Child = append(e.Child[:idx], e.Child[idx+1:]...)
		}
		sort.Slice(e.Wars, func(i, j int) bool { return e.Wars[i].Id_ < e.Wars[j].Id_ })
	}

	// check events texts
	// for _, e := range w.HistoricalEvents {
	// 	e.Details.Html(&Context{World: w})
	// }

}

func (w *DfWorld) processEvents() {
	list := maps.Values(w.HistoricalEvents)
	sort.Slice(list, func(i, j int) bool { return list[i].Id_ < list[j].Id_ })

	for _, e := range list {
		switch d := e.Details.(type) {
		case *HistoricalEventHfDoesInteraction:
			if strings.HasPrefix(d.Interaction, "DEITY_CURSE_WEREBEAST_") {
				w.HistoricalFigures[d.TargetHfid].Werebeast = true
				w.HistoricalFigures[d.TargetHfid].WerebeastSince = e.Year
			}
			if strings.HasPrefix(d.Interaction, "DEITY_CURSE_VAMPIRE_") {
				w.HistoricalFigures[d.TargetHfid].Vampire = true
				w.HistoricalFigures[d.TargetHfid].VampireSince = e.Year
			}
		case *HistoricalEventHfLearnsSecret:
			if strings.HasPrefix(d.Interaction, "SECRET_") {
				w.HistoricalFigures[d.StudentHfid].Necromancer = true
				w.HistoricalFigures[d.StudentHfid].NecromancerSince = e.Year
			}
		case *HistoricalEventCreatedSite:
			w.addEntitySite(d.CivId, d.SiteId)
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.Sites[d.SiteId].Ruin = false
			w.Sites[d.SiteId].Owner = d.CivId
		case *HistoricalEventDestroyedSite:
			w.addEntitySite(d.DefenderCivId, d.SiteId)
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.Sites[d.SiteId].Ruin = true
		case *HistoricalEventSiteTakenOver:
			w.addEntitySite(d.AttackerCivId, d.SiteId)
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.addEntitySite(d.DefenderCivId, d.SiteId)
			w.addEntitySite(d.NewSiteCivId, d.SiteId)
			w.Sites[d.SiteId].Ruin = false
			w.Sites[d.SiteId].Owner = d.AttackerCivId
		case *HistoricalEventHfDestroyedSite:
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.addEntitySite(d.DefenderCivId, d.SiteId)
			w.Sites[d.SiteId].Ruin = true
		case *HistoricalEventReclaimSite:
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.Sites[d.SiteId].Ruin = false
			w.Sites[d.SiteId].Owner = d.CivId
		case *HistoricalEventAddHfEntityLink:
			if d.Link == HistoricalEventAddHfEntityLinkLink_Position {
				if hf, ok := w.HistoricalFigures[d.Hfid]; ok {
					for _, l := range hf.EntityPositionLink {
						if l.EntityId == d.CivId && l.StartYear == e.Year {
							l.PositionProfileId = d.PositionId
						}
					}
					for _, l := range hf.EntityFormerPositionLink {
						if l.EntityId == d.CivId && l.StartYear == e.Year {
							l.PositionProfileId = d.PositionId
						}
					}
				}
			}
		case *HistoricalEventHfReachSummit:
			id, _, _ := util.FindInMap(w.MountainPeaks, func(m *MountainPeak) bool { return m.Coords == d.Coords })
			d.MountainPeakId = id
		case *HistoricalEventCreatedWorldConstruction:
			if master, ok := w.WorldConstructions[d.MasterWcid]; ok {
				master.Parts = append(master.Parts, d.Wcid)
			}
		case *HistoricalEventBuildingProfileAcquired:
			if site, ok := w.Sites[d.SiteId]; ok {
				if property, ok := site.SiteProperties[d.BuildingProfileId]; ok {
					if structure, ok := site.Structures[property.StructureId]; ok {
						structure.Ruin = false
					}
					d.StructureId = property.StructureId
				}
			}
		case *HistoricalEventRazedStructure:
			if site, ok := w.Sites[d.SiteId]; ok {
				if structure, ok := site.Structures[d.StructureId]; ok {
					structure.Ruin = true
				}
			}
		case *HistoricalEventReplacedStructure:
			if site, ok := w.Sites[d.SiteId]; ok {
				if structure, ok := site.Structures[d.OldAbId]; ok {
					structure.Ruin = true
				}
			}
		}
	}
}

func (w *DfWorld) processCollections() {
	list := maps.Values(w.HistoricalEventCollections)
	sort.Slice(list, func(i, j int) bool { return list[i].Id_ < list[j].Id_ })

	for _, col := range list {
		for _, eventId := range col.Event {
			if e, ok := w.HistoricalEvents[eventId]; ok {
				e.Collection = col.Id_
			}
		}

		switch cd := col.Details.(type) {
		case *HistoricalEventCollectionAbduction:
			targets := make(map[int]bool)
			for _, eventId := range col.Event {
				if e, ok := w.HistoricalEvents[eventId]; ok {
					switch d := e.Details.(type) {
					case *HistoricalEventHfAbducted:
						targets[d.TargetHfid] = true
					}
				}
			}
			delete(targets, -1)
			cd.TargetHfids = util.Keys(targets)
		case *HistoricalEventCollectionBeastAttack:
			attackers := make(map[int]bool)
			for _, eventId := range col.Event {
				if e, ok := w.HistoricalEvents[eventId]; ok {
					switch d := e.Details.(type) {
					case *HistoricalEventHfSimpleBattleEvent:
						attackers[d.Group1Hfid] = true
					case *HistoricalEventHfAttackedSite:
						attackers[d.AttackerHfid] = true
					case *HistoricalEventHfDestroyedSite:
						attackers[d.AttackerHfid] = true
					case *HistoricalEventAddHfEntityLink:
						attackers[d.Hfid] = true
					case *HistoricalEventCreatureDevoured:
						attackers[d.Eater] = true
					case *HistoricalEventItemStolen:
						attackers[d.Histfig] = true
					}
				}
			}
			delete(attackers, -1)
			cd.AttackerHfIds = util.Keys(attackers)
		case *HistoricalEventCollectionJourney:
		HistoricalEventCollectionJourneyLoop:
			for _, eventId := range col.Event {
				if e, ok := w.HistoricalEvents[eventId]; ok {
					switch d := e.Details.(type) {
					case *HistoricalEventHfTravel:
						cd.TravellerHfIds = d.GroupHfid
						break HistoricalEventCollectionJourneyLoop
					}
				}
			}
		case *HistoricalEventCollectionOccasion:
			for _, eventcolId := range col.Eventcol {
				if e, ok := w.HistoricalEventCollections[eventcolId]; ok {
					switch d := e.Details.(type) {
					case *HistoricalEventCollectionCeremony:
						d.OccasionEventcol = col.Id_
					case *HistoricalEventCollectionCompetition:
						d.OccasionEventcol = col.Id_
					case *HistoricalEventCollectionPerformance:
						d.OccasionEventcol = col.Id_
					case *HistoricalEventCollectionProcession:
						d.OccasionEventcol = col.Id_
					}
				}
			}
		case *HistoricalEventCollectionWar:
			if e, ok := w.Entities[cd.AggressorEntId]; ok {
				e.Wars = append(e.Wars, col)
			}
			if e, ok := w.Entities[cd.DefenderEntId]; ok {
				e.Wars = append(e.Wars, col)
			}
		}
	}
}

func (w *DfWorld) addEntitySite(entityId, siteId int) {
	if e, ok := w.Entities[entityId]; ok {
		if !slices.Contains(e.Sites, siteId) {
			e.Sites = append(e.Sites, siteId)
		}
	}
}

func (w *DfWorld) addRelationshipEvents() {
	for _, r := range w.HistoricalEventRelationships {
		w.HistoricalEvents[r.Event] = &HistoricalEvent{
			Id_:        r.Event,
			Year:       r.Year,
			Collection: -1,
			Seconds72:  -1,
			Details: &HistoricalEventAddHfHfLink{
				Hfid:         r.SourceHf,
				HfidTarget:   r.TargetHf,
				Relationship: r.Relationship,
				LinkType:     HistoricalEventAddHfHfLinkLinkType_Unknown,
			},
		}
	}
}

func (w *DfWorld) processHistoricalFigures() {
	// for _, hf := range w.HistoricalFigures {
	// 	for _, i := range hf.ActiveInteraction {
	// 		if strings.HasPrefix(i, "DEITY_CURSE_WEREBEAST_") {
	// 			hf.Werebeast = true
	// 		}
	// 		if strings.HasPrefix(i, "DEITY_CURSE_VAMPIRE_") {
	// 			hf.Vampire = true
	// 		}
	// 		if strings.HasPrefix(i, "SECRET_") {
	// 			hf.Necromancer = true
	// 		}
	// 	}
	// }
}
