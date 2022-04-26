package server

import (
	"fmt"
	"html/template"

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
		"hf":        func(id int) template.HTML { return model.LinkHf(srv.world, id) },
		"getHf":     func(id int) *model.HistoricalFigure { return srv.world.HistoricalFigures[id] },
		"entity":    func(id int) template.HTML { return model.LinkEntity(srv.world, id) },
		"getEntity": func(id int) *model.Entity { return srv.world.Entities[id] },
		"site":      func(id int) template.HTML { return model.LinkSite(srv.world, id) },
		"getSite":   func(id int) *model.Site { return srv.world.Sites[id] },
		"region":    func(id int) template.HTML { return model.LinkRegion(srv.world, id) },
		"getRegion": func(id int) *model.Region { return srv.world.Regions[id] },
		"events":    func(obj any) *model.EventList { return model.NewEventList(srv.world, obj) },
		"season":    model.Season,
		"time":      model.Time,
		"html": func(value any) template.HTML {
			return template.HTML(fmt.Sprint(value))
		},
	}
	srv.templates = templates.New(functions)
}
