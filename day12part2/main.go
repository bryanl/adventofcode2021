package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(input []string) error {
	al := parse(input)

	paths := visit(al, []string{"start"})
	fmt.Println(len(paths))

	return nil
}

var previousLookup = map[string][][]string{}

func visit(al adjacencyList, previous []string) [][]string {
	key := strings.Join(previous, ",")
	if paths, ok := previousLookup[key]; ok {
		return paths
	}

	var paths [][]string

	for _, next := range Available(al, previous) {
		if next == "start" {
			continue
		}

		if next == "end" {
			paths = append(paths, append(previous, next))
			continue
		}

		paths = append(paths, visit(al, append(previous, next))...)
	}

	previousLookup[key] = paths
	return paths
}

var availableLookup = map[string][]string{}

func Available(al adjacencyList, base []string) []string {
	key := strings.Join(base, ",")
	if sl, ok := availableLookup[key]; ok {
		return sl
	}

	var sl []string

	for _, name := range al[base[len(base)-1]] {
		if name == "start" {
			continue
		} else if !isLower(name) || name == "end" {
			sl = append(sl, name)
		} else {
			count := support.CountString(name, base)
			first, found := MaxedCave(base)

			n := 2
			if found && name != first {
				n = 1
			}

			if count < n {
				sl = append(sl, name)
			}
		}
	}

	availableLookup[key] = sl
	return sl
}

var maxedLookup = map[string]string{}

func MaxedCave(base []string) (string, bool) {
	key := strings.Join(base, ",")
	if s, ok := maxedLookup[key]; ok {
		return s, true
	}

	m := map[string]int{}

	for _, name := range base {
		if name == "start" || !isLower(name) {
			continue
		}

		m[name] += 1
		if m[name] == 2 {
			maxedLookup[key] = name
			return name, true
		}
	}

	return "", false
}

var lowerMap = map[string]bool{}

func isLower(s string) bool {
	if tf, ok := lowerMap[s]; ok {
		return tf
	}

	lowerMap[s] = unicode.IsLower(rune(s[0]))
	return lowerMap[s]
}

type adjacencyList map[string][]string

func (al adjacencyList) Connect(start, end string) {
	al[start], al[end] = append(al[start], end), append(al[end], start)
}

func parse(input []string) adjacencyList {
	al := adjacencyList{}

	for _, row := range input {
		connection := strings.Split(row, "-")
		al.Connect(connection[0], connection[1])
	}

	return al
}
