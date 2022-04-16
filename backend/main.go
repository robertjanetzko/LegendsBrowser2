package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/profile"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
)

var world *model.DfWorld

//go:embed resources/frontend
var frontend embed.FS

func main() {
	f := flag.String("f", "", "open a file")
	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)

	functions := template.FuncMap{
		"json": func(obj any) string {
			b, err := json.MarshalIndent(obj, "", "  ")
			if err != nil {
				fmt.Println(err)
				return ""
			}
			return string(b)
		},
		"check": func(condition bool, v any) any {
			if condition {
				return v
			}
			return nil
		},
		"title": func(input string) string {
			words := strings.Split(input, " ")
			smallwords := " a an on the to of "

			for index, word := range words {
				if strings.Contains(smallwords, " "+word+" ") && index > 0 {
					words[index] = word
				} else {
					words[index] = strings.Title(word)
				}
			}
			return strings.Join(words, " ")
		},
		"getHf":     func(id int) *model.HistoricalFigure { return world.HistoricalFigures[id] },
		"getEntity": func(id int) *model.Entity { return world.Entities[id] },
		"events": func(obj model.Identifiable) []*model.HistoricalEvent {
			id := obj.Id()
			var list []*model.HistoricalEvent
			switch obj.(type) {
			case *model.HistoricalFigure:
				for _, e := range world.HistoricalEvents {
					if e.Details.RelatedToHf(id) {
						list = append(list, e)
					}
				}
			default:
				fmt.Printf("unknown type %T\n", obj)
			}
			return list
		},
	}
	t := templates.New(functions)

	if len(*f) > 0 {
		defer profile.Start(profile.MemProfile).Stop()
		go func() {
			http.ListenAndServe(":8081", nil)
		}()

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

		RegisterPage(router, "/entity/{id}", t, "entity.html", func(id int) any { return world.Entities[id] })
		RegisterPage(router, "/hf/{id}", t, "hf.html", func(id int) any { return world.HistoricalFigures[id] })

	}

	spa := spaHandler{staticFS: frontend, staticPath: "resources/frontend", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	fmt.Println("Serving at :8080")
	http.ListenAndServe(":8080", router)

}

func RegisterPage(router *mux.Router, path string, templates *templates.Template, template string, accessor func(int) any) {
	get := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			fmt.Fprintln(w, err)
			fmt.Println(err)
		}
		fmt.Println("render", template, id)
		err = templates.Render(w, template, accessor(id))
		if err != nil {
			fmt.Fprintln(w, err)
			fmt.Println(err)
		}
	}

	router.HandleFunc(path, get).Methods("GET")
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
	fmt.Println(path)

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
