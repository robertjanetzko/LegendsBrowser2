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
	d := flag.Bool("d", false, "debug templates")
	flag.Parse()

	if *p {
		defer profile.Start(profile.ProfilePath(".")).Stop()
		go func() {
			http.ListenAndServe(":8081", nil)
		}()
	}

	templates.DebugTemplates = *d

	var world *model.DfWorld

	if len(*f) > 0 {
		w, err := model.Parse(*f, nil)
		if err != nil {
			log.Fatal(err)
		}

		runtime.GC()
		world = w
	}

	err := server.StartServer(world, static)
	if err != nil {
		log.Fatal(err)
	}
}
