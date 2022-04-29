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

func (srv *DfServer) LoadTemplates() {
	functions := template.FuncMap{
		"json": util.Json,
		"check": func(condition bool, v any) any {
			if condition {
				return v
			}
			return nil
		},
		"title":                util.Title,
		"kebab":                func(s string) string { return strcase.ToKebab(s) },
		"hf":                   func(id int) template.HTML { return model.LinkHf(srv.context.world, id) },
		"getHf":                func(id int) *model.HistoricalFigure { return srv.context.world.HistoricalFigures[id] },
		"entity":               func(id int) template.HTML { return model.LinkEntity(srv.context.world, id) },
		"getEntity":            func(id int) *model.Entity { return srv.context.world.Entities[id] },
		"site":                 func(id int) template.HTML { return model.LinkSite(srv.context.world, id) },
		"getSite":              func(id int) *model.Site { return srv.context.world.Sites[id] },
		"structure":            func(siteId, id int) template.HTML { return model.LinkStructure(srv.context.world, siteId, id) },
		"region":               func(id int) template.HTML { return model.LinkRegion(srv.context.world, id) },
		"getRegion":            func(id int) *model.Region { return srv.context.world.Regions[id] },
		"worldconstruction":    func(id int) template.HTML { return model.LinkWorldConstruction(srv.context.world, id) },
		"getWorldconstruction": func(id int) *model.WorldConstruction { return srv.context.world.WorldConstructions[id] },
		"artifact":             func(id int) template.HTML { return model.LinkArtifact(srv.context.world, id) },
		"getArtifact":          func(id int) *model.Artifact { return srv.context.world.Artifacts[id] },
		"danceForm":            func(id int) template.HTML { return model.LinkDanceForm(srv.context.world, id) },
		"musicalForm":          func(id int) template.HTML { return model.LinkMusicalForm(srv.context.world, id) },
		"poeticForm":           func(id int) template.HTML { return model.LinkPoeticForm(srv.context.world, id) },
		"writtencontent":       func(id int) template.HTML { return model.LinkWrittenContent(srv.context.world, id) },
		"events": func(obj any) *model.EventList {
			return model.NewEventList(srv.context.world, obj)
		},
		"season":       model.Season,
		"time":         model.Time,
		"url":          url.PathEscape,
		"query":        url.QueryEscape,
		"isLegendsXml": isLegendsXml,
		"html": func(value any) template.HTML {
			return template.HTML(fmt.Sprint(value))
		},
		"bytes":   func(s int64) string { return humanize.Bytes(uint64(s)) },
		"first":   util.FirstInMap,
		"ifFirst": func(m any, k string, r string) string { return util.If(util.FirstInMap(m, k), r, "") },
	}
	srv.templates = templates.New(functions)
}
