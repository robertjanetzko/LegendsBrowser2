package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

type EntityLeader struct {
	Hf        *HistoricalFigure
	StartYear int
	EndYear   int
}

func (w *DfWorld) LoadHistory() {
	fmt.Println("")

	path := strings.ReplaceAll(w.FilePath, "-legends.xml", "-world_history.txt")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("no world history found")
		} else {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("found world history", path)
	leaderRegEx := regexp.MustCompile(`  \[\*\] (.+?) \(.*?Reign Began: (-?\d+)\)`)
	results := regexp.MustCompile(`\n([^ ].*?), [^\n]+(?:\n [^\n]+)*`).FindAllStringSubmatch(util.ConvertCp473(data), -1)
	for _, result := range results {
		if _, civ, ok := util.FindInMap(w.Entities, nameMatches[*Entity](result[1])); ok {
			leaders := leaderRegEx.FindAllStringSubmatch(result[0], -1)
			var last *EntityLeader
			for _, leader := range leaders {
				year, _ := strconv.Atoi(leader[2])
				l := &EntityLeader{StartYear: year, EndYear: -1}
				if _, hf, ok := util.FindInMap(w.HistoricalFigures, nameMatches[*HistoricalFigure](leader[1])); ok {
					hf.Leader = true
					l.Hf = hf
					civ.Leaders = append(civ.Leaders, l)
				}
				if last != nil {
					last.EndYear = year
				}
				last = l
			}
		}
	}
}

func nameMatches[T Named](name string) func(T) bool {
	name = strings.ToLower(name)
	return func(t T) bool { return t.Name() == name }
}
