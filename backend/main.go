package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/pkg/profile"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/server"
)

var world *model.DfWorld

func main() {
	f := flag.String("f", "", "open a file")
	flag.Parse()

	if len(*f) > 0 {
		defer profile.Start(profile.MemProfile).Stop()
		go func() {
			http.ListenAndServe(":8081", nil)
		}()

		w, err := model.Parse(*f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		world = w

		fmt.Println("Hallo Welt!")
		runtime.GC()
		// world.Process()

		// model.ListOtherElements("world", &[]*model.World{&world})
		// model.ListOtherElements("region", &world.Regions)
		// model.ListOtherElements("underground regions", &world.UndergroundRegions)
		// model.ListOtherElements("landmasses", &world.Landmasses)
		// model.ListOtherElements("sites", &world.Sites)
		// model.ListOtherElements("world constructions", &world.WorldConstructions)
		// model.ListOtherElements("artifacts", &world.Artifacts)
		// model.ListOtherElements("entities", &world.Entities)
		// model.ListOtherElements("hf", &world.HistoricalFigures)
		// model.ListOtherElements("events", &world.HistoricalEvents)
		// model.ListOtherElements("collections", &world.HistoricalEventCollections)
		// model.ListOtherElements("era", &world.HistoricalEras)
		// model.ListOtherElements("danceForm", &world.DanceForms)
		// model.ListOtherElements("musicalForm", &world.MusicalForms)
		// model.ListOtherElements("poeticForm", &world.PoeticForms)
		// model.ListOtherElements("written", &world.WrittenContents)

		router := mux.NewRouter().StrictSlash(true)

		server.RegisterResource(router, "region", world.Regions)
		// server.RegisterResource(router, "undergroundRegion", world.UndergroundRegions)
		server.RegisterResource(router, "landmass", world.Landmasses)
		server.RegisterResource(router, "site", world.Sites)
		server.RegisterResource(router, "worldConstruction", world.WorldConstructions)
		server.RegisterResource(router, "artifact", world.Artifacts)
		server.RegisterResource(router, "hf", world.HistoricalFigures)
		server.RegisterResource(router, "collection", world.HistoricalEventCollections)
		server.RegisterResource(router, "entity", world.Entities)
		server.RegisterResource(router, "event", world.HistoricalEvents)
		// server.RegisterResource(router, "era", world.HistoricalEras)
		server.RegisterResource(router, "danceForm", world.DanceForms)
		server.RegisterResource(router, "musicalForm", world.MusicalForms)
		server.RegisterResource(router, "poeticForm", world.PoeticForms)
		// server.RegisterResource(router, "written", world.WrittenContents)

		spa := server.SpaHandler{StaticPath: "frontend/dist/legendsbrowser", IndexPath: "index.html"}
		router.PathPrefix("/").Handler(spa)

		fmt.Println("Serving at :8080")
		http.ListenAndServe(":8080", router)
	}

}
