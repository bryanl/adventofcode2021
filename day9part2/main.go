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

func run(text []string) error {

	var heightmap [][]int
	for _, rowText := range text {
		var row []int
		for _, c := range strings.Split(rowText, "") {
			row = append(row, support.ParseInt(c))
		}

		heightmap = append(heightmap, row)
	}

	var basins []int

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
				basins = append(basins, countBasin(heightmap, i, j))
			}
		}
	}

	sort.Ints(basins)
	basins = basins[len(basins)-3:]
	fmt.Println(basins[0] * basins[1] * basins[2])

	return nil
}

func countBasin(heightmap [][]int, row, col int) int {
	n := newNode(row, col, heightmap, nil)
	return n.Search()
}

type point struct {
	row int
	col int
}

func (p *point) String() string {
	return fmt.Sprintf("%d,%d", p.row, p.col)
}

type node struct {
	point     point
	heightmap [][]int
	visited   map[point]bool
}

func newNode(row, col int, heightmap [][]int, visited map[point]bool) *node {
	if visited == nil {
		visited = map[point]bool{}
	}
	n := &node{
		point: point{
			row: row,
			col: col,
		},
		heightmap: heightmap,
		visited:   visited,
	}

	return n
}

func (n *node) IsFirstRow() bool {
	return n.point.row == 0
}

func (n *node) IsLastRow() bool {
	return n.point.row == len(n.heightmap)-1
}

func (n *node) IsFirstColumn() bool {
	return n.point.col == 0
}

func (n *node) IsLastColumn() bool {
	return n.point.col == len(n.heightmap[n.point.row])-1
}

func (n *node) Value() int {
	return n.heightmap[n.point.row][n.point.col]
}

func (n *node) Search() int {
	if _, ok := n.visited[n.point]; ok {
		return 0
	}

	if n.Value() == 9 {
		return 0
	}

	n.visited[n.point] = true

	sum := 1

	// up
	if !n.IsFirstRow() {
		sum += newNode(n.point.row-1, n.point.col, n.heightmap, n.visited).Search()
	}

	// left
	if !n.IsFirstColumn() {
		sum += newNode(n.point.row, n.point.col-1, n.heightmap, n.visited).Search()
	}

	// right
	if !n.IsLastColumn() {
		sum += newNode(n.point.row, n.point.col+1, n.heightmap, n.visited).Search()
	}

	// down
	if !n.IsLastRow() {
		sum += newNode(n.point.row+1, n.point.col, n.heightmap, n.visited).Search()
	}

	return sum
}
