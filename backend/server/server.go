package server

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
)

type DfServerContext struct {
	world     *model.DfWorld
	isLoading bool
}

type DfServer struct {
	router    *mux.Router
	templates *templates.Template
	context   *DfServerContext
}

func StartServer(world *model.DfWorld, static embed.FS) {
	srv := &DfServer{
		router: mux.NewRouter().StrictSlash(true),
		context: &DfServerContext{
			world:     world,
			isLoading: false,
		},
	}
	srv.LoadTemplates()

	srv.RegisterWorldResourcePage("/entity/{id}", "entity.html", func(id int) any { return srv.context.world.Entities[id] })
	srv.RegisterWorldResourcePage("/hf/{id}", "hf.html", func(id int) any { return srv.context.world.HistoricalFigures[id] })
	srv.RegisterWorldResourcePage("/region/{id}", "region.html", func(id int) any { return srv.context.world.Regions[id] })
	srv.RegisterWorldResourcePage("/site/{id}", "site.html", func(id int) any { return srv.context.world.Sites[id] })
	srv.RegisterWorldResourcePage("/artifact/{id}", "artifact.html", func(id int) any { return srv.context.world.Artifacts[id] })

	srv.RegisterWorldPage("/", "eventTypes.html", func(p Parms) any { return srv.context.world.AllEventTypes() })
	srv.RegisterWorldPage("/events", "eventTypes.html", func(p Parms) any { return srv.context.world.AllEventTypes() })
	srv.RegisterWorldPage("/events/{type}", "eventType.html", func(p Parms) any { return srv.context.world.EventsOfType(p["type"]) })

	srv.router.PathPrefix("/load").Handler(loadHandler{server: srv})

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

type loadHandler struct {
	server *DfServer
}

func (h loadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("p")

	p := &paths{
		Current: path,
	}
	if p.Current == "" {
		p.Current = "."
	}

	if f, err := os.Stat(p.Current); err == nil {
		if f.IsDir() {
			p.List, err = ioutil.ReadDir(p.Current)
			if err != nil {
				fmt.Fprintln(w, err)
				fmt.Println(err)
			}

			err = h.server.templates.Render(w, "load.html", p)
			if err != nil {
				fmt.Fprintln(w, err)
				fmt.Println(err)
			}
			return
		} else {
			h.server.context.isLoading = true
			wrld, _ := model.Parse(p.Current)
			h.server.context.world = wrld
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "/load", http.StatusSeeOther)
}

func isLegendsXml(f fs.FileInfo) bool {
	return strings.HasSuffix(f.Name(), "-legends.xml")
}
