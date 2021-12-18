package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/RyanCarrier/dijkstra"
	"github.com/bryan/adventofcode2021/internal/support"
)

var writer = bufio.NewWriter(os.Stdout)

func println(f string) { fmt.Fprintln(writer, f) }

var lines []string
var result = 0

var problem [][]int

func main() {
	support.SetupInput(run)
}

func run(rows []string) error {
	lines = rows
	stringLinesTo2dIntArray()
	result = solve()
	fmt.Println(result)
	return nil
}

func stringLinesTo2dIntArray() {
	problem = make([][]int, len(lines))

	for i, line := range lines {
		problem[i] = make([]int, len(line))

		for j, char := range line {
			problem[i][j] = int(char) - 48
		}
	}
}

func solve() int {
	graph := dijkstra.NewGraph()

	for i, row := range problem {
		for j := range row {
			graph.AddVertex(i*len(row) + j)
		}
	}

	for i, row := range problem {
		for j, value := range row {
			index := i*len(row) + j
			if checkLegitPoint(i-1, j) {
				graph.AddArc((i-1)*len(row)+j, index, int64(value))
			}
			if checkLegitPoint(i+1, j) {
				graph.AddArc((i+1)*len(row)+j, index, int64(value))
			}
			if checkLegitPoint(i, j-1) {
				graph.AddArc((i)*len(row)+j-1, index, int64(value))
			}
			if checkLegitPoint(i, j+1) {
				graph.AddArc((i)*len(row)+j+1, index, int64(value))
			}
		}
	}

	best, err := graph.Shortest(0, len(problem)*len(problem[0])-1)
	if err != nil {
		log.Fatal(err)
	}

	return int(best.Distance)
}

func checkLegitPoint(x, y int) bool {
	if x < 0 || x >= len(problem) {
		return false
	}

	if y < 0 || y >= len(problem[x]) {
		return false
	}

	return true
}

//
// import (
// 	"fmt"
// 	"log"
// 	"strings"
//
// 	"github.com/RyanCarrier/dijkstra"
// 	"github.com/bryan/adventofcode2021/internal/support"
// )
//
// func main() {
// 	support.SetupInput(run)
// }
//
// func run(rows []string) error {
// 	grid := Parse2(rows)
// 	solution := Solve(grid)
// 	fmt.Println(solution)
//
// 	return nil
// }
//
// func Parse2(rows []string) [][]int {
// 	out := make([][]int, len(rows))
// 	for i, row := range rows {
// 		out[i] = make([]int, len(row))
//
// 		for j, char := range strings.Split(row, "") {
// 			out[i][j] = support.ParseInt(char)
// 		}
// 	}
//
// 	return out
// }
//
// func Solve(grid [][]int) int {
// 	graph := dijkstra.NewGraph()
//
// 	for i, row := range grid {
// 		for j := range row {
// 			graph.AddVertex(i * len(row) * j)
// 		}
// 	}
//
// 	for i, row := range grid {
// 		for j, value := range row {
// 			index := i*len(row) + j
// 			if verify(grid, i-1, j) {
// 				graph.AddArc((i-1)*len(row)+j, index, int64(value))
// 			}
// 			if verify(grid, i+1, j) {
// 				graph.AddArc((i+1)*len(row)+j, index, int64(value))
// 			}
// 			if verify(grid, i, j-1) {
// 				graph.AddArc((i)*len(row)+j-1, index, int64(value))
// 			}
// 			if verify(grid, i, j+1) {
// 				graph.AddArc((i)*len(row)+j+1, index, int64(value))
// 			}
// 		}
// 	}
//
// 	best, err := graph.Shortest(0, len(grid)*len(grid[0])-1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	return int(best.Distance)
// }
//
// func verify(grid [][]int, x, y int) bool {
// 	if x < 0 || x >= len(grid) {
// 		return false
// 	}
//
// 	if y < 0 || y >= len(grid[x]) {
// 		return false
// 	}
//
// 	return true
// }
