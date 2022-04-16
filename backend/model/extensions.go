package model

func (e *Entity) Position(id int) *EntityPosition {
	for _, p := range e.EntityPosition {
		if p.Id_ == id {
			return p
		}
	}
	return nil
}

type HistoricalEventDetails interface {
	RelatedToHf(int) bool
}

type HistoricalEventCollectionDetails interface {
}
