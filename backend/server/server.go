package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

type DfServerContext struct {
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

func StartServer(world *model.DfWorld, static embed.FS) {
	srv := &DfServer{
		router: mux.NewRouter().StrictSlash(true),
		context: &DfServerContext{
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

	srv.RegisterWorldPage("/writtencontents", "writtencontents.html", func(p Parms) any { return groupByType(srv.context.world.WrittenContents) })
	srv.RegisterWorldResourcePage("/writtencontent/{id}", "writtencontent.html", func(id int) any { return srv.context.world.WrittenContents[id] })
	srv.RegisterWorldResourcePage("/popover/writtencontent/{id}", "popoverWrittencontent.html", func(id int) any { return srv.context.world.WrittenContents[id] })

	srv.RegisterWorldResourcePage("/hf/{id}", "hf.html", func(id int) any { return srv.context.world.HistoricalFigures[id] })
	srv.RegisterWorldResourcePage("/popover/hf/{id}", "popoverHf.html", func(id int) any { return srv.context.world.HistoricalFigures[id] })

	srv.RegisterWorldPage("/events", "eventTypes.html", func(p Parms) any { return srv.context.world.AllEventTypes() })
	srv.RegisterWorldPage("/events/{type}", "eventType.html", func(p Parms) any { return srv.context.world.EventsOfType(p["type"]) })

	srv.RegisterWorldPage("/", "index.html", func(p Parms) any {
		return &struct {
			Civilizations map[string][]*model.Entity
		}{
			Civilizations: groupBy(srv.context.world.Entities,
				func(e *model.Entity) string { return e.Race },
				func(e *model.Entity) bool { return e.Name() != "" && e.Type_ == model.EntityType_Civilization },
				func(e *model.Entity) string { return e.Name() }),
		}
	})

	srv.router.PathPrefix("/search").Handler(searchHandler{server: srv})

	srv.router.PathPrefix("/load").Handler(srv.loader)

	spa := spaHandler{server: srv, staticFS: static, staticPath: "static", indexPath: "index.html"}
	srv.router.PathPrefix("/").Handler(spa)

	OpenBrowser("http://localhost:8080")

	fmt.Println("Serving at :8080")
	http.ListenAndServe(":8080", srv.router)
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

type spaHandler struct {
	server     *DfServer
	staticFS   embed.FS
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path := r.URL.Path
	// if err != nil {
	// 	// if we failed to get the absolute path respond with a 400 bad request and stop
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// prepend the path with the path to the static directory
	path = h.staticPath + path

	_, err := h.staticFS.Open(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		fmt.Println(path)
		index, err := h.staticFS.ReadFile(h.staticPath + "/" + h.indexPath)
		if err != nil {
			h.server.notFound(w)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusAccepted)
		w.Write(index)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the subdirectory of the static dir
	statics, err := fs.Sub(h.staticFS, h.staticPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.FS(statics)).ServeHTTP(w, r)
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

func groupByType[K comparable, T namedTyped](input map[K]T) map[string][]T {
	return groupBy(input, func(t T) string { return t.Type() }, func(t T) bool { return t.Name() != "" }, func(t T) string { return t.Name() })
}

func groupBy[K comparable, T any](input map[K]T, mapper func(T) string, filter func(T) bool, sortMapper func(T) string) map[string][]T {
	output := make(map[string][]T)

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
