package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

const MaxInt = int(^uint(0) >> 1)

func run(input []string) error {
	runs := 1
	if len(os.Args) > 1 {
		runs = support.ParseInt(os.Args[1])
	}

	score := map[int]int{}
	for _, s := range strings.Split(input[0], ",") {
		count := support.ParseInt(s)
		score[count] += 1
	}

	for day := 1; day <= runs; day++ {
		spawn := score[0]
		score[0] = score[1]
		score[1] = score[2]
		score[2] = score[3]
		score[3] = score[4]
		score[4] = score[5]
		score[5] = score[6]
		score[6] = score[7] + spawn
		score[7] = score[8]
		score[8] = spawn
	}

	total := 0
	for _, v := range score {
		total += v
	}

	fmt.Println(total)

	return nil
}
