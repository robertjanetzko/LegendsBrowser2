package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func (e *HistoricalEvent) Name() string           { return "" }
func (e *HistoricalEventCollection) Name() string { return "" }

func Parse(file string) (*DfWorld, error) {
	InitSameFields()

	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened", file)
	defer xmlFile.Close()

	converter := util.NewConvertReader(xmlFile)
	d := xml.NewDecoder(converter)

	var world *DfWorld
BaseLoop:
	for {
		tok, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == "df_world" {
				world, err = parseDfWorld(d, &t)
				if err != nil {
					return nil, err
				}
				break BaseLoop
			}
		}
	}

	plus := true

	if plus {
		file = strings.Replace(file, "-legends.xml", "-legends_plus.xml", 1)
		xmlFile, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return world, nil
		}

		fmt.Println("Successfully Opened", file)
		defer xmlFile.Close()

		converter := util.NewConvertReader(xmlFile)
		d := xml.NewDecoder(converter)

	PlusLoop:
		for {
			tok, err := d.Token()
			if err != nil {
				return nil, err
			}
			switch t := tok.(type) {
			case xml.StartElement:
				if t.Name.Local == "df_world" {
					world, err = parseDfWorldPlus(d, &t, world)
					if err != nil {
						return nil, err
					}
					break PlusLoop
				}
			}
		}
	}

	same, err := json.MarshalIndent(exportSameFields(), "", "  ")
	if err != nil {
		return world, err
	}
	ioutil.WriteFile("same.json", same, 0644)

	return world, nil
}

func parseArray[T any](d *xml.Decoder, dest *[]T, creator func(*xml.Decoder, *xml.StartElement) (T, error)) {
	for {
		tok, err := d.Token()
		if err != nil {
			return // nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			x, _ := creator(d, &t)
			*dest = append(*dest, x)

		case xml.EndElement:
			return
		}
	}
}

func parseMap[T Identifiable](d *xml.Decoder, dest *map[int]T, creator func(*xml.Decoder, *xml.StartElement) (T, error)) {
	for {
		tok, err := d.Token()
		if err != nil {
			return // nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			x, _ := creator(d, &t)
			(*dest)[x.Id()] = x

		case xml.EndElement:
			return
		}
	}
}

func parseMapPlus[T Identifiable](d *xml.Decoder, dest *map[int]T, creator func(*xml.Decoder, *xml.StartElement, T) (T, error)) {
	for {
		tok, err := d.Token()
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			id, err := parseId(d)
			if err != nil {
				log.Fatal(err)
			}
			x, err := creator(d, &t, (*dest)[id])
			if err != nil {
				return
			}
			(*dest)[id] = x

		case xml.EndElement:
			return
		}
	}
}
func parseId(d *xml.Decoder) (int, error) {
	var data []byte
	for {
		tok, err := d.Token()
		if err != nil {
			return -1, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			data = nil
			if t.Name.Local != "id" {
				d.Skip()
				// return -1, fmt.Errorf("expected id at: %d", d.InputOffset())
			}

		case xml.CharData:
			data = append(data, t...)

		case xml.EndElement:
			if t.Name.Local == "id" {
				return strconv.Atoi(string(data))
			}
		}
	}
}

var sameFields map[string]map[string]map[string]bool

func exportSameFields() map[string]map[string]string {
	export := make(map[string]map[string]string)

	for objectType, v := range sameFields {
		fields := make(map[string]string)
		for field, v2 := range v {
			c := 0
			f := ""
			for field2, same := range v2 {
				if same {
					c++
					f = field2
				}
			}
			if c == 1 {
				fields[field] = f
			}
		}
		if len(fields) > 0 {
			export[objectType] = fields
		}
	}

	return export
}
