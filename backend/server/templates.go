package server

import (
	"fmt"
	"html/template"
	"net/url"

	humanize "github.com/dustin/go-humanize"
	"github.com/iancoleman/strcase"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

var DebugJSON = false

func (srv *DfServer) LoadTemplates() {
	functions := template.FuncMap{
		"json": func(obj any) template.HTML {
			if !DebugJSON {
				return ""
			} else {
				return util.Json(obj)
			}
		},
		"check": func(condition bool, v any) any {
			if condition {
				return v
			}
			return nil
		},
		"title":   util.Title,
		"kebab":   func(s string) string { return strcase.ToKebab(s) },
		"andList": model.AndList,
		"world":   func() *model.DfWorld { return srv.context.world },
		"context": func(r any) *model.Context { return model.NewContext(srv.context.world, r) },
		"initMap": func() template.HTML {
			return template.HTML(fmt.Sprintf(`<script>var worldWidth = %d, worldHeight = %d;</script><script src="/js/map.js"></script>`,
				srv.context.world.Width, srv.context.world.Height))
		},
		"hf":                   func(id int) template.HTML { return model.LinkHf(srv.context.world, id) },
		"hfShort":              func(id int) template.HTML { return model.LinkHfShort(srv.context.world, id) },
		"getHf":                func(id int) *model.HistoricalFigure { return srv.context.world.HistoricalFigures[id] },
		"hfList":               func(ids []int) template.HTML { return model.LinkHfList(srv.context.world, ids) },
		"identity":             func(id int) template.HTML { return model.LinkIdentity(srv.context.world, id) },
		"entity":               func(id int) template.HTML { return model.LinkEntity(srv.context.world, id) },
		"getEntity":            func(id int) *model.Entity { return srv.context.world.Entities[id] },
		"site":                 func(id int) template.HTML { return model.LinkSite(srv.context.world, id) },
		"getSite":              func(id int) *model.Site { return srv.context.world.Sites[id] },
		"structure":            func(siteId, id int) template.HTML { return model.LinkStructure(srv.context.world, siteId, id) },
		"region":               func(id int) template.HTML { return model.LinkRegion(srv.context.world, id) },
		"getRegion":            func(id int) *model.Region { return srv.context.world.Regions[id] },
		"worldConstruction":    func(id int) template.HTML { return model.LinkWorldConstruction(srv.context.world, id) },
		"getWorldConstruction": func(id int) *model.WorldConstruction { return srv.context.world.WorldConstructions[id] },
		"artifact":             func(id int) template.HTML { return model.LinkArtifact(srv.context.world, id) },
		"getArtifact":          func(id int) *model.Artifact { return srv.context.world.Artifacts[id] },
		"danceForm":            func(id int) template.HTML { return model.LinkDanceForm(srv.context.world, id) },
		"musicalForm":          func(id int) template.HTML { return model.LinkMusicalForm(srv.context.world, id) },
		"poeticForm":           func(id int) template.HTML { return model.LinkPoeticForm(srv.context.world, id) },
		"writtenContent":       func(id int) template.HTML { return model.LinkWrittenContent(srv.context.world, id) },
		"landmass":             func(id int) template.HTML { return model.LinkLandmass(srv.context.world, id) },
		"mountain":             func(id int) template.HTML { return model.LinkMountain(srv.context.world, id) },
		"river":                func(id int) template.HTML { return model.LinkRiver(srv.context.world, id) },

		"addLandmass":          func(id int) template.HTML { return model.AddMapLandmass(srv.context.world, id) },
		"addRegion":            func(id int) template.HTML { return model.AddMapRegion(srv.context.world, id) },
		"addSite":              func(id int, color bool) template.HTML { return model.AddMapSite(srv.context.world, id, color) },
		"addMountain":          func(id int, color bool) template.HTML { return model.AddMapMountain(srv.context.world, id, color) },
		"addWorldConstruction": func(id int) template.HTML { return model.AddMapWorldConstruction(srv.context.world, id) },
		"addRiver":             func(id int) template.HTML { return model.AddMapRiver(srv.context.world, id) },

		"events": func(obj any) *model.EventList {
			return model.NewEventList(srv.context.world, obj)
		},
		"history": func(siteId int) []*model.HistoricalEvent {
			return srv.context.world.SiteHistory(siteId)
		},
		"collection":    func(id int) template.HTML { return model.LinkCollection(srv.context.world, id) },
		"getCollection": func(id int) *model.HistoricalEventCollection { return srv.context.world.HistoricalEventCollections[id] },
		"getOccasion": func(civId, occasionId int) *model.Occasion {
			if civ, ok := srv.context.world.Entities[civId]; ok {
				return civ.Occasion[occasionId]
			}
			return nil
		},
		"story": func(id int) template.HTML {
			if e, ok := srv.context.world.HistoricalEvents[id]; ok {
				return template.HTML(e.Details.Html(&model.Context{World: srv.context.world, Story: true}) + " in " + model.Time(e.Year, e.Seconds72))
			}
			return template.HTML("")
		},
		"description":  func(d string) template.HTML { return model.LinkDescription(srv.context.world, d) },
		"season":       model.Season,
		"time":         model.Time,
		"url":          url.PathEscape,
		"query":        url.QueryEscape,
		"isLegendsXml": isLegendsXml,
		"html": func(value any) template.HTML {
			return template.HTML(fmt.Sprint(value))
		},
		"bytes":           func(s int64) string { return humanize.Bytes(uint64(s)) },
		"first":           util.FirstInMap,
		"ifFirst":         func(m any, k string, r string) string { return util.If(util.FirstInMap(m, k), r, "") },
		"strip":           util.Strip,
		"string":          util.String,
		"capitalize":      util.Capitalize,
		"add":             func(a, b int) int { return a + b },
		"breakYearColumn": func(c, m int) bool { return (c % ((m + 2) / 4)) == 0 },
	}
	srv.templates = templates.New(functions)
}
