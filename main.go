package main

import (
	"fmt"
	"legendsbrowser/model"
	"legendsbrowser/server"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/gorilla/mux"
)

var world model.World

func main() {
	// defer profile.Start(profile.MemProfile).Stop()
	// go func() {
	// 	http.ListenAndServe(":8081", nil)
	// }()

	fmt.Println("Hallo Welt!")

	// world.Load("region1-00152-01-01-legends_plus.xml")
	world.Load("region2-00195-01-01-legends.xml")
	// world.Load("Agora-00033-01-01-legends_plus.xml")
	runtime.GC()
	world.Process()

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

	server.RegisterResource(router, "region", world.RegionMap)
	server.RegisterResource(router, "undergroundRegion", world.UndergroundRegionMap)
	server.RegisterResource(router, "landmass", world.LandmassMap)
	server.RegisterResource(router, "site", world.SiteMap)
	server.RegisterResource(router, "worldConstruction", world.WorldConstructionMap)
	server.RegisterResource(router, "artifact", world.ArtifactMap)
	server.RegisterResource(router, "hf", world.HistoricalFigureMap)
	server.RegisterResource(router, "collection", world.HistoricalEventCollectionMap)
	server.RegisterResource(router, "entity", world.EntityMap)
	server.RegisterResource(router, "event", world.HistoricalEventMap)
	server.RegisterResource(router, "era", world.HistoricalEraMap)
	server.RegisterResource(router, "danceForm", world.DanceFormMap)
	server.RegisterResource(router, "musicalForm", world.MusicalFormMap)
	server.RegisterResource(router, "poeticForm", world.PoeticFormMap)
	server.RegisterResource(router, "written", world.WrittenContentMap)

	spa := server.SpaHandler{StaticPath: "frontend/dist/legendsbrowser", IndexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	fmt.Println("Serving at :8080")
	http.ListenAndServe(":8080", router)
}
