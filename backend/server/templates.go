package server

import (
	"fmt"
	"html/template"
	"net/url"

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
		"title":     util.Title,
		"hf":        func(id int) template.HTML { return model.LinkHf(srv.context.world, id) },
		"getHf":     func(id int) *model.HistoricalFigure { return srv.context.world.HistoricalFigures[id] },
		"entity":    func(id int) template.HTML { return model.LinkEntity(srv.context.world, id) },
		"getEntity": func(id int) *model.Entity { return srv.context.world.Entities[id] },
		"site":      func(id int) template.HTML { return model.LinkSite(srv.context.world, id) },
		"getSite":   func(id int) *model.Site { return srv.context.world.Sites[id] },
		"region":    func(id int) template.HTML { return model.LinkRegion(srv.context.world, id) },
		"getRegion": func(id int) *model.Region { return srv.context.world.Regions[id] },
		"events": func(obj any) *model.EventList {
			fmt.Println("W", srv.context.world)
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
	}
	srv.templates = templates.New(functions)
}
