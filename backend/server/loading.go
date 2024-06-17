package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
	"github.com/shirou/gopsutil/disk"
)

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
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(h.Progress())
		return
	}

	if h.server.context.config.ServerMode {
		err := h.server.templates.Render(w, "serverMode.html", nil)
		if err != nil {
			httpError(w, err)
		}
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
		p.Current = h.server.context.config.LastPath
	}
	var err error
	p.Current, err = filepath.Abs(p.Current)
	if err != nil {
		httpError(w, err)
		return
	}

	if f, err := os.Stat(p.Current); err == nil {
		if f.IsDir() {
			h.server.context.config.LastPath = p.Current
			h.server.context.config.Save()

			p.List, err = os.ReadDir(p.Current)
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
			h.server.context.config.LastFile = p.Current
			h.server.context.config.Save()

			h.server.context.isLoading = true
			h.server.context.world = nil
			go loadWorld(h.server, p.Current)
			http.Redirect(w, r, h.server.context.config.SubUri+"/", http.StatusSeeOther)
			return
		}
	}
	if err != nil {
		if !os.IsNotExist(err) {
			httpError(w, err)
		}
	}
	http.Redirect(w, r, h.server.context.config.SubUri+"/load?p=%2f", http.StatusSeeOther)
}

func isLegendsXml(f os.DirEntry) bool {
	return strings.HasSuffix(f.Name(), "-legends.xml")
}

func isLegendsXmlWithPlus(path string, f os.DirEntry) bool {
	plusName := strings.ReplaceAll(f.Name(), "-legends.xml", "-legends_plus.xml")
	_, err := os.Stat(filepath.Join(path, plusName))
	return err == nil
}

func loadWorld(server *DfServer, file string) {
	runtime.GC()
	wrld, _ := model.Parse(file, server.context.progress)
	server.context.world = wrld
	server.context.isLoading = false
}

type paths struct {
	Current    string
	List       []os.DirEntry
	Partitions []string
}

func (srv *DfServer) renderLoading(w http.ResponseWriter, r *http.Request) {
	if srv.context.isLoading {
		err := srv.templates.Render(w, "loading.html", srv.loader.Progress())
		if err != nil {
			httpError(w, err)
		}
	} else {
		http.Redirect(w, r, srv.context.config.SubUri+"/load", http.StatusSeeOther)
	}
}
