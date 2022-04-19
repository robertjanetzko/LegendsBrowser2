package model

import "fmt"

func (e *Entity) Position(id int) *EntityPosition {
	for _, p := range e.EntityPosition {
		if p.Id_ == id {
			return p
		}
	}
	return nil
}

type HistoricalEventDetails interface {
	RelatedToEntity(int) bool
	RelatedToHf(int) bool
	Html() string
}

type HistoricalEventCollectionDetails interface {
}

func containsInt(list []int, id int) bool {
	for _, v := range list {
		if v == id {
			return true
		}
	}
	return false
}

var world *DfWorld

func hf(id int) string {
	if x, ok := world.HistoricalFigures[id]; ok {
		return fmt.Sprintf(`<a href="/hf/%d">%s</a>`, x.Id(), x.Name())
	}
	return "UNKNOWN HISTORICAL FIGURE"
}

func site(id int, prefix string) string {
	if x, ok := world.Sites[id]; ok {
		return fmt.Sprintf(`%s <a href="/site/%d">%s</a>`, prefix, x.Id(), x.Name())
	}
	return "UNKNOWN SITE"
}
