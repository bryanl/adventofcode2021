package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	var input string
	flag.StringVar(&input, "input", "sample.txt", "input")
	flag.Parse()

	if err := run(input); err != nil {
		log.Fatalf("failed: %v", err)
	}
}

func run(input string) error {
	data, err := support.ReadData(input)
	if err != nil {
		return err
	}

	oc := support.ParseBinary(Oxygen(data))
	co2 := support.ParseBinary(CO2(data))

	fmt.Println(oc * co2)

	return nil
}

func Oxygen(appData []string) string {
	col := 0

	data := appData

	for {
		score := countData(data, col)

		if score['0'] == score['1'] {
			data = colMatch('1', col, data)
		} else {
			data = findGreater(data, col, score)
		}

		if len(data) == 1 {
			return data[0]
		}

		col += 1
	}

	return data[0]
}

func CO2(appData []string) string {
	data := appData
	col := 0

	for {
		score := countData(data, col)

		if score['0'] == score['1'] {
			data = colMatch('0', col, data)
		} else {
			data = findLesser(data, col, score)
		}

		if len(data) == 1 {
			return data[0]
		}

		col += 1
	}

}

func findGreater(data []string, col int, score map[uint8]int) []string {
	var cur uint8
	var scoreMax int

	for k, v := range score {
		if v > scoreMax {
			cur = k
			scoreMax = v
		}
	}

	return colMatch(cur, col, data)
}

func findLesser(data []string, col int, score map[uint8]int) []string {
	var cur uint8
	scoreMax := len(data)

	for k, v := range score {
		if v < scoreMax {
			cur = k
			scoreMax = v
		}
	}

	return colMatch(cur, col, data)
}

func colMatch(n byte, col int, data []string) []string {
	var out []string

	for _, item := range data {
		if item[col] == n {
			out = append(out, item)
		}
	}

	return out
}

func countData(data []string, col int) map[uint8]int {
	out := map[uint8]int{}

	for _, item := range data {
		out[item[col]]++
	}

	return out
}
