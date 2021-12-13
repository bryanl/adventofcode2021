package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(input []string) error {
	al := parse(input)

	paths := visit(al, []string{"start"}, 0)
	fmt.Println(printPaths(paths))
	fmt.Println(len(paths))

	return nil
}

func printPaths(paths [][]string) string {
	sort.Slice(paths, func(i, j int) bool {
		a := strings.Join(paths[i], ",")
		b := strings.Join(paths[j], ",")

		return a < b
	})

	var sb strings.Builder

	for _, path := range paths {
		sb.WriteString(printPath(path) + "\n")
	}

	return sb.String()
}

func printPath(path []string) string {
	return strings.Join(path, ",")
}

func visit(al adjacencyList, previous []string, level int) [][]string {
	prev := append(previous)

	current := prev[len(prev)-1]
	var paths [][]string

	changes := 0

	for _, next := range available(al, prev, current) {
		if next == "start" {
			continue
		}

		if next == "end" {
			paths = append(paths, append(prev, next))
			continue
		}

		paths = append(paths, visit(al, append(prev, next), level+1)...)
		changes += 1
	}

	dup := make([][]string, len(paths))
	for i := range paths {
		dup[i] = make([]string, len(paths[i]))
		copy(dup[i], paths[i])
	}

	return dup
}

func available(al adjacencyList, base []string, current string) []string {
	var sl []string
	for _, c := range al[current] {
		if onlyOnce(c) && !support.ContainsString(c, base) {
			sl = append(sl, c)
		} else if !onlyOnce(c) {
			sl = append(sl, c)
		}
	}

	return sl
}

func onlyOnce(s string) bool {
	return unicode.IsLower(rune(s[0]))
}

type adjacencyList map[string][]string

func (al adjacencyList) Connect(start, end string) {
	al[start] = append(al[start], end)
	al[end] = append(al[end], start)
}

func (al adjacencyList) Remove(s string) adjacencyList {
	out := adjacencyList{}
	for k, v := range al {
		if s == k {
			continue
		}

		out[k] = append(out[k], v...)
	}

	return out
}

func parse(input []string) adjacencyList {
	al := adjacencyList{}

	for _, row := range input {
		connection := strings.Split(row, "-")
		al.Connect(connection[0], connection[1])
	}

	return al
}
