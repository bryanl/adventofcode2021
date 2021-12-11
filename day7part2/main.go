package main

import (
	"fmt"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(strings []string) error {
	input := support.ParseStringList(strings[0])

	mean := Mean(input)
	answer := Adjust(input, mean)
	fmt.Println(answer)

	return nil
}

var (
	cache = map[int]int{}
)

func Adjust(in []int, to int) int {
	sum := 0

	for _, x := range in {
		r := x - to
		if n, ok := cache[r]; ok {
			sum += n
			continue
		}

		if r < 0 {
			r = -r
		}

		v := TriangleSum(r)
		cache[r] = v
		cache[-r] = v

		sum += v
	}

	return sum
}

func TriangleSum(n int) int {
	return n * (n + 1) / 2
}

func Median(in []int) int {
	l := len(in)

	return in[l/2]
}

func Mean(in []int) int {
	sum := 0
	for i := range in {
		sum += in[i]
	}

	return sum / len(in)
}
