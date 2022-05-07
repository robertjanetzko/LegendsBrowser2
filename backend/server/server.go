package server

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
	"golang.org/x/exp/constraints"
)

type DfServerContext struct {
	config    *Config
	world     *model.DfWorld
	isLoading bool
	progress  *model.LoadProgress
}

type DfServer struct {
	router    *mux.Router
	loader    *loadHandler
	templates *templates.Template
	context   *DfServerContext
}

func StartServer(config *Config, world *model.DfWorld, static embed.FS) error {
	srv := &DfServer{
		router: mux.NewRouter().StrictSlash(true),
		context: &DfServerContext{
			config:    config,
			world:     world,
			isLoading: false,
			progress:  &model.LoadProgress{},
		},
	}
	srv.loader = &loadHandler{server: srv}
	srv.LoadTemplates()

	srv.RegisterWorldPage("/entities", "entities.html", func(p Parms) any { return groupByType(srv.context.world.Entities) })
	srv.RegisterWorldResourcePage("/entity/{id}", "entity.html", func(id int) any { return srv.context.world.Entities[id] })
	srv.RegisterWorldResourcePage("/popover/entity/{id}", "popoverEntity.html", func(id int) any { return srv.context.world.Entities[id] })

	srv.RegisterWorldPage("/geography", "geography.html", func(p Parms) any {
		return &struct {
			Regions       map[string][]*model.Region
			Landmasses    map[string][]*model.Landmass
			MountainPeaks map[string][]*model.MountainPeak
			Rivers        map[string][]*model.River
		}{
			Regions:       singleGroup(srv.context.world.Regions, "region"),
			Landmasses:    singleGroup(srv.context.world.Landmasses, "landmass"),
			MountainPeaks: singleGroup(srv.context.world.MountainPeaks, "mountain"),
			Rivers: map[string][]*model.River{
				"rivers": srv.context.world.Rivers,
			},
		}
	})
	srv.RegisterWorldResourcePage("/landmass/{id}", "landmass.html", func(id int) any { return srv.context.world.Landmasses[id] })
	srv.RegisterWorldResourcePage("/popover/landmass/{id}", "popoverLandmass.html", func(id int) any { return srv.context.world.Landmasses[id] })

	srv.RegisterWorldResourcePage("/mountain/{id}", "mountain.html", func(id int) any { return srv.context.world.MountainPeaks[id] })
	srv.RegisterWorldResourcePage("/popover/mountain/{id}", "popoverMountain.html", func(id int) any { return srv.context.world.MountainPeaks[id] })

	srv.RegisterWorldResourcePage("/river/{id}", "river.html", func(id int) any { return srv.context.world.Rivers[id] })
	srv.RegisterWorldResourcePage("/popover/river/{id}", "popoverRiver.html", func(id int) any { return srv.context.world.Rivers[id] })

	srv.RegisterWorldPage("/regions", "regions.html", func(p Parms) any { return groupByType(srv.context.world.Regions) })
	srv.RegisterWorldResourcePage("/region/{id}", "region.html", func(id int) any { return srv.context.world.Regions[id] })
	srv.RegisterWorldResourcePage("/popover/region/{id}", "popoverRegion.html", func(id int) any { return srv.context.world.Regions[id] })

	srv.RegisterWorldPage("/sites", "sites.html", func(p Parms) any { return groupByType(srv.context.world.Sites) })
	srv.RegisterWorldResourcePage("/site/{id}", "site.html", func(id int) any { return srv.context.world.Sites[id] })
	srv.RegisterWorldResourcePage("/popover/site/{id}", "popoverSite.html", func(id int) any { return srv.context.world.Sites[id] })

	srv.RegisterWorldPage("/structures", "structures.html", func(p Parms) any {
		return flatGrouped(srv.context.world.Sites, func(s *model.Site) []*model.Structure { return util.Values(s.Structures) })
	})
	srv.RegisterWorldPage("/site/{siteId}/structure/{id}", "structure.html", srv.findStructure)
	srv.RegisterWorldPage("/popover/site/{siteId}/structure/{id}", "popoverStructure.html", srv.findStructure)

	srv.RegisterWorldPage("/worldconstructions", "worldconstructions.html", func(p Parms) any { return groupByType(srv.context.world.WorldConstructions) })
	srv.RegisterWorldResourcePage("/worldconstruction/{id}", "worldconstruction.html", func(id int) any { return srv.context.world.WorldConstructions[id] })
	srv.RegisterWorldResourcePage("/popover/worldconstruction/{id}", "popoverWorldconstruction.html", func(id int) any { return srv.context.world.WorldConstructions[id] })

	srv.RegisterWorldPage("/artifacts", "artifacts.html", func(p Parms) any { return groupByType(srv.context.world.Artifacts) })
	srv.RegisterWorldResourcePage("/artifact/{id}", "artifact.html", func(id int) any { return srv.context.world.Artifacts[id] })
	srv.RegisterWorldResourcePage("/popover/artifact/{id}", "popoverArtifact.html", func(id int) any { return srv.context.world.Artifacts[id] })

	srv.RegisterWorldPage("/artforms", "artforms.html", func(p Parms) any {
		return &struct {
			DanceForms   map[string][]*model.DanceForm
			MusicalForms map[string][]*model.MusicalForm
			PoeticForms  map[string][]*model.PoeticForm
		}{
			DanceForms:   groupByType(srv.context.world.DanceForms),
			MusicalForms: groupByType(srv.context.world.MusicalForms),
			PoeticForms:  groupByType(srv.context.world.PoeticForms),
		}
	})

	srv.RegisterWorldResourcePage("/danceform/{id}", "artform.html", func(id int) any { return srv.context.world.DanceForms[id] })
	srv.RegisterWorldResourcePage("/musicalform/{id}", "artform.html", func(id int) any { return srv.context.world.MusicalForms[id] })
	srv.RegisterWorldResourcePage("/poeticform/{id}", "artform.html", func(id int) any { return srv.context.world.PoeticForms[id] })

	srv.RegisterWorldPage("/writtencontents", "writtencontents.html", func(p Parms) any { return groupByType(srv.context.world.WrittenContents) })
	srv.RegisterWorldResourcePage("/writtencontent/{id}", "writtencontent.html", func(id int) any { return srv.context.world.WrittenContents[id] })
	srv.RegisterWorldResourcePage("/popover/writtencontent/{id}", "popoverWrittencontent.html", func(id int) any { return srv.context.world.WrittenContents[id] })

	srv.RegisterWorldPage("/hfs", "hfs.html", srv.searchHf)
	srv.RegisterWorldResourcePage("/hf/{id}", "hf.html", func(id int) any { return srv.context.world.HistoricalFigures[id] })
	srv.RegisterWorldResourcePage("/popover/hf/{id}", "popoverHf.html", func(id int) any { return srv.context.world.HistoricalFigures[id] })

	srv.RegisterWorldPage("/identities", "identities.html", func(p Parms) any { return srv.context.world.Identities })
	srv.RegisterWorldResourcePage("/identity/{id}", "identity.html", func(id int) any { return srv.context.world.Identities[id] })
	srv.RegisterWorldResourcePage("/popover/identity/{id}", "popoverIdentity.html", func(id int) any { return srv.context.world.Identities[id] })

	srv.RegisterWorldPage("/years", "years.html", func(p Parms) any {
		return groupBy(srv.context.world.HistoricalEvents,
			func(e *model.HistoricalEvent) int { return e.Year },
			func(e *model.HistoricalEvent) bool { return true },
			func(e *model.HistoricalEvent) int { return e.Id_ })
	})
	srv.RegisterWorldResourcePage("/year/{id}", "year.html", func(id int) any {
		return util.FilterMap(srv.context.world.HistoricalEvents,
			func(v *model.HistoricalEvent) bool { return v.Year == id },
			func(a, b *model.HistoricalEvent) bool { return a.Id_ < b.Id_ },
		)
	})

	srv.RegisterWorldPage("/events", "eventTypes.html", func(p Parms) any { return srv.context.world.AllEventTypes() })
	srv.RegisterWorldPage("/events/{type}", "eventType.html", func(p Parms) any { return srv.context.world.EventsOfType(p["type"]) })
	srv.RegisterWorldResourcePage("/event/{id}", "event.html", func(id int) any { return srv.context.world.HistoricalEvents[id] })

	srv.RegisterWorldPage("/collections", "collections.html", func(p Parms) any {
		return groupBy(srv.context.world.HistoricalEventCollections,
			func(e *model.HistoricalEventCollection) string { return e.Type() },
			func(e *model.HistoricalEventCollection) bool { return true },
			func(e *model.HistoricalEventCollection) string { return model.Time(e.StartYear, e.StartSeconds72) },
		)
	})
	srv.RegisterWorldResourcePage("/collection/{id}", "collection.html", func(id int) any { return srv.context.world.HistoricalEventCollections[id] })
	srv.RegisterWorldResourcePage("/popover/collection/{id}", "popoverCollection.html", func(id int) any { return srv.context.world.HistoricalEventCollections[id] })

	srv.RegisterWorldPage("/worldmap", "worldMap.html", func(p Parms) any {
		return &struct {
			Landmasses         map[int]*model.Landmass
			Regions            map[int]*model.Region
			Sites              map[int]*model.Site
			MountainPeaks      map[int]*model.MountainPeak
			WorldConstructions map[int]*model.WorldConstruction
			Rivers             []*model.River
		}{
			Landmasses:         srv.context.world.Landmasses,
			Regions:            srv.context.world.Regions,
			Sites:              srv.context.world.Sites,
			MountainPeaks:      srv.context.world.MountainPeaks,
			WorldConstructions: srv.context.world.WorldConstructions,
			Rivers:             srv.context.world.Rivers,
		}
	})

	srv.RegisterWorldPage("/", "index.html", func(p Parms) any {
		return &struct {
			Civilizations map[string][]*model.Entity
		}{
			Civilizations: groupBy(srv.context.world.Entities,
				func(e *model.Entity) string { return util.If(e.Necromancer, "necromancer", e.Race) },
				func(e *model.Entity) bool {
					return e.Name() != "" && (e.Type_ == model.EntityType_Civilization || e.Necromancer)
				},
				func(e *model.Entity) string { return e.Name() }),
		}
	})

	srv.router.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		w.Write(srv.loader.server.context.world.MapData)
	})

	srv.router.PathPrefix("/search").Handler(searchHandler{server: srv})

	srv.router.PathPrefix("/load").Handler(srv.loader)

	spa := spaHandler{server: srv, staticFS: static, staticPath: "static", indexPath: "index.html"}
	if templates.DebugTemplates {
		spa.staticFS = os.DirFS(".")
	}
	srv.router.PathPrefix("/").Handler(spa)

	OpenBrowser("http://localhost:8080")

	fmt.Println("Serving at :8080")
	http.ListenAndServe(":8080", srv.router)
	return nil
}

func (srv *DfServer) findStructure(p Parms) any {
	siteId, err := strconv.Atoi(p["siteId"])
	if err != nil {
		return nil
	}
	structureId, err := strconv.Atoi(p["id"])
	if err != nil {
		return nil
	}
	if site, ok := srv.context.world.Sites[siteId]; ok {
		return site.Structures[structureId]
	}
	return nil
}

func (srv *DfServer) searchHf(p Parms) any {
	var list []*model.HistoricalFigure

	for _, hf := range srv.context.world.HistoricalFigures {
		if p["leader"] == "1" && !hf.Leader {
			continue
		}
		if p["deity"] == "1" && !hf.Deity {
			continue
		}
		if p["force"] == "1" && !hf.Force {
			continue
		}
		if p["vampire"] == "1" && !hf.Vampire {
			continue
		}
		if p["werebeast"] == "1" && !hf.Werebeast {
			continue
		}
		if p["necromancer"] == "1" && !hf.Necromancer {
			continue
		}
		if p["alive"] == "1" && hf.DeathYear != -1 {
			continue
		}
		if p["ghost"] == "1" && false { // TODO ghost
			continue
		}
		if p["adventurer"] == "1" && !hf.Adventurer {
			continue
		}
		list = append(list, hf)
	}

	switch p["sort"] {
	case "race":
		sort.Slice(list, func(i, j int) bool { return list[i].Race < list[j].Race })
	case "birth":
		sort.Slice(list, func(i, j int) bool { return list[i].BirthYear < list[j].BirthYear })
	case "death":
		sort.Slice(list, func(i, j int) bool { return list[i].DeathYear < list[j].DeathYear })
	case "kills":
		sort.Slice(list, func(i, j int) bool { return len(list[i].Kills) > len(list[j].Kills) })
	default:
		sort.Slice(list, func(i, j int) bool { return list[i].Name_ < list[j].Name_ })
	}

	return map[string]any{
		"Params": p,
		"Hfs":    list,
	}
}

func (srv *DfServer) notFound(w http.ResponseWriter) {
	err := srv.templates.Render(w, "notFound.html", nil)
	if err != nil {
		httpError(w, err)
	}
}

func httpError(w http.ResponseWriter, err error) {
	fmt.Fprintln(w, err)
	fmt.Println(err)
}

type namedTyped interface {
	model.Named
	model.Typed
}

func flatGrouped[K comparable, U any, V namedTyped](input map[K]U, mapper func(U) []V) map[string][]V {
	output := make(map[string][]V)

	for _, x := range input {
		for _, v := range mapper(x) {
			k := v.Type()
			if v.Name() != "" {
				output[k] = append(output[k], v)
			}
		}
	}

	for _, v := range output {
		sort.Slice(v, func(i, j int) bool { return v[i].Name() < v[j].Name() })
	}

	return output
}

func singleGroup[K comparable, T model.Named](input map[K]T, group string) map[string][]T {
	return groupBy(input, func(t T) string { return group }, func(t T) bool { return t.Name() != "" }, func(t T) string { return t.Name() })
}

func groupByType[K comparable, T namedTyped](input map[K]T) map[string][]T {
	return groupBy(input, func(t T) string { return t.Type() }, func(t T) bool { return t.Name() != "" }, func(t T) string { return t.Name() })
}

func groupBy[K comparable, N comparable, T any, S constraints.Ordered](input map[K]T, mapper func(T) N, filter func(T) bool, sortMapper func(T) S) map[N][]T {
	output := make(map[N][]T)

	for _, v := range input {
		k := mapper(v)
		if filter(v) {
			output[k] = append(output[k], v)
		}
	}

	for _, v := range output {
		sort.Slice(v, func(i, j int) bool { return sortMapper(v[i]) < sortMapper(v[j]) })
	}

	return output
}
