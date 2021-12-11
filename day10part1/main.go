package main

import (
	"fmt"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(rows []string) error {
	score := 0

	for _, row := range rows {
		got, ok := parse(row)
		if !ok {
			score += tagValue[got]
		}
	}

	fmt.Println(score)

	return nil
}

var tagValue = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func parse(row string) (string, bool) {
	var a []string

	for _, c := range strings.Split(row, "") {
		tag, ok := tags[c]
		if ok {
			a = append(a, tag)
		} else if isClose(c) {
			var cur string
			cur, a = a[len(a)-1], a[:len(a)-1]
			if cur != c {
				return c, false
			}
		}

	}

	return "", true
}

var tags = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func isOpen(in string) bool {
	_, ok := tags[in]
	return ok
}

func isClose(in string) bool {
	for _, v := range tags {
		if in == v {
			return true
		}
	}

	return false
}
