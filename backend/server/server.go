package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
)

type DfServer struct {
	router    *mux.Router
	templates *templates.Template
	world     *model.DfWorld
}

func StartServer(world *model.DfWorld, static embed.FS) {
	srv := &DfServer{
		router: mux.NewRouter().StrictSlash(true),
		world:  world,
	}
	srv.LoadTemplates()

	srv.RegisterResourcePage("/entity/{id}", "entity.html", func(id int) any { return srv.world.Entities[id] })
	srv.RegisterResourcePage("/hf/{id}", "hf.html", func(id int) any { return srv.world.HistoricalFigures[id] })
	srv.RegisterResourcePage("/region/{id}", "region.html", func(id int) any { return srv.world.Regions[id] })
	srv.RegisterResourcePage("/site/{id}", "site.html", func(id int) any { return srv.world.Sites[id] })
	srv.RegisterResourcePage("/artifact/{id}", "artifact.html", func(id int) any { return srv.world.Artifacts[id] })
	srv.RegisterPage("/events", "eventTypes.html", func(p Parms) any { return srv.world.AllEventTypes() })
	srv.RegisterPage("/events/{type}", "eventType.html", func(p Parms) any { return srv.world.EventsOfType(p["type"]) })

	spa := spaHandler{staticFS: static, staticPath: "static", indexPath: "index.html"}
	srv.router.PathPrefix("/").Handler(spa)

	OpenBrowser("http://localhost:8080")

	fmt.Println("Serving at :8080")
	http.ListenAndServe(":8080", srv.router)
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
