package server

import (
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type spaHandler struct {
	server     *DfServer
	staticFS   fs.FS
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path := r.URL.Path
	path = strings.TrimPrefix(path, h.server.context.config.SubUri)
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
		file, err := h.staticFS.Open(h.staticPath + "/" + h.indexPath)
		if err != nil {
			h.server.notFound(w)
			return
		}
		index, err := ioutil.ReadAll(file)
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
	http.StripPrefix(h.server.context.config.SubUri, http.FileServer(http.FS(statics))).ServeHTTP(w, r)
}
