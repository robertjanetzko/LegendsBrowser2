package model

import (
	"fmt"

	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

type HistoricalEventDetails interface {
	RelatedToEntity(int) bool
	RelatedToHf(int) bool
	RelatedToArtifact(int) bool
	RelatedToSite(int) bool
	RelatedToStructure(int, int) bool
	RelatedToRegion(int) bool
	RelatedToWorldConstruction(int) bool
	RelatedToWrittenContent(int) bool
	RelatedToDanceForm(int) bool
	RelatedToMusicalForm(int) bool
	RelatedToPoeticForm(int) bool
	RelatedToMountain(int) bool
	Html(*Context) string
	Type() string
}

type HistoricalEventCollectionDetails interface {
	Type() string
	Html(*HistoricalEventCollection, *Context) string
}

type EventList struct {
	Events  []*HistoricalEvent
	Context *Context
}

func NewEventList(world *DfWorld, obj any) *EventList {
	el := EventList{
		Context: &Context{
			World: world,
			HfId:  -1,
		},
	}

	switch x := obj.(type) {
	case *Entity:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToEntity(x.Id()) })
	case *HistoricalFigure:
		el.Context.HfId = x.Id()
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToHf(x.Id()) })
	case *Artifact:
		el.Context.HfId = x.HolderHfid
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToArtifact(x.Id()) })
	case *Site:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToSite(x.Id()) })
	case *Structure:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToStructure(x.SiteId, x.Id()) })
	case *Region:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToRegion(x.Id()) })
	case *WorldConstruction:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToWorldConstruction(x.Id()) })
	case *WrittenContent:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToWrittenContent(x.Id()) })
	case *DanceForm:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToDanceForm(x.Id()) })
	case *MusicalForm:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToMusicalForm(x.Id()) })
	case *PoeticForm:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToPoeticForm(x.Id()) })
	case *MountainPeak:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToMountain(x.Id()) })
	case []*HistoricalEvent:
		el.Events = x
	case []int:
		el.Events = util.Map(x, func(id int) *HistoricalEvent { return world.HistoricalEvents[id] })
	default:
		fmt.Printf("unknown type %T\n", obj)
	}

	return &el
}
