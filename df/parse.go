package df

import (
	"encoding/xml"
	"fmt"
	"legendsbrowser/util"
	"os"
)

// type DfWorld struct{}

// func (x *DfWorld) Parse(d *xml.Decoder, start *xml.StartElement) {}

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
				w := DfWorld{}
				w.Parse(d, &t)
				return &w, nil
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

func parseArray[T Parsable](d *xml.Decoder, dest *[]T, creator func() T) {
	for {
		tok, err := d.Token()
		if err != nil {
			return // nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			x := creator()
			x.Parse(d, &t)
			*dest = append(*dest, x)

		case xml.EndElement:
			return
		}
	}
}

func parseMap[T IdentifiableParsable](d *xml.Decoder, dest *map[int]T, creator func() T) {
	for {
		tok, err := d.Token()
		if err != nil {
			return // nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			x := creator()
			x.Parse(d, &t)
			(*dest)[x.Id()] = x

		case xml.EndElement:
			return
		}
	}
}
