package model

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type World struct {
	XMLName xml.Name `xml:"df_world"`
	Name    string   `xml:"name"`
	AltName string   `xml:"altname"`

	Regions                    []*Region                    `xml:"regions>region"`
	UndergroundRegions         []*UndergroundRegion         `xml:"underground_regions>underground_region"`
	Landmasses                 []*Landmass                  `xml:"landmasses>landmass"`
	Sites                      []*Site                      `xml:"sites>site"`
	WorldConstructions         []*WorldConstruction         `xml:"world_constructions>world_construction"`
	Artifacts                  []*Artifact                  `xml:"artifacts>artifact"`
	HistoricalFigures          []*HistoricalFigure          `xml:"historical_figures>historical_figure"`
	HistoricalEvents           []*HistoricalEvent           `xml:"historical_events>historical_event"`
	HistoricalEventCollections []*HistoricalEventCollection `xml:"historical_event_collections>historical_event_collection"`
	Entities                   []*Entity                    `xml:"entities>entity"`

	RegionMap                    map[int]*Region
	UndergroundRegionMap         map[int]*UndergroundRegion
	LandmassMap                  map[int]*Landmass
	SiteMap                      map[int]*Site
	WorldConstructionMap         map[int]*WorldConstruction
	ArtifactMap                  map[int]*Artifact
	HistoricalFigureMap          map[int]*HistoricalFigure
	HistoricalEventMap           map[int]*HistoricalEvent
	HistoricalEventCollectionMap map[int]*HistoricalEventCollection
	EntityMap                    map[int]*Entity
}

func (w *World) Load(file string) {
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	fmt.Println(len(byteValue))

	err = xml.Unmarshal(byteValue, w)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("World loaded")
}

func (w *World) Process() {
	w.RegionMap = make(map[int]*Region)
	mapObjects(&w.Regions, &w.RegionMap)

	w.UndergroundRegionMap = make(map[int]*UndergroundRegion)
	mapObjects(&w.UndergroundRegions, &w.UndergroundRegionMap)

	w.LandmassMap = make(map[int]*Landmass)
	mapObjects(&w.Landmasses, &w.LandmassMap)

	w.SiteMap = make(map[int]*Site)
	mapObjects(&w.Sites, &w.SiteMap)

	w.WorldConstructionMap = make(map[int]*WorldConstruction)
	mapObjects(&w.WorldConstructions, &w.WorldConstructionMap)

	w.ArtifactMap = make(map[int]*Artifact)
	mapObjects(&w.Artifacts, &w.ArtifactMap)

	w.HistoricalFigureMap = make(map[int]*HistoricalFigure)
	mapObjects(&w.HistoricalFigures, &w.HistoricalFigureMap)

	w.HistoricalEventMap = make(map[int]*HistoricalEvent)
	mapObjects(&w.HistoricalEvents, &w.HistoricalEventMap)

	w.HistoricalEventCollectionMap = make(map[int]*HistoricalEventCollection)
	mapObjects(&w.HistoricalEventCollections, &w.HistoricalEventCollectionMap)

	w.EntityMap = make(map[int]*Entity)
	mapObjects(&w.Entities, &w.EntityMap)

	w.processEvents()
}

func (w *World) processEvents() {
	legendFields := make(map[string][]int)

	t := reflect.TypeOf(HistoricalEvent{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		l, ok := f.Tag.Lookup("legend")
		if ok {
			legendFields[l] = append(legendFields[l], i)
		}
	}

	for eventIndex := 0; eventIndex < len(w.HistoricalEvents); eventIndex++ {
		e := w.HistoricalEvents[eventIndex]
		v := reflect.ValueOf(*e)
		processEvent(e, &v, legendFields["entity"], &w.EntityMap)
		processEvent(e, &v, legendFields["site"], &w.SiteMap)
		processEvent(e, &v, legendFields["hf"], &w.HistoricalFigureMap)
		processEvent(e, &v, legendFields["artifact"], &w.ArtifactMap)
		// processEvent(e, &v, legendFields["wc"], &w.WorldConstructionMap)
		// processEvent(e, &v, legendFields["structure"], &w.St)
	}
}

func processEvent[T HasEvents](event *HistoricalEvent, v *reflect.Value, fields []int, objectMap *map[int]T) {
	for _, i := range fields {
		val := v.Field(i)
		if !val.IsZero() {
			switch val.Elem().Kind() {
			case reflect.Slice:
				ids := val.Interface().(*[]int)
				for _, id := range *ids {
					x, ok := (*objectMap)[id]
					if ok {
						x.SetEvents(append(x.GetEvents(), event))
					}
				}
			case reflect.Int:
				id := int(val.Elem().Int())
				x, ok := (*objectMap)[id]
				if ok {
					x.SetEvents(append(x.GetEvents(), event))
				}
			default:
				fmt.Println("unknown", val.Elem().Kind())
			}
		}
	}
}

func mapObjects[T Identifiable](objects *[]T, objectMap *map[int]T) {
	for i, obj := range *objects {
		(*objectMap)[obj.Id()] = (*objects)[i]
	}
}
