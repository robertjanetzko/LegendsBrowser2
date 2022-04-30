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

	var results []SearchResult

	results = seachMap(term, h.server.context.world.HistoricalFigures, results, "/hf")
	results = seachMap(term, h.server.context.world.Entities, results, "/entity")
	results = seachMap(term, h.server.context.world.Sites, results, "/site")
	for _, site := range h.server.context.world.Sites {
		results = seachMap(term, site.Structures, results, fmt.Sprintf("/site/%d/structure", site.Id_))
	}
	results = seachMap(term, h.server.context.world.Regions, results, "/region")
	results = seachMap(term, h.server.context.world.Artifacts, results, "/artifavt")
	results = seachMap(term, h.server.context.world.WorldConstructions, results, "/worldconstruction")
	results = seachMap(term, h.server.context.world.DanceForms, results, "/danceForm")
	results = seachMap(term, h.server.context.world.MusicalForms, results, "/musicalForm")
	results = seachMap(term, h.server.context.world.PoeticForms, results, "/poeticForm")
	results = seachMap(term, h.server.context.world.WrittenContents, results, "/writtencontent")
	results = seachMap(term, h.server.context.world.Landmasses, results, "/landmass")
	results = seachMap(term, h.server.context.world.MountainPeaks, results, "/mountain")

	sort.Slice(results, func(i, j int) bool { return results[i].Label < results[j].Label })

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func seachMap[T model.Named](s string, input map[int]T, output []SearchResult, baseUrl string) []SearchResult {
	for id, v := range input {
		if strings.Contains(v.Name(), s) {
			output = append(output, SearchResult{
				Label: util.Title(v.Name()),
				Value: fmt.Sprintf("%s/%d", baseUrl, id),
			})
		}
	}
	return output
}
