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

	text := string(content)
	fmt.Println(text)

	r := regexp.MustCompile(`\[DIM:(\d+):(\d+)\]`)
	result := r.FindAllStringSubmatch(text, 1)
	if result == nil {
		return
	}
	w.Width, _ = strconv.Atoi(result[0][2])
	w.Height, _ = strconv.Atoi(result[0][1])
}
