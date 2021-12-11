package main

import (
	"fmt"
	"sort"

	"github.com/bryan/adventofcode2021/internal/support"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	support.SetupInput(run)
}

func run(strings []string) error {
	input := support.ParseStringList(strings[0])

	sort.Ints(input)
	spew.Dump(input)

	median := Median(input)
	fmt.Println("the median is", median)
	fmt.Println("adjustment", Adjust(input, median))

	return nil
}

func Adjust(in []int, to int) int {
	sum := 0

	for i, x := range in {
		r := x - to
		if r < 0 {
			r = -r
		}

		fmt.Printf("%d: %d -> %d\n", i, x, r)

		sum += r
	}

	return sum
}

func Median(in []int) int {
	l := len(in)

	return in[l/2]
}
