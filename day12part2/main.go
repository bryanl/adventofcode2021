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

	fmt.Println(visit2(al, []string{"start"}))

	return nil
}

var previousLookup = map[string][][]string{}

func visit2(al adjacencyList, previous []string) int {
	sum := 0

	for _, next := range Available(al, previous) {
		if next == "end" {
			sum += 1
			continue
		}

		sum += visit2(al, append(previous, next))
	}

	return sum
}

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

func Available(al adjacencyList, base []string) []string {
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

	return sl
}

func MaxedCave(base []string) (string, bool) {
	m := map[string]int{}

	for _, name := range base {
		if name == "start" || !isLower(name) {
			continue
		}

		m[name] += 1
		if m[name] == 2 {
			return name, true
		}
	}

	return "", false
}

func isLower(s string) bool {
	return unicode.IsLower(rune(s[0]))
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
