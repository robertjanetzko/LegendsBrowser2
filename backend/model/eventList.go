package model

import (
	"fmt"
)

type HistoricalEventDetails interface {
	RelatedToEntity(int) bool
	RelatedToHf(int) bool
	RelatedToArtifact(int) bool
	RelatedToSite(int) bool
	RelatedToRegion(int) bool
	Html(*Context) string
	Type() string
}

type HistoricalEventCollectionDetails interface {
}

type EventList struct {
	Events  []*HistoricalEvent
	Context *Context
}

func NewEventList(world *DfWorld, obj any) *EventList {
	el := EventList{
		Context: &Context{HfId: -1},
	}

	switch x := obj.(type) {
	case *Entity:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToEntity(x.Id()) })
	case *HistoricalFigure:
		el.Context.HfId = x.Id()
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToHf(x.Id()) })
	case *Artifact:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToArtifact(x.Id()) })
	case *Site:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToSite(x.Id()) })
	case *Region:
		el.Events = world.EventsMatching(func(d HistoricalEventDetails) bool { return d.RelatedToRegion(x.Id()) })
	case []*HistoricalEvent:
		el.Events = x
	default:
		fmt.Printf("unknown type %T\n", obj)
	}

	return &el
}
