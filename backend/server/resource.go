package server

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Parms map[string]string

func (srv *DfServer) RegisterWorldPage(path string, template string, accessor func(Parms) any) {
	get := func(w http.ResponseWriter, r *http.Request) {
		if srv.context.world == nil {
			srv.renderLoading(w, r)
			return
		}

		data := accessor(mux.Vars(r))
		if data == nil || (reflect.ValueOf(data).Kind() == reflect.Ptr && reflect.ValueOf(data).IsNil()) {
			srv.notFound(w)
			return
		}

		err := srv.templates.Render(w, template, data)
		if err != nil {
			fmt.Fprintln(w, err)
			fmt.Println(err)
		}
	}

	srv.router.HandleFunc(path, get).Methods("GET")
}

func (srv *DfServer) RegisterWorldResourcePage(path string, template string, accessor func(int) any) {
	srv.RegisterWorldPage(path, template, func(params Parms) any {
		id, _ := strconv.Atoi(params["id"])
		return accessor(id)
	})
}
