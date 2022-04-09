package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	sf "github.com/sa-/slicefunk"
)

type World struct {
	XMLName                    xml.Name                    `xml:"df_world"`
	Name                       string                      `xml:"name"`
	AltName                    string                      `xml:"altname"`
	Regions                    []Region                    `xml:"regions>region"`
	UndergroundRegions         []UndergroundRegion         `xml:"underground_regions>underground_region"`
	Landmasses                 []Landmass                  `xml:"landmasses>landmass"`
	Sites                      []Site                      `xml:"sites>site"`
	WorldConstructions         []WorldConstruction         `xml:"world_constructions>world_construction"`
	Artifacts                  []Artifact                  `xml:"artifacts>artifact"`
	HistoricalFigures          []HistoricalFigure          `xml:"historical_figures>historical_figure"`
	HistoricalEvents           []HistoricalEvent           `xml:"historical_events>historical_event"`
	HistoricalEventCollections []HistoricalEventCollection `xml:"historical_event_collections>historical_event_collection"`
}

type NamedObject struct {
	Id   int    `xml:"id" json:"id"`
	Name string `xml:"name" json:"name"`
}

func (r NamedObject) id() int      { return r.Id }
func (r NamedObject) name() string { return r.Name }

type Named interface {
	id() int
	name() string
}

type Region struct {
	XMLName xml.Name `xml:"region" json:"-"`
	NamedObject
	Type string `xml:"type" json:"type"`
}

type UndergroundRegion struct {
	XMLName xml.Name `xml:"underground_region"`
	NamedObject
	Type string `xml:"type" json:"type"`
}

type Landmass struct {
	XMLName xml.Name `xml:"landmass"`
	NamedObject
}

type Site struct {
	XMLName xml.Name `xml:"site" json:"-"`
	NamedObject
	Type       string      `xml:"type" json:"type"`
	Coords     string      `xml:"coords" json:"coords"`
	Rectangle  string      `xml:"rectangle" json:"rectangle"`
	Structures []Structure `xml:"structures>structure" json:"structures"`
}

// func (obj Site) id() int      { return obj.Id }
// func (obj Site) name() string { return obj.Name }

type Structure struct {
	XMLName xml.Name `xml:"structure" json:"-"`
	LocalId int      `xml:"local_id" json:"localId"`
	Name    string   `xml:"name" json:"name"`
	Type    string   `xml:"type" json:"type"`
}

type WorldConstruction struct {
	XMLName xml.Name `xml:"world_construction"`
	NamedObject
}

type Artifact struct {
	XMLName xml.Name `xml:"artifact"`
	NamedObject
	SiteId int `xml:"site_id" json:"siteId"`
}

type Element struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

type HistoricalFigure struct {
	XMLName xml.Name `xml:"historical_figure"`
	NamedObject
	Race  string    `xml:"race" json:"race"`
	Caste string    `xml:"caste" json:"caste"`
	Other []Element `xml:",any" json:"-"`
}

type HistoricalEvent struct {
	XMLName xml.Name `xml:"historical_event"`
	Id      int      `xml:"id"`
	Year    int      `xml:"year"`
	Seconds int      `xml:"seconds72"`
	Type    string   `xml:"type"`
}

type HistoricalEventCollection struct {
	XMLName xml.Name `xml:"historical_event_collection"`
	NamedObject
	StartYear    int    `xml:"year"`
	StartSeconds int    `xml:"seconds72"`
	EndYear      int    `xml:"end_year"`
	EndSeconds   int    `xml:"end_seconds72"`
	Type         string `xml:"type" json:"type"`
	EventIds     []int  `xml:"event" json:"eventIds"`
}

var world World

func main() {
	fmt.Println("Hallo Welt!")

	xmlFile, err := os.Open("region1-00152-01-01-legends.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	fmt.Println(len(byteValue))

	err = xml.Unmarshal(byteValue, &world)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Sites: %d\n", len(world.Sites))
	fmt.Printf("artifacts: %d\n", len(world.Artifacts))
	fmt.Printf("events: %d\n", len(world.HistoricalEvents))
	fmt.Printf("collections: %d\n", len(world.HistoricalEventCollections))
	fmt.Printf("  events: %v\n", len(world.HistoricalEventCollections[0].EventIds))
	fmt.Printf("figures: %d\n", len(world.HistoricalFigures))
	fmt.Printf("  len: %d\n", len(world.HistoricalFigures[0].Other))
	// fmt.Printf("  other: %v\n", world.HistoricalFigures[0].Other)

	// for i := 0; i < len(world.Regions); i++ {
	// 	fmt.Println("Regions Name: " + world.Regions[i].Name)
	// }

	// for i := 0; i < len(world.Sites); i++ {
	// 	fmt.Println("Sites Name: " + world.Sites[i].Name)
	// }

	router := mux.NewRouter().StrictSlash(true)
	registerResource(router, "region", world.Regions)
	registerResource(router, "undergroundRegion", world.UndergroundRegions)
	registerResource(router, "landmass", world.Landmasses)
	registerResource(router, "site", world.Sites)
	registerResource(router, "worldConstruction", world.WorldConstructions)
	registerResource(router, "artifact", world.Artifacts)
	registerResource(router, "hf", world.HistoricalFigures)
	registerResource(router, "collection", world.HistoricalEventCollections)
	http.ListenAndServe(":8080", router)
}

type Info struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func registerResource[T Named](router *mux.Router, resourceName string, resources []T) {

	list := func(w http.ResponseWriter, r *http.Request) {

		values := sf.Map(resources, func(item T) *Info { return &Info{Id: item.id(), Name: item.name()} })

		json.NewEncoder(w).Encode(values)
	}

	get := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			fmt.Println(err)
		}

		for _, item := range resources {
			if item.id() == id {
				json.NewEncoder(w).Encode(item)
			}
		}
	}

	router.HandleFunc(fmt.Sprintf("/%s", resourceName), list).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/%s/{id}", resourceName), get).Methods("GET")
}
