package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(rows []string) error {
	var scores []int

	for _, row := range rows {
		a := parse(row)
		if a != nil {
			scores = append(scores, score(a))
		}
	}

	sort.Ints(scores)
	index := len(scores) / 2
	fmt.Println(index, scores[index])

	return nil
}

var tags = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func parse(row string) []string {
	var a []string

	corrupted := false

	for _, c := range strings.Split(row, "") {
		tag, ok := tags[c]
		if ok {
			a = append(a, tag)
		} else if isCloseToken(c) {
			var cur string
			cur, a = a[len(a)-1], a[:len(a)-1]
			if cur != c {
				corrupted = true
				break
			}
		}
	}

	if !corrupted {
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}

		return a
	}

	return nil
}

func isCloseToken(in string) bool {
	for _, v := range tags {
		if in == v {
			return true
		}
	}

	return false
}

var tagValue = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func score(sl []string) int {
	sum := 0

	for _, s := range sl {
		sum = sum*5 + tagValue[s]

	}

	return sum
}
