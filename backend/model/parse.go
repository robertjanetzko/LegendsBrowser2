package model

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func (e *HistoricalEvent) Name() string           { return "" }
func (e *HistoricalEventCollection) Name() string { return "" }

func NewLegendsDecoder(file string) (*xml.Decoder, *os.File, *pb.ProgressBar, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return nil, nil, nil, err
	}
	size := fi.Size()
	bar := pb.Full.Start64(size)

	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened", file)

	converter := util.NewConvertReader(xmlFile)
	barReader := bar.NewProxyReader(converter)
	d := xml.NewDecoder(barReader)

	return d, xmlFile, bar, err
}

func NewLegendsParser(file string) (*util.XMLParser, *os.File, *pb.ProgressBar, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return nil, nil, nil, err
	}
	size := fi.Size()
	bar := pb.Full.Start64(size)

	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened", file)

	barReader := bar.NewProxyReader(xmlFile)
	d := util.NewXMLParser(bufio.NewReader(barReader))

	return d, xmlFile, bar, err
}

func Parse(file string) (*DfWorld, error) {
	InitSameFields()

	p, xmlFile, bar, err := NewLegendsParser(file)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	world := &DfWorld{}

BaseLoop:
	for {
		t, n, err := p.Token()
		if err != nil {
			return nil, err
		}
		switch t {
		case util.StartElement:
			if n == "df_world" {
				world, err = parseDfWorld(p)
				if err != nil {
					return nil, err
				}
				break BaseLoop
			}
		}
	}

	bar.Finish()

	plus := true

	if plus {
		file = strings.Replace(file, "-legends.xml", "-legends_plus.xml", 1)

		p, xmlFile, bar, err = NewLegendsParser(file)
		if err != nil {
			return nil, err
		}
		defer xmlFile.Close()

	PlusLoop:
		for {
			t, n, err := p.Token()
			if err != nil {
				return nil, err
			}
			switch t {
			case util.StartElement:
				if n == "df_world" {
					world, err = parseDfWorldPlus(p, world)
					if err != nil {
						return nil, err
					}
					break PlusLoop
				}
			}
		}

		bar.Finish()
	}

	same, err := json.MarshalIndent(exportSameFields(), "", "  ")
	if err != nil {
		return world, err
	}
	ioutil.WriteFile("same.json", same, 0644)

	return world, nil
}

func parseArray[T any](p *util.XMLParser, dest *[]T, creator func(*util.XMLParser) (T, error)) {
	for {
		t, _, err := p.Token()
		if err != nil {
			return // nil, err
		}
		switch t {
		case util.StartElement:
			x, _ := creator(p)
			*dest = append(*dest, x)

		case util.EndElement:
			return
		}
	}
}

func parseMap[T Identifiable](p *util.XMLParser, dest *map[int]T, creator func(*util.XMLParser) (T, error)) {
	for {
		t, _, err := p.Token()
		if err != nil {
			return // nil, err
		}
		switch t {
		case util.StartElement:
			x, _ := creator(p)
			(*dest)[x.Id()] = x

		case util.EndElement:
			return
		}
	}
}

func parseMapPlus[T Identifiable](p *util.XMLParser, dest *map[int]T, creator func(*util.XMLParser, T) (T, error)) {
	for {
		t, _, err := p.Token()
		if err != nil {
			return
		}
		switch t {
		case util.StartElement:
			id, err := parseId(p)
			if err != nil {
				log.Fatal(err)
			}
			x, err := creator(p, (*dest)[id])
			if err != nil {
				return
			}
			(*dest)[id] = x

		case util.EndElement:
			return
		}
	}
}
func parseId(p *util.XMLParser) (int, error) {
	for {
		t, n, err := p.Token()
		if err != nil {
			return -1, err
		}
		switch t {
		case util.StartElement:
			if n == "id" {
				d, err := p.Value()
				if err != nil {
					return -1, err
				}
				return strconv.Atoi(string(d))
			} else {
				p.Skip()
			}
		}
	}
}

func num(b []byte) int {
	v, _ := strconv.Atoi(string(b))
	return v
}

func txt(b []byte) string {
	return util.ConvertCp473(b)
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
