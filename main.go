package main

import (
	"fmt"
	"legendsbrowser/model"
	"legendsbrowser/server"
	"net/http"

	"github.com/gorilla/mux"
)

var world model.World

func main() {
	fmt.Println("Hallo Welt!")

	world.Load("region1-00152-01-01-legends.xml")
	world.Process()

	model.ListOtherElements(&world.HistoricalEvents)
	// listOtherElements(&world.HistoricalFigures)

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

	spa := server.SpaHandler{StaticPath: "frontend/dist/legendsbrowser", IndexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	fmt.Println("Serving at :8080")
	http.ListenAndServe(":8080", router)
}
