package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	var input string
	flag.StringVar(&input, "input", "1.txt", "input")
	flag.Parse()

	if err := run(input); err != nil {
		log.Fatalf("failed: %v", err)
	}
}

func run(input string) error {
	data, err := readData(input)
	if err != nil {
		return err
	}

	movements, err := parseMovements(data)
	if err != nil {
		return err
	}

	p := move(movements)

	spew.Dump(p.product())

	return nil
}

func readData(name string) ([]string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	raw := strings.Split(strings.TrimSpace(string(data)), "\n")
	return raw, nil
}

type movement struct {
	direction string
	distance  int
}

func parseMovements(data []string) ([]movement, error) {
	var out []movement
	for _, row := range data {
		parts := strings.SplitN(row, " ", 2)

		d, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		m := movement{
			direction: parts[0],
			distance:  d,
		}

		out = append(out, m)
	}

	return out, nil
}

func move(movements []movement) position {
	p := position{}

	for _, m := range movements {
		p.move(m)
	}

	return p
}

type position struct {
	horizontal int
	depth      int
}

func (p *position) move(m movement) {
	switch m.direction {
	case "forward":
		p.forward(m.distance)
	case "down":
		p.down(m.distance)
	case "up":
		p.up(m.distance)
	}
}

func (p *position) forward(dist int) {
	p.horizontal += dist
}

func (p *position) down(dist int) {
	p.depth += dist
}

func (p *position) up(dist int) {
	p.depth -= dist
}

func (p *position) product() int {
	return p.depth * p.horizontal
}
