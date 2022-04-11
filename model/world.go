package model

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type World struct {
	XMLName xml.Name `xml:"df_world"`
	Name    string   `xml:"name"`
	AltName string   `xml:"altname"`

	// Regions                    []*Region                    `xml:"regions>region"`
	// UndergroundRegions         []*UndergroundRegion         `xml:"underground_regions>underground_region"`
	// Landmasses                 []*Landmass                  `xml:"landmasses>landmass"`
	// Sites                      []*Site                      `xml:"sites>site"`
	// WorldConstructions         []*WorldConstruction         `xml:"world_constructions>world_construction"`
	// Artifacts                  []*Artifact                  `xml:"artifacts>artifact"`
	// HistoricalFigures          []*HistoricalFigure          //`xml:"historical_figures>historical_figure"`
	// HistoricalEvents           []*HistoricalEvent           //`xml:"historical_events>historical_event"`
	// HistoricalEventCollections []*HistoricalEventCollection `xml:"historical_event_collections>historical_event_collection"`
	// HistoricalEras             []*HistoricalEra             `xml:"historical_eras>historical_era"`
	// Entities                   []*Entity                    `xml:"entities>entity"`
	// EntityPopulations          []*EntityPopulation          `xml:"entity_populations>entity_population"`
	// DanceForms                 []*DanceForm                 `xml:"dance_forms>dance_form"`
	// MusicalForms               []*MusicalForm               `xml:"musical_forms>musical_form"`
	// PoeticForms                []*PoeticForm                `xml:"poetic_forms>poetic_form"`
	// WrittenContents            []*WrittenContent            `xml:"written_contents>written_content"`
	OtherElements

	RegionMap                    map[int]*Region                    `xml:"regions>region"`
	UndergroundRegionMap         map[int]*UndergroundRegion         `xml:"underground_regions>underground_region"`
	LandmassMap                  map[int]*Landmass                  `xml:"landmasses>landmass"`
	SiteMap                      map[int]*Site                      `xml:"sites>site"`
	WorldConstructionMap         map[int]*WorldConstruction         `xml:"world_constructions>world_construction"`
	ArtifactMap                  map[int]*Artifact                  `xml:"artifacts>artifact"`
	HistoricalFigureMap          map[int]*HistoricalFigure          `xml:"historical_figures>historical_figure"`
	HistoricalEventMap           map[int]*HistoricalEvent           `xml:"historical_events>historical_event"`
	HistoricalEventCollectionMap map[int]*HistoricalEventCollection `xml:"historical_event_collections>historical_event_collection"`
	HistoricalEraMap             map[int]*HistoricalEra             `xml:"historical_eras>historical_era"`
	EntityMap                    map[int]*Entity                    `xml:"entities>entity"`
	DanceFormMap                 map[int]*DanceForm                 `xml:"dance_forms>dance_form"`
	MusicalFormMap               map[int]*MusicalForm               //`xml:"musical_forms>musical_form"`
	PoeticFormMap                map[int]*PoeticForm                // `xml:"poetic_forms>poetic_form"`
	WrittenContentMap            map[int]*WrittenContent            `xml:"written_contents>written_content"`
}

var cp437 = []byte("         \t\n  \r                   !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~ CueaaaaceeeiiiAAEaAooouuyOU    faiounN                                                                                                ")

// var cp437 = []byte("         \t\n  \r                   !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~ ÇüéâäàåçêëèïîìÄÅÉæÆôöòûùÿÖÜ¢£¥₧ƒáíóúñÑ                                                                                                ")

type ConvertReader struct {
	r    io.Reader
	read int
}

func (c *ConvertReader) Read(b []byte) (n int, err error) {
	n, err = c.r.Read(b)
	if c.read == 0 && n > 35 {
		copy(b[30:35], []byte("UTF-8"))
	}
	c.read += n
	for i := range b {
		b[i] = cp437[b[i]]
	}
	return n, err
}

func (w *World) Load(file string) {
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	defer xmlFile.Close()

	converter := &ConvertReader{r: xmlFile}

	// byteValue, _ := io.ReadAll(converter)
	// fmt.Println(len(byteValue))

	fillTypes(reflect.TypeOf(w))
	fmt.Println(types["Region"])

	d := xml.NewDecoder(converter)
	parseObject(d, nil, reflect.ValueOf(w))

	// err = xml.Unmarshal(byteValue, w)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println("World loaded")
}

var types = make(map[string]map[string]reflect.StructField)

func fillTypes(t reflect.Type) {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	fmt.Println(t.Name())

	if _, ok := types[t.Name()]; ok {
		return
	}

	info := make(map[string]reflect.StructField)
	DeepFields(t, &info, make([]int, 0))

	types[t.Name()] = info
}

func DeepFields(t reflect.Type, info *map[string]reflect.StructField, index []int) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		f.Index = append(index, f.Index[0])
		if xml, ok := f.Tag.Lookup("xml"); ok {
			if p := strings.Index(xml, ">"); p >= 0 {
				(*info)[xml[0:p]] = f
			} else {
				for _, s := range strings.Split(xml, "|") {
					(*info)[s] = f
				}
			}
			if f.Type.Kind() == reflect.Map || f.Type.Kind() == reflect.Slice {
				fillTypes(f.Type.Elem())
			}
			fmt.Println(i, f)
		}
		if f.Type.Kind() == reflect.Struct && f.Anonymous {
			DeepFields(f.Type, info, f.Index)
		}
	}
}

func parseObject(d *xml.Decoder, start *xml.StartElement, val reflect.Value) error {
	if start == nil {
		for {
			tok, err := d.Token()
			if err != nil {
				return err
			}
			if t, ok := tok.(xml.StartElement); ok {
				start = &t
				break
			}
		}
	}

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	typ, ok := types[val.Type().Name()]
	if !ok {
		d.Skip()
		return nil
	}

Loop:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if ty, ok := typ[t.Name.Local]; ok {
				if ty.Type.Kind() == reflect.Map {
					fmt.Println("   ", t.Name.Local, val.Type().Name(), ty)
					f := val.Field(ty.Index[0])
					if f.IsNil() {
						f.Set(reflect.MakeMapWithSize(ty.Type, 0))
					}
					parseMap(d, ty, f)
					val.Field(ty.Index[0]).SetMapIndex(reflect.ValueOf(6), reflect.New(ty.Type.Elem().Elem()))
				}
			} else {
				d.Skip()
			}
			// parseObject(d, &t, val)
		case xml.EndElement:
			break Loop
		}
	}
	return nil
}

func parseMap(d *xml.Decoder, field reflect.StructField, dest reflect.Value) error {
	x, ok := field.Tag.Lookup("xml")
	if !ok {
		return errors.New("no xml tag")
	}
	elementName := strings.Split(x, ">")[1]

	var lastStart *xml.StartElement
	var id int

Loop:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == elementName {
				lastStart = &t
				id = -1
			} else if t.Name.Local == "id" {
				if id != -1 {
					return errors.New("ID at invalid place")
				}
				d.DecodeElement(&id, &t)

				obj := dest.MapIndex(reflect.ValueOf(id))
				if !obj.IsValid() {
					obj = reflect.New(field.Type.Elem().Elem())
					dest.SetMapIndex(reflect.ValueOf(id), obj)
					obj.Elem().FieldByIndex(types[obj.Type().Elem().Name()]["id"].Index).SetInt(int64(id))
				}
				d.DecodeElement(obj.Interface(), lastStart)
			} else {
				fmt.Println("SKIP", elementName, t.Name.Local)
				d.Skip()
			}
		case xml.EndElement:
			if t.Name.Local != elementName {
				break Loop
			}
		}
	}
	return nil
}

func (w *World) Process() {
	// w.RegionMap = make(map[int]*Region)
	// mapObjects(&w.Regions, &w.RegionMap)

	// w.UndergroundRegionMap = make(map[int]*UndergroundRegion)
	// mapObjects(&w.UndergroundRegions, &w.UndergroundRegionMap)

	// w.LandmassMap = make(map[int]*Landmass)
	// mapObjects(&w.Landmasses, &w.LandmassMap)

	// w.SiteMap = make(map[int]*Site)
	// mapObjects(&w.Sites, &w.SiteMap)

	// w.WorldConstructionMap = make(map[int]*WorldConstruction)
	// mapObjects(&w.WorldConstructions, &w.WorldConstructionMap)

	// w.ArtifactMap = make(map[int]*Artifact)
	// mapObjects(&w.Artifacts, &w.ArtifactMap)

	// w.HistoricalFigureMap = make(map[int]*HistoricalFigure)
	// mapObjects(&w.HistoricalFigures, &w.HistoricalFigureMap)

	// w.HistoricalEventMap = make(map[int]*HistoricalEvent)
	// mapObjects(&w.HistoricalEvents, &w.HistoricalEventMap)

	// w.HistoricalEventCollectionMap = make(map[int]*HistoricalEventCollection)
	// mapObjects(&w.HistoricalEventCollections, &w.HistoricalEventCollectionMap)

	// w.EntityMap = make(map[int]*Entity)
	// mapObjects(&w.Entities, &w.EntityMap)

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

	// for eventIndex := 0; eventIndex < len(w.HistoricalEvents); eventIndex++ {
	// 	e := w.HistoricalEvents[eventIndex]
	// 	v := reflect.ValueOf(*e)
	// 	processEvent(e, &v, legendFields["entity"], &w.EntityMap)
	// 	processEvent(e, &v, legendFields["site"], &w.SiteMap)
	// 	processEvent(e, &v, legendFields["hf"], &w.HistoricalFigureMap)
	// 	processEvent(e, &v, legendFields["artifact"], &w.ArtifactMap)
	// 	// processEvent(e, &v, legendFields["wc"], &w.WorldConstructionMap)
	// 	// processEvent(e, &v, legendFields["structure"], &w.St)
	// }
}

func processEvent[T HasEvents](event *HistoricalEvent, v *reflect.Value, fields []int, objectMap *map[int]T) {
	for _, i := range fields {
		val := v.Field(i)
		if !val.IsZero() {
			switch val.Elem().Kind() {
			case reflect.Slice:
				ids := val.Interface().(*[]int)
				for _, id := range *ids {
					if x, ok := (*objectMap)[id]; ok {
						x.SetEvents(append(x.GetEvents(), event))
					}
				}
			case reflect.Int:
				id := int(val.Elem().Int())
				if x, ok := (*objectMap)[id]; ok {
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
