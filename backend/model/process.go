package model

func (w *DfWorld) Process() {

	// set site in structure
	for _, site := range w.Sites {
		for _, structure := range site.Structures {
			structure.SiteId = site.Id_
		}
	}

	// check events texts
	for _, e := range w.HistoricalEvents {
		e.Details.Html(&Context{World: w})
	}
}
