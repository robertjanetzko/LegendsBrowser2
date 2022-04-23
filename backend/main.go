package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/profile"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

var world *model.DfWorld

//go:embed static
var static embed.FS

func main() {
	f := flag.String("f", "", "open a file")
	p := flag.Bool("p", false, "start profiling")
	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)

	functions := template.FuncMap{
		"json": util.Json,
		"check": func(condition bool, v any) any {
			if condition {
				return v
			}
			return nil
		},
		"title":     util.Title,
		"hf":        model.LinkHf,
		"getHf":     func(id int) *model.HistoricalFigure { return world.HistoricalFigures[id] },
		"entity":    model.LinkEntity,
		"getEntity": func(id int) *model.Entity { return world.Entities[id] },
		"site":      model.LinkSite,
		"getSite":   func(id int) *model.Site { return world.Sites[id] },
		"region":    model.LinkRegion,
		"getRegion": func(id int) *model.Region { return world.Regions[id] },
		"events":    model.NewEventList,
		"season": func(seconds int) string {
			r := ""
			month := seconds % 100800
			if month <= 33600 {
				r += "early "
			} else if month <= 67200 {
				r += "mid"
			} else if month <= 100800 {
				r += "late "
			}

			season := seconds % 403200
			if season < 100800 {
				r += "spring"
			} else if season < 201600 {
				r += "summer"
			} else if season < 302400 {
				r += "autumn"
			} else if season < 403200 {
				r += "winter"
			}

			return r
		},
		"html": func(value any) template.HTML {
			return template.HTML(fmt.Sprint(value))
		},
	}
	t := templates.New(functions)

	if len(*f) > 0 {
		if *p {
			defer profile.Start(profile.ProfilePath(".")).Stop()
			go func() {
				http.ListenAndServe(":8081", nil)
			}()
		}

		w, err := model.Parse(*f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		world = w

		fmt.Println("Hallo Welt!")
		runtime.GC()
		// world.Process()

		// model.ListOtherElements("world", &[]*model.World{&world})
		// model.ListOtherElements("region", &world.Regions)
		// model.ListOtherElements("underground regions", &world.UndergroundRegions)
		// model.ListOtherElements("landmasses", &world.Landmasses)
		// model.ListOtherElements("sites", &world.Sites)
		// model.ListOtherElements("world constructions", &world.WorldConstructions)
		// model.ListOtherElements("artifacts", &world.Artifacts)
		// model.ListOtherElements("entities", &world.Entities)
		// model.ListOtherElements("hf", &world.HistoricalFigures)
		// model.ListOtherElements("events", &world.HistoricalEvents)
		// model.ListOtherElements("collections", &world.HistoricalEventCollections)
		// model.ListOtherElements("era", &world.HistoricalEras)
		// model.ListOtherElements("danceForm", &world.DanceForms)
		// model.ListOtherElements("musicalForm", &world.MusicalForms)
		// model.ListOtherElements("poeticForm", &world.PoeticForms)
		// model.ListOtherElements("written", &world.WrittenContents)

		// server.RegisterResource(router, "region", world.Regions)
		// // server.RegisterResource(router, "undergroundRegion", world.UndergroundRegions)
		// server.RegisterResource(router, "landmass", world.Landmasses)
		// server.RegisterResource(router, "site", world.Sites)
		// server.RegisterResource(router, "worldConstruction", world.WorldConstructions)
		// server.RegisterResource(router, "artifact", world.Artifacts)
		// server.RegisterResource(router, "hf", world.HistoricalFigures)
		// server.RegisterResource(router, "collection", world.HistoricalEventCollections)
		// server.RegisterResource(router, "entity", world.Entities)
		// server.RegisterResource(router, "event", world.HistoricalEvents)
		// // server.RegisterResource(router, "era", world.HistoricalEras)
		// server.RegisterResource(router, "danceForm", world.DanceForms)
		// server.RegisterResource(router, "musicalForm", world.MusicalForms)
		// server.RegisterResource(router, "poeticForm", world.PoeticForms)
		// server.RegisterResource(router, "written", world.WrittenContents)

		RegisterResourcePage(router, "/entity/{id}", t, "entity.html", func(id int) any { return world.Entities[id] })
		RegisterResourcePage(router, "/hf/{id}", t, "hf.html", func(id int) any { return world.HistoricalFigures[id] })
		RegisterResourcePage(router, "/region/{id}", t, "region.html", func(id int) any { return world.Regions[id] })
		RegisterResourcePage(router, "/site/{id}", t, "site.html", func(id int) any { return world.Sites[id] })
		RegisterResourcePage(router, "/artifact/{id}", t, "artifact.html", func(id int) any { return world.Artifacts[id] })
		RegisterPage(router, "/events", t, "eventTypes.html", func(p Parms) any { return allEventTypes() })
		RegisterPage(router, "/events/{type}", t, "eventType.html", func(p Parms) any { return eventsOfType(p["type"]) })
	}

	spa := spaHandler{staticFS: static, staticPath: "static", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	fmt.Println("Serving at :8080")
	http.ListenAndServe(":8080", router)
}

func allEventTypes() []string {
	types := make(map[string]bool)
	for _, e := range world.HistoricalEvents {
		types[e.Details.Type()] = true
	}
	var list = util.Keys(types)
	sort.Strings(list)
	return list
}

func eventsOfType(t string) any {
	var list []*model.HistoricalEvent
	for _, e := range world.HistoricalEvents {
		if e.Details.Type() == t {
			list = append(list, e)
		}
	}

	sort.Slice(list, func(i, j int) bool { return list[i].Id_ < list[j].Id_ })

	return struct {
		Type   string
		Events []*model.HistoricalEvent
	}{
		Type:   t,
		Events: list,
	}
}

type Parms map[string]string

func RegisterPage(router *mux.Router, path string, templates *templates.Template, template string, accessor func(Parms) any) {
	get := func(w http.ResponseWriter, r *http.Request) {
		err := templates.Render(w, template, accessor(mux.Vars(r)))
		if err != nil {
			fmt.Fprintln(w, err)
			fmt.Println(err)
		}
	}

	router.HandleFunc(path, get).Methods("GET")
}

func RegisterResourcePage(router *mux.Router, path string, templates *templates.Template, template string, accessor func(int) any) {
	RegisterPage(router, path, templates, template, func(params Parms) any {
		id, _ := strconv.Atoi(params["id"])
		return accessor(id)
	})
}

type spaHandler struct {
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
	fmt.Println(r.URL, "->", path)

	_, err := h.staticFS.Open(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		fmt.Println(path)
		index, err := h.staticFS.ReadFile(h.staticPath + "/" + h.indexPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusAccepted)
		w.Write(index)
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the subdirectory of the static dir
	statics, err := fs.Sub(h.staticFS, h.staticPath)
	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.FS(statics)).ServeHTTP(w, r)
}
