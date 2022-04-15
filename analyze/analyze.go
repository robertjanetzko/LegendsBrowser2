package main

import (
	"flag"
	"log"

	"github.com/robertjanetzko/LegendsBrowser2/analyze/df"
)

func main() {
	a := flag.String("a", "", "analyze a file")
	g := flag.Bool("g", false, "generate model")
	flag.Parse()

	if len(*a) > 0 {
		df.AnalyzeStructure(*a)
	}

	if *g {
		err := df.Generate()
		if err != nil {
			log.Fatal(err)
		}
	}
}
