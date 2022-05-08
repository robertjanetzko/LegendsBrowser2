package model

import (
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func (w *DfWorld) process() {
	for id, r := range w.Rivers {
		r.Id_ = id
	}

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

	if !w.Plus {
		trimRace := func(s string) string { return strings.Trim(strcase.ToDelimited(s, ' '), " 0123456789") }
		for _, hf := range w.HistoricalFigures {
			hf.Race = trimRace(hf.Race)
		}

		for _, e := range w.Entities {
			if len(e.Leaders) > 0 {
				switch r := e.Leaders[0].Hf.Race; {
				case r == "demon":
					e.Race = "goblin"
				default:
					e.Race = r
				}
			} else {
				if !strings.Contains(e.Name_, " ") {
					e.Race = "kobold"
				}
			}
		}

		for _, a := range w.DanceForms {
			a.Name_ = a.Description[:strings.Index(a.Description, " is a ")]
		}
		for _, a := range w.MusicalForms {
			a.Name_ = a.Description[:strings.Index(a.Description, " is a ")]
		}
		for _, a := range w.PoeticForms {
			a.Name_ = a.Description[:strings.Index(a.Description, " is a ")]
		}

		setEntityType := func(id int, t EntityType) {
			if c, ok := w.Entities[id]; ok {
				if c.Type_ == EntityType_Unknown {
					c.Type_ = t
				}
			}
		}
		setParent := func(id, parent int) {
			if id == -1 || parent == -1 {
				return
			}
			if c, ok := w.Entities[parent]; ok {
				if c.Parent != -1 {
					parent = c.Parent
				}
			}
			if c, ok := w.Entities[id]; ok {
				c.Parent = parent
				if p, ok := w.Entities[parent]; ok {
					c.Race = p.Race
				}
			}
			if c, ok := w.Entities[parent]; ok {
				c.Child = append(c.Child, id)
			}
		}

		list := maps.Values(w.HistoricalEvents)
		sort.Slice(list, func(i, j int) bool { return list[i].Id_ < list[j].Id_ })
		for _, e := range list {
			switch d := e.Details.(type) {
			case *HistoricalEventCreatedSite:
				setParent(d.SiteCivId, d.CivId)
				setEntityType(d.CivId, EntityType_Civilization)
				setEntityType(d.SiteCivId, EntityType_Sitegovernment)
			case *HistoricalEventDestroyedSite:
				setParent(d.SiteCivId, d.DefenderCivId)
				setEntityType(d.AttackerCivId, EntityType_Civilization)
				setEntityType(d.DefenderCivId, EntityType_Civilization)
				setEntityType(d.SiteCivId, EntityType_Sitegovernment)
			case *HistoricalEventSiteTakenOver:
				setParent(d.SiteCivId, d.DefenderCivId)
				setParent(d.NewSiteCivId, d.AttackerCivId)
				setEntityType(d.DefenderCivId, EntityType_Civilization)
				setEntityType(d.SiteCivId, EntityType_Sitegovernment)
				setEntityType(d.AttackerCivId, EntityType_Civilization)
				setEntityType(d.NewSiteCivId, EntityType_Sitegovernment)
			case *HistoricalEventHfDestroyedSite:
				setParent(d.SiteCivId, d.DefenderCivId)
				setEntityType(d.DefenderCivId, EntityType_Civilization)
				setEntityType(d.SiteCivId, EntityType_Sitegovernment)
			case *HistoricalEventReclaimSite:
				setParent(d.SiteCivId, d.CivId)
				setEntityType(d.CivId, EntityType_Civilization)
				setEntityType(d.SiteCivId, EntityType_Sitegovernment)
			case *HistoricalEventCreatedStructure:
				setParent(d.SiteCivId, d.CivId)
				setEntityType(d.CivId, EntityType_Civilization)
				setEntityType(d.SiteCivId, EntityType_Sitegovernment)
				if site, ok := w.Sites[d.SiteId]; ok {
					if structure, ok := site.Structures[d.StructureId]; ok {
						if structure.Type_ == StructureType_Guildhall {
							setEntityType(d.SiteCivId, EntityType_Guild)
						}
					}
				}
			case *HistoricalEventChangedCreatureType:
				d.NewRace = trimRace(d.NewRace)
				d.OldRace = trimRace(d.OldRace)
			case *HistoricalEventCreatureDevoured:
				if col, ok := w.HistoricalEventCollections[e.Collection]; ok {
					if cd, ok := col.Details.(*HistoricalEventCollectionBeastAttack); ok {
						if len(cd.AttackerHfIds) > 0 {
							d.Eater = cd.AttackerHfIds[0]
						}
					}
				}
				d.Race = "creature"
			case *HistoricalEventHfNewPet:
				d.Pets = "creature" // TODO from hf pets?
			case *HistoricalEventItemStolen:
				if col, ok := w.HistoricalEventCollections[e.Collection]; ok {
					if cd, ok := col.Details.(*HistoricalEventCollectionBeastAttack); ok {
						if len(cd.AttackerHfIds) > 0 {
							d.Histfig = cd.AttackerHfIds[0]
						}
						d.Site = cd.SiteId
					}
				}
				d.ItemType = "item"
			case *HistoricalEventMasterpieceItem:
				d.ItemType = "item"
			}
		}

		for _, e := range w.Entities {
			switch e.Race {
			case "dwarf":
				e.EntityPosition = dwarfPositions
			case "elf":
				e.EntityPosition = elfPositions
			case "human":
				e.EntityPosition = humanPositions
			case "goblin":
				e.EntityPosition = goblinPositions
			}
			for i, p := range e.EntityPosition {
				p.Id_ = i
			}
		}
	}

	// check events texts
	for _, e := range w.HistoricalEvents {
		e.Details.Html(&Context{World: w})
	}
	for _, e := range w.HistoricalEventCollections {
		e.Details.Html(e, &Context{World: w})
	}

}

func (w *DfWorld) processEvents() {
	list := maps.Values(w.HistoricalEvents)
	sort.Slice(list, func(i, j int) bool { return list[i].Id_ < list[j].Id_ })

	for _, e := range list {
		switch d := e.Details.(type) {
		case *HistoricalEventHfDoesInteraction:
			if hf, ok := w.HistoricalFigures[d.TargetHfid]; ok {
				if strings.HasPrefix(d.Interaction, "DEITY_CURSE_WEREBEAST_") {
					hf.Werebeast = true
					hf.WerebeastSince = e.Year
				}
				if strings.HasPrefix(d.Interaction, "DEITY_CURSE_VAMPIRE_") {
					hf.Vampire = true
					hf.VampireSince = e.Year
				}
			}
		case *HistoricalEventHfLearnsSecret:
			if strings.HasPrefix(d.Interaction, "SECRET_") {
				if hf, ok := w.HistoricalFigures[d.StudentHfid]; ok {
					hf.Necromancer = true
					hf.NecromancerSince = e.Year
				}
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
		case *HistoricalEventAssumeIdentity:
			if hf, ok := w.HistoricalFigures[d.TricksterHfid]; ok {
				if id, ok := w.Identities[d.IdentityId]; ok {
					id.HistfigId = hf.Id_
				}
			}
		case *HistoricalEventHfDied:
			if hf, ok := w.HistoricalFigures[d.SlayerHfid]; ok {
				hf.Kills = append(hf.Kills, d.Hfid)
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
