package model

import (
	"fmt"
	"strings"
)

func (w *DfWorld) process() {

	// set site in structure
	for _, site := range w.Sites {
		for _, structure := range site.Structures {
			structure.SiteId = site.Id_
		}
	}

	w.processEvents()

	// check events texts
	for _, e := range w.HistoricalEvents {
		e.Details.Html(&Context{World: w})
	}
}

func (w *DfWorld) processEvents() {
	for _, e := range w.HistoricalEvents {
		switch d := e.Details.(type) {
		case *HistoricalEventHfDoesInteraction:
			if strings.HasPrefix(d.Interaction, "DEITY_CURSE_WEREBEAST_") {
				w.HistoricalFigures[d.TargetHfid].Werebeast = true
			}
			if strings.HasPrefix(d.Interaction, "DEITY_CURSE_VAMPIRE_") {
				w.HistoricalFigures[d.TargetHfid].Vampire = true
			}
		case *HistoricalEventCreatedSite:
			w.addEntitySite(d.CivId, d.SiteId)
			w.addEntitySite(d.SiteCivId, d.SiteId)
		case *HistoricalEventDestroyedSite:
			w.addEntitySite(d.DefenderCivId, d.SiteId)
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.Sites[d.SiteId].Ruin = true
		case *HistoricalEventSiteTakenOver:
			w.addEntitySite(d.AttackerCivId, d.SiteId)
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.addEntitySite(d.DefenderCivId, d.SiteId)
			w.addEntitySite(d.NewSiteCivId, d.SiteId)
		case *HistoricalEventHfDestroyedSite:
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.addEntitySite(d.DefenderCivId, d.SiteId)
			w.Sites[d.SiteId].Ruin = true
		case *HistoricalEventReclaimSite:
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.addEntitySite(d.SiteCivId, d.SiteId)
			w.Sites[d.SiteId].Ruin = false
		}
	}
}

func (w *DfWorld) addEntitySite(entityId, siteId int) {
	fmt.Println("add site", entityId, siteId)
	if e, ok := w.Entities[entityId]; ok {
		if s, ok := w.Sites[siteId]; ok {
			e.Sites[s.Id_] = s
		}
	}
}
