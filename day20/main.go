package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(Run)
}

func Run(rows []string) error {
	algo := rows[0]
	rawImage := rows[2:]

	PrintImage(rawImage)
	width := len(rawImage)

	lookup1 := Scan(rawImage)
	image2 := BuildImage(lookup1, algo, width)

	PrintImage(image2)

	lookup2 := Scan(image2)
	fmt.Println("l2 size", len(lookup2))

	image3 := BuildImage(lookup2, algo, width)

	PrintImage(image3)

	fmt.Println(Count(image3))

	return nil
}

func Count(in []string) int {
	sum := 0

	for _, row := range in {
		for _, ch := range row {
			if ch == '#' {
				sum += 1
			}
		}
	}

	return sum
}

func PrintImage(in []string) {
	for i := range in {
		fmt.Println(in[i])
	}
}

func BuildImage(lookup map[int]int64, algo string, width int) []string {
	var update []string
	var row string

	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			cur := i*width + j
			row += string(algo[lookup[cur]])
		}
		update = append(update, row)
		row = ""
	}

	if len(update) != width {
		panic("build is not complete")
	}

	return update
}

func Scan(rawImage []string) map[int]int64 {
	lookup := map[int]int64{}

	sum := 0

	for y, row := range rawImage {
		for x := range strings.Split(row, "") {
			index := y*len(rawImage) + x
			sum += 1
			coordsStr := ""
			if y > 0 {
				if x > 0 {
					coordsStr += CharValue(rawImage[y-1][x-1])
				}

				coordsStr += CharValue(rawImage[y-1][x])

				if x < len(row)-1 {
					coordsStr += CharValue(rawImage[y-1][x+1])
				}

			}

			if x > 0 {
				coordsStr += CharValue(rawImage[y][x-1])
			}

			coordsStr += CharValue(rawImage[y][x])

			if x < len(row)-1 {
				coordsStr += CharValue(rawImage[y][x+1])
			}

			if y < len(rawImage)-1 {
				if x > 0 {
					coordsStr += CharValue(rawImage[y+1][x-1])
				}

				coordsStr += CharValue(rawImage[y+1][x])

				if x < len(row)-1 {
					coordsStr += CharValue(rawImage[y+1][x+1])
				}
			}

			coord, err := strconv.ParseInt(coordsStr, 2, 64)
			if err != nil {
				panic(err)
			}

			lookup[index] = coord
		}
	}

	fmt.Println("scanned", sum, len(lookup))
	return lookup
}

func CharValue(in byte) string {
	if in == '.' {
		return "0"
	}

	return "1"
}
