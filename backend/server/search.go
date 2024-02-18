package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

type searchHandler struct {
	server *DfServer
}

type SearchResult struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (h searchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.server.context.world == nil {
		h.server.renderLoading(w, r)
		return
	}

	term := r.URL.Query().Get("term")

	world := h.server.context.world
	if term != "" {
		var results []SearchResult
		results = searchMap(term, world.HistoricalFigures, results, "/hf")
		results = searchMap(term, world.Entities, results, "/entity")
		results = searchMap(term, world.Sites, results, "/site")
		for _, site := range world.Sites {
			results = searchMap(term, site.Structures, results, fmt.Sprintf("/site/%d/structure", site.Id_))
		}
		results = searchMap(term, world.Regions, results, "/region")
		results = searchMap(term, world.Artifacts, results, "/artifact")
		results = searchMap(term, world.WorldConstructions, results, "/worldconstruction")
		results = searchMap(term, world.DanceForms, results, "/danceform")
		results = searchMap(term, world.MusicalForms, results, "/musicalform")
		results = searchMap(term, world.PoeticForms, results, "/poeticform")
		results = searchMap(term, world.WrittenContents, results, "/writtencontent")
		results = searchMap(term, world.Landmasses, results, "/landmass")
		results = searchMap(term, world.MountainPeaks, results, "/mountain")

		sort.Slice(results, func(i, j int) bool { return results[i].Label < results[j].Label })

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	} else {
		term = r.URL.Query().Get("search")

		var structures []*model.Structure
		for _, site := range world.Sites {
			structures = search(term, site.Structures, structures)
		}

		results := struct {
			Term               string
			HistoricalFigures  []*model.HistoricalFigure
			Entities           []*model.Entity
			Sites              []*model.Site
			Structures         []*model.Structure
			Regions            []*model.Region
			Artifacts          []*model.Artifact
			WorldConstructions []*model.WorldConstruction
			DanceForms         []*model.DanceForm
			MusicalForms       []*model.MusicalForm
			PoeticForms        []*model.PoeticForm
			WrittenContents    []*model.WrittenContent
			Landmasses         []*model.Landmass
			MountainPeaks      []*model.MountainPeak
		}{
			Term:               term,
			HistoricalFigures:  search(term, world.HistoricalFigures, nil),
			Entities:           search(term, world.Entities, nil),
			Sites:              search(term, world.Sites, nil),
			Structures:         structures,
			Regions:            search(term, world.Regions, nil),
			Artifacts:          search(term, world.Artifacts, nil),
			WorldConstructions: search(term, world.WorldConstructions, nil),
			DanceForms:         search(term, world.DanceForms, nil),
			MusicalForms:       search(term, world.MusicalForms, nil),
			PoeticForms:        search(term, world.PoeticForms, nil),
			WrittenContents:    search(term, world.WrittenContents, nil),
			Landmasses:         search(term, world.Landmasses, nil),
			MountainPeaks:      search(term, world.MountainPeaks, nil),
		}

		err := h.server.templates.Render(w, "search.html", results)
		if err != nil {
			httpError(w, err)
		}
	}
}

func searchMap[T model.Named](s string, input map[int]T, output []SearchResult, baseUrl string) []SearchResult {
	s = strings.ToLower(s)
	for id, v := range input {
		if strings.Contains(strings.ToLower(v.Name()), s) {
			output = append(output, SearchResult{
				Label: util.Title(v.Name()),
				Value: fmt.Sprintf("%s/%d", baseUrl, id),
			})
		}
	}
	return output
}

func search[T model.Named](s string, input map[int]T, output []T) []T {
	s = strings.ToLower(s)
	for _, v := range input {
		if strings.Contains(strings.ToLower(v.Name()), s) {
			output = append(output, v)
		}
	}
	sort.Slice(output, func(i, j int) bool { return output[i].Name() < output[j].Name() })
	return output
}
