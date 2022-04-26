package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Parms map[string]string

func (srv *DfServer) RegisterPage(path string, template string, accessor func(Parms) any) {
	get := func(w http.ResponseWriter, r *http.Request) {
		err := srv.templates.Render(w, template, accessor(mux.Vars(r)))
		if err != nil {
			fmt.Fprintln(w, err)
			fmt.Println(err)
		}
	}

	srv.router.HandleFunc(path, get).Methods("GET")
}

func (srv *DfServer) RegisterResourcePage(path string, template string, accessor func(int) any) {
	srv.RegisterPage(path, template, func(params Parms) any {
		id, _ := strconv.Atoi(params["id"])
		return accessor(id)
	})
}
