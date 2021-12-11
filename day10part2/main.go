package main

import (
	"fmt"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	support.SetupInput(run)
}

func run(rows []string) error {
	score := 0

	for _, row := range rows {
		_, ok := parse(row)
		if ok {
			fmt.Println(row)
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

func parse(row string) ([]string, bool) {
	var a []string

	//

	for _, c := range strings.Split(row, "") {
		tag, ok := tags[c]
		if ok {
			a = append(a, tag)
		} else if isClose(c) {
			cur := a[len(a)-1]
			if cur != c {
				fmt.Println("hello", c, a)
				return nil, false
			} else {
			}
		}
	}

	fmt.Println(a)

	spew.Dump(a)

	return a, true
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
