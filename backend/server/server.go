package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
	"github.com/shirou/gopsutil/disk"
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

	srv.RegisterWorldPage("/entities", "entities.html", func(p Parms) any { return grouped(srv.context.world.Entities) })
	srv.RegisterWorldResourcePage("/entity/{id}", "entity.html", func(id int) any { return srv.context.world.Entities[id] })

	srv.RegisterWorldResourcePage("/hf/{id}", "hf.html", func(id int) any { return srv.context.world.HistoricalFigures[id] })
	srv.RegisterWorldResourcePage("/region/{id}", "region.html", func(id int) any { return srv.context.world.Regions[id] })
	srv.RegisterWorldResourcePage("/site/{id}", "site.html", func(id int) any { return srv.context.world.Sites[id] })
	srv.RegisterWorldResourcePage("/artifact/{id}", "artifact.html", func(id int) any { return srv.context.world.Artifacts[id] })

	srv.RegisterWorldPage("/", "eventTypes.html", func(p Parms) any { return srv.context.world.AllEventTypes() })
	srv.RegisterWorldPage("/events", "eventTypes.html", func(p Parms) any { return srv.context.world.AllEventTypes() })
	srv.RegisterWorldPage("/events/{type}", "eventType.html", func(p Parms) any { return srv.context.world.EventsOfType(p["type"]) })

	srv.router.PathPrefix("/load").Handler(srv.loader)

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

type loadProgress struct {
	Msg      string  `json:"msg"`
	Progress float64 `json:"progress"`
	Done     bool    `json:"done"`
}

func (h loadHandler) Progress() *loadProgress {
	percent := 0.0
	p := h.server.context.progress
	if p.ProgressBar != nil {
		percent = float64(p.ProgressBar.Current()*100) / float64(p.ProgressBar.Total())
	}

	return &loadProgress{
		Msg:      h.server.context.progress.Message,
		Progress: percent,
		Done:     h.server.context.world != nil,
	}
}

func (h loadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/load/progress" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(h.Progress())
		return
	}

	var partitions []string
	if runtime.GOOS == "windows" {
		ps, _ := disk.Partitions(false)
		partitions = util.Map(ps, func(p disk.PartitionStat) string { return p.Mountpoint + `\` })
	} else {
		partitions = append(partitions, "/")
	}

	path := r.URL.Query().Get("p")

	p := &paths{
		Partitions: partitions,
		Current:    path,
	}
	if p.Current == "" {
		p.Current = "."
	}
	var err error
	p.Current, err = filepath.Abs(p.Current)
	if err != nil {
		httpError(w, err)
		return
	}

	if f, err := os.Stat(p.Current); err == nil {
		if f.IsDir() {
			p.List, err = ioutil.ReadDir(p.Current)
			if err != nil {
				httpError(w, err)
				return
			}

			err = h.server.templates.Render(w, "load.html", p)
			if err != nil {
				httpError(w, err)
			}
			return
		} else {
			h.server.context.isLoading = true
			h.server.context.world = nil
			go loadWorld(h.server, p.Current)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "/load", http.StatusSeeOther)
}

func isLegendsXml(f fs.FileInfo) bool {
	return strings.HasSuffix(f.Name(), "-legends.xml")
}

func loadWorld(server *DfServer, file string) {
	runtime.GC()
	wrld, _ := model.Parse(file, server.context.progress)
	server.context.world = wrld
	server.context.isLoading = false
}

type paths struct {
	Current    string
	List       []fs.FileInfo
	Partitions []string
}

func (srv *DfServer) renderLoading(w http.ResponseWriter, r *http.Request) {
	if srv.context.isLoading {
		err := srv.templates.Render(w, "loading.html", srv.loader.Progress())
		if err != nil {
			httpError(w, err)
		}
	} else {
		http.Redirect(w, r, "/load", http.StatusSeeOther)
	}
}

func httpError(w http.ResponseWriter, err error) {
	fmt.Fprintln(w, err)
	fmt.Println(err)
}

func grouped[T model.Typed](input map[int]T) map[string][]T {
	output := make(map[string][]T)

	for _, v := range input {
		k := v.Type()
		output[k] = append(output[k], v)
	}

	return output
}
