package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

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

	d := NewDiagnoser()
	for _, item := range data {
		d.Add(item)
	}

	fmt.Printf("result is %d\n", d.Result())

	return nil
}

type Position struct {
	zero int
	one  int
}

func (p *Position) Epsilon() string {
	if p.zero < p.one {
		return "0"
	}

	return "1"
}

func (p *Position) Gamma() string {
	if p.one > p.zero {
		return "1"
	}

	return "0"

}

type Diagnoser struct {
	results map[int]Position
}

func NewDiagnoser() *Diagnoser {
	return &Diagnoser{
		results: map[int]Position{},
	}
}

func (d *Diagnoser) Result() int64 {
	g := support.ParseBinary(d.Gamma())
	fmt.Println("g", d.Gamma(), g)
	e := support.ParseBinary(d.Epsilon())
	fmt.Println("e", d.Epsilon(), e)

	return g * e
}

func (d *Diagnoser) Gamma() string {
	out := ""

	for i := 0; i < len(d.results); i++ {
		v := d.results[i]
		out += v.Gamma()
	}

	return out
}

func (d *Diagnoser) Epsilon() string {
	out := ""

	for i := 0; i < len(d.results); i++ {
		v := d.results[i]
		out += v.Epsilon()
	}

	return out
}

func (d *Diagnoser) Add(in string) {
	parts := strings.Split(in, "")

	for i, r := range in {
		cur, ok := d.results[i]
		if !ok {
			cur = Position{}
		}

		switch r {
		case '0':
			cur.zero++
		case '1':
			cur.one++
		default:
			panic(fmt.Sprintf("can't handle '%s'", parts[i]))
		}

		d.results[i] = cur
	}
}
