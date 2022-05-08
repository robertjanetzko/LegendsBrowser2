package model

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	_ "golang.org/x/image/bmp"
)

func (w *DfWorld) LoadMap() {
	w.LoadDimensions()

	path := ""
	files, err := filepath.Glob(strings.ReplaceAll(w.FilePath, "-legends.xml", "-world_map.*"))
	if err == nil && len(files) > 0 {
		path = files[len(files)-1]
	}
	files, err = filepath.Glob(strings.ReplaceAll(w.FilePath, "-legends.xml", "-detailed.*"))
	if err == nil && len(files) > 0 {
		path = files[len(files)-1]
	}

	if path == "" {
		return
	}

	mapImage, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Found Map", path)
	img, format, err := image.Decode(mapImage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("loaded img", format)
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.MapData = buf.Bytes()
	w.MapReady = true
}

func (w *DfWorld) LoadDimensions() {
	files, err := filepath.Glob(filepath.Join(filepath.Dir(w.FilePath), "*-world_gen_param.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	path := ""
	for _, f := range files {
		prefix := filepath.Base(f)[:len(filepath.Base(f))-len("world_gen_param.txt")]
		if strings.HasPrefix(filepath.Base(w.FilePath), prefix) {
			path = f
			break
		}
	}
	if path == "" {
		return
	}

	fmt.Println("Found Worldgen", path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	r := regexp.MustCompile(`\[DIM:(\d+):(\d+)\]`)
	result := r.FindAllStringSubmatch(string(content), 1)
	if result == nil {
		return
	}
	w.Width, _ = strconv.Atoi(result[0][2])
	w.Height, _ = strconv.Atoi(result[0][1])
}

type Coord struct {
	X, Y int
}

func Coords(s string) []Coord {
	var coords []Coord
	for _, c := range strings.Split(s, "|") {
		if c == "" {
			continue
		}
		d := strings.Split(c, ",")
		x, _ := strconv.Atoi(d[0])
		y, _ := strconv.Atoi(d[1])
		coords = append(coords, Coord{X: x, Y: y})
	}
	return coords
}

func maxCoords(coords []Coord) Coord {
	var max Coord
	for _, c := range coords {
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
	}
	return max
}

func (r *Region) Outline() []Coord {
	var outline []Coord

	if r.Coords == "" {
		return outline
	}
	// if (cacheOutline != null)
	// 	return cacheOutline;

	/* draw the region in a matrix */
	coords := Coords(r.Coords)
	max := maxCoords(coords)

	var region = make([][]bool, max.X+3)
	for i := range region {
		region[i] = make([]bool, max.Y+3)
	}
	for _, c := range coords {
		region[c.X+1][c.Y+1] = true
	}

	var curdir, prevdir rune
	curdir = 'n'
	if len(coords) == 1 || coords[0].X == coords[1].X-1 {
		curdir = 'e'
	}

	x0 := coords[0].X + 1
	y0 := coords[0].Y + 1
	x := x0
	y := y0

	/* follow the outline by keeping the right hand inside */
Loop:
	for {
		if !(x != x0 || y != y0 || prevdir == 0) {
			break Loop
		}

		prevdir = curdir
		switch {
		case curdir == 'n' && y > 1:
			y -= 1
			if region[x-1][y-1] {
				curdir = 'w'
			} else if !region[x][y-1] {
				curdir = 'e'
			}
		case curdir == 's' && y < max.Y+2:
			y += 1
			if region[x][y] {
				curdir = 'e'
			} else if !region[x-1][y] {
				curdir = 'w'
			}
		case curdir == 'w' && x > 1:
			x -= 1
			if region[x-1][y] {
				curdir = 's'
			} else if !region[x-1][y-1] {
				curdir = 'n'
			}
		case curdir == 'e' && x < max.X+2:
			x += 1
			if region[x][y-1] {
				curdir = 'n'
			} else if !region[x][y] {
				curdir = 's'
			}
		}
		if curdir != prevdir {
			/* change of direction: record point */
			outline = append(outline, Coord{X: x - 1, Y: y - 1})
			if len(outline) > 256*10 {
				break Loop
			}
		}
	}
	return outline
}

func (x *WorldConstruction) Line() []Coord {
	return Coords(x.Coords)
}
