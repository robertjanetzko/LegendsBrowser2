package main

import (
	"flag"
	"log"

	"github.com/robertjanetzko/LegendsBrowser2/analyze/df"
)

func main() {
	a := flag.String("a", "", "analyze a file")
	g := flag.Bool("g", false, "generate model")
	e := flag.Bool("e", false, "regenerate events")
	flag.Parse()

	if len(*a) > 0 {
		df.AnalyzeStructure(*a)
	}

	if *g {
		df.LoadSameFields()

		a, err := df.LoadAnalyzeData()
		if err != nil {
			log.Fatal(err)
		}

		df.ListEnumCandidates(a)

		m, err := df.CreateMetadata(a)
		if err != nil {
			log.Fatal(err)
		}
		if err := df.GenerateBackendCode(m); err != nil {
			log.Fatal(err)
		}

		if *e {
			if err := df.GenerateEventsCode(m); err != nil {
				log.Fatal(err)
			}
		}

	}
}
