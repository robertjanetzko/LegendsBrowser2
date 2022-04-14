package main

import (
	"flag"

	"github.com/robertjanetzko/LegendsBrowser2/analyze/df"
)

func main() {
	a := flag.String("a", "", "analyze a file")
	g := flag.Bool("g", false, "generate model")
	flag.Parse()

	if len(*a) > 0 {
		df.Analyze(*a)
	}

	if *g {
		df.Generate()
	}
}
