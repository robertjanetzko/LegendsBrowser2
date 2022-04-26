package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	"github.com/pkg/profile"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/server"
)

//go:embed static
var static embed.FS

func main() {
	f := flag.String("f", "", "open a file")
	p := flag.Bool("p", false, "start profiling")
	flag.Parse()

	if len(*f) > 0 {
		if *p {
			defer profile.Start(profile.ProfilePath(".")).Stop()
			go func() {
				http.ListenAndServe(":8081", nil)
			}()
		}

		w, err := model.Parse(*f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		runtime.GC()

		server.StartServer(w, static)
	}

}
