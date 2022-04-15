package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/pkg/profile"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/server"
)

var world *model.DfWorld

//go:embed resources/frontend
var frontend embed.FS

func main() {
	f := flag.String("f", "", "open a file")
	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)

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

		server.RegisterResource(router, "region", world.Regions)
		// server.RegisterResource(router, "undergroundRegion", world.UndergroundRegions)
		server.RegisterResource(router, "landmass", world.Landmasses)
		server.RegisterResource(router, "site", world.Sites)
		server.RegisterResource(router, "worldConstruction", world.WorldConstructions)
		server.RegisterResource(router, "artifact", world.Artifacts)
		server.RegisterResource(router, "hf", world.HistoricalFigures)
		server.RegisterResource(router, "collection", world.HistoricalEventCollections)
		server.RegisterResource(router, "entity", world.Entities)
		server.RegisterResource(router, "event", world.HistoricalEvents)
		// server.RegisterResource(router, "era", world.HistoricalEras)
		server.RegisterResource(router, "danceForm", world.DanceForms)
		server.RegisterResource(router, "musicalForm", world.MusicalForms)
		server.RegisterResource(router, "poeticForm", world.PoeticForms)
		// server.RegisterResource(router, "written", world.WrittenContents)
	}

	spa := spaHandler{staticFS: frontend, staticPath: "resources/frontend", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	fmt.Println("Serving at :8080")
	http.ListenAndServe(":8080", router)

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
