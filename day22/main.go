package main

import (
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
	"golang.org/x/tools/container/intsets"
)

func main() {

}

func Parse(text []string) []Step {
	var steps []Step

	for _, row := range text {
		steps = append(steps, ParseRow(row))
	}

	return steps
}

func ParseRow(row string) Step {
	step := Step{}

	part1 := strings.SplitN(row, " ", 2)
	step.Status = part1[0]

	part2 := strings.SplitN(part1[1], ",", 3)
	for _, s := range part2 {
		directive := strings.SplitN(s, "=", 2)
		r := ExtractRange(directive[1])

		switch directive[0] {
		case "x":
			step.X = r
		case "y":
			step.Y = r
		case "z":
			step.Z = r
		}
	}

	return step
}

func ExtractRange(s string) Range {
	parts := strings.SplitN(s, "..", 2)
	return Range{
		Min: support.ParseInt(parts[0]),
		Max: support.ParseInt(parts[1]),
	}
}

type Set struct {
	set *intsets.Sparse
}

func NewSet(r Range) *Set {
	s := &Set{
		set: &intsets.Sparse{},
	}

	s.Add(r)

	return s
}

func (s *Set) Add(r Range) {
	min := r.Min
	max := r.Max

	if min > max {
		min, max = max, min
	}

	for i := min; i <= max; i++ {
		s.set.Insert(i)
	}
}

func (s *Set) Sub(r Range) {
	otherSet := NewSet(r)
	s.set.DifferenceWith(otherSet.set)
}

func (s *Set) Range() Range {
	return Range{
		Min: s.set.Min(),
		Max: s.set.Max(),
	}
}

type Range struct {
	Min int
	Max int
}

type Step struct {
	Status string
	X      Range
	Y      Range
	Z      Range
}
