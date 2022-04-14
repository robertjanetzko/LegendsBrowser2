package model

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func (e *HistoricalEvent) Name() string           { return "" }
func (e *HistoricalEventCollection) Name() string { return "" }

func Parse(file string) (*DfWorld, error) {
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened", file)
	defer xmlFile.Close()

	converter := util.NewConvertReader(xmlFile)
	d := xml.NewDecoder(converter)

	for {
		tok, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == "df_world" {
				return parseDfWorld(d, &t)
			}
		}
	}
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
