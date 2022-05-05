package main

import (
	"embed"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/pkg/profile"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/server"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
)

//go:embed static
var static embed.FS

func main() {
	f := flag.String("f", "", "open a file")
	p := flag.Bool("p", false, "start profiling")
	c := flag.String("c", "", "config file")
	l := flag.Bool("l", false, "open last file")
	flag.Parse()

	if *p {
		defer profile.Start(profile.ProfilePath(".")).Stop()
		go func() {
			http.ListenAndServe(":8081", nil)
		}()
	}

	config, err := server.LoadConfig(*c)
	if err != nil {
		log.Fatal(err)
	}

	templates.DebugTemplates = config.DebugTemplates

	var world *model.DfWorld

	if *l {
		*f = config.LastFile
	}

	if len(*f) > 0 {
		w, err := model.Parse(*f, nil)
		if err != nil {
			log.Fatal(err)
		}

		runtime.GC()
		world = w
	}

	err = server.StartServer(config, world, static)
	if err != nil {
		log.Fatal(err)
	}
}
