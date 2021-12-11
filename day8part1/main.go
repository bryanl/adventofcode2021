package main

import (
	"fmt"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(text []string) error {
	var inputs []Input

	for _, s := range text {
		inputs = append(inputs, ReadInput(s))
	}

	sum := 0
	for _, input := range inputs {
		sum += CountKnown(input)
	}

	fmt.Println(sum)

	return nil
}

type Input struct {
	SignalPatterns []string
	OutputValue    []string
}

func ReadInput(in string) Input {
	parts := strings.SplitN(in, "|", 2)

	i := Input{
		SignalPatterns: strings.Split(strings.TrimSpace(parts[0]), " "),
		OutputValue:    strings.Split(strings.TrimSpace(parts[1]), " "),
	}
	return i
}

func CountKnown(in Input) int {
	sum := 0

	for _, s := range in.OutputValue {
		if _, ok := GuessByLength(s); ok {
			sum += 1
		}
	}

	return sum
}

var lenLookup = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

func GuessByLength(in string) (int, bool) {
	if res, ok := lenLookup[len(in)]; ok {
		return res, true
	}

	return 0, false
}
