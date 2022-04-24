package util

import (
	"encoding/json"
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

func Keys[K comparable, V any](input map[K]V) []K {
	keys := make([]K, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](input map[K]V) []V {
	values := make([]V, 0, len(input))
	for _, v := range input {
		values = append(values, v)
	}
	return values
}

func ContainsAny(s string, substrings ...string) bool {
	for _, substring := range substrings {
		if strings.Contains(s, substring) {
			return true
		}
	}
	return false
}

func MatchesAny(s string, substrings ...string) bool {
	for _, substring := range substrings {
		if ok, _ := regexp.MatchString(s, substring); ok {
			return true
		}
	}
	return false
}

func Title(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to of "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && index > 0 {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func Capitalize(input string) string {
	if input == "" {
		return ""
	}
	return strings.ToUpper(input[:1]) + input[1:]
}

func Json(obj any) template.HTML {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return template.HTML(`<span class="json">` + string(b) + `</span>`)
}

func If[T any](cond bool, v1, v2 T) T {
	if cond {
		return v1
	} else {
		return v2
	}
}

func Map[U, V any](list []U, mapper func(U) V) []V {
	var newList = make([]V, 0, len(list))
	for _, i := range list {
		newList = append(newList, mapper(i))
	}
	return newList
}

type Identifiable interface {
	Id() int
}

func Find[U any](list []*U, predicate func(*U) bool) (*U, bool) {
	for _, x := range list {
		if predicate(x) {
			return x, true
		}
	}
	return nil, false
}
