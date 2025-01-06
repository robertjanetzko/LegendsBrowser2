package model

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

func (e *HistoricalEvent) Name() string { return "" }

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

	fmt.Println("Loading:", file)

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

	fmt.Println("\nLoading:", file)

	barReader := bar.NewProxyReader(xmlFile)
	d := util.NewXMLParser(bufio.NewReader(barReader))

	return d, xmlFile, bar, err
}

type LoadProgress struct {
	Message     string
	ProgressBar *pb.ProgressBar
}

func Parse(file string, lp *LoadProgress) (*DfWorld, error) {
	InitSameFields()

	p, xmlFile, bar, err := NewLegendsParser(file)
	if lp != nil {
		lp.Message = "Loading " + file
		lp.ProgressBar = bar
	}
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	var world *DfWorld

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
	world.FilePath = file

	bar.Finish()

	plus := false
	file = strings.Replace(file, "-legends.xml", "-legends_plus.xml", 1)
	if _, err := os.Stat(file); err == nil {
		plus = true
	} else {
		fmt.Println("\nno legends_plus.xml found")
	}

	if plus {
		p, xmlFile, bar, err = NewLegendsParser(file)
		if lp != nil {
			lp.Message = "Loading " + file
			lp.ProgressBar = bar
		}
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
		world.Plus = true
		world.PlusFilePath = file

		bar.Finish()
	}

	// same, err := json.MarshalIndent(exportSameFields(), "", "  ")
	// if err != nil {
	// 	return world, err
	// }
	// ioutil.WriteFile("same.json", same, 0644)

	world.LoadMap()
	world.LoadHistory()

	world.process()

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

func parseArrayPlus[T any](p *util.XMLParser, dest *[]T, creator func(*util.XMLParser, T) (T, error)) {
	for {
		t, _, err := p.Token()
		if err != nil {
			return // nil, err
		}
		switch t {
		case util.StartElement:
			var x T
			x, _ = creator(p, x)
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
			x.setId(id)
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
