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

func run(text []string) error {

	var heightmap [][]int
	for _, rowText := range text {
		var row []int
		for _, c := range strings.Split(rowText, "") {
			row = append(row, support.ParseInt(c))
		}

		heightmap = append(heightmap, row)
	}

	var points []int

	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			low := true

			// above
			if i > 0 && heightmap[i][j] >= heightmap[i-1][j] {
				low = false
			}

			// left
			if j > 0 && heightmap[i][j] >= heightmap[i][j-1] {
				low = false
			}

			// right
			if j < len(heightmap[i])-1 && heightmap[i][j] >= heightmap[i][j+1] {
				low = false
			}

			// below
			if i < len(heightmap)-1 && heightmap[i][j] >= heightmap[i+1][j] {
				low = false
			}

			if low {
				fmt.Println(i, j)
				points = append(points, heightmap[i][j])
			}
		}
	}

	spew.Dump(points)
	fmt.Println(calcRisk(points))

	return nil
}

func calcRisk(in []int) int {
	sum := 0

	for _, i := range in {
		sum += i + 1
	}

	return sum
}
