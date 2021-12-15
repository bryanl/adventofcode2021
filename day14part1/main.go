package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(rows []string) error {
	data := Parse(rows)

	for i := 0; i < 40; i++ {
		fmt.Println(i)
		data.Perform()
	}

	fmt.Println(data.Score())

	return nil
}

type Data struct {
	Template string

	Rules map[string]string
}

var (
	reTemplate = regexp.MustCompile(`^([A-Z]+)$`)
	reInsert   = regexp.MustCompile(`^(\w{2}) -> (\w)$`)
)

func Parse(rows []string) *Data {
	data := &Data{
		Rules: map[string]string{},
	}

	for _, row := range rows {
		switch {
		case reTemplate.MatchString(row):
			match := reTemplate.FindStringSubmatch(row)
			data.Template = match[1]
		case reInsert.MatchString(row):
			match := reInsert.FindStringSubmatch(row)
			data.Rules[match[1]] = match[2]
		}
	}

	return data
}

func (d *Data) Count() map[string]int {
	m := map[string]int{}

	for _, s := range strings.Split(d.Template, "") {
		m[s] += 1
	}

	return m
}

func (d *Data) Score() int {
	count := d.Count()
	max := 0
	min := 99999999

	for _, v := range count {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max - min
}

func (d *Data) Pairs() []string {
	var pairs []string

	for i := 0; i < len(d.Template)-1; i++ {
		pair := d.Template[i : i+2]
		pairs = append(pairs, pair)
	}

	return pairs
}

func (d *Data) Perform() {
	var sb strings.Builder

	pairs := d.Pairs()

	for i, pair := range pairs {
		update := pair

		if v, ok := d.Rules[pair]; ok {
			update = pair[:1] + v + pair[1:]
		}

		if i > 0 {
			update = update[1:]
		}
		sb.WriteString(update)
	}

	d.Template = sb.String()
}
