package df

import (
	"encoding/xml"
	"fmt"
	"legendsbrowser/util"
	"os"
)

// type DfWorld struct{}

// func parseDfWorld(d *xml.Decoder, start *xml.StartElement) (*DfWorld, error) { return nil, nil }

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
	// return nil, errors.New("Fehler!")
}

type Identifiable interface {
	Id() int
}

type Parsable interface {
	Parse(d *xml.Decoder, start *xml.StartElement) error
}

type IdentifiableParsable interface {
	Identifiable
	Parsable
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
