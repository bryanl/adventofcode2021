package main

import (
	"fmt"
	"regexp"

	"github.com/bryan/adventofcode2021/internal/support"
)

func main() {
	support.SetupInput(run)
}

func run(rows []string) error {
	data := Parse(rows)

	pairs := map[string]int{}
	counter := map[string]int{}

	for i := 0; i < len(data.Template)-1; i++ {
		pair := string(data.Template[i]) + string(data.Template[i+1])
		pairs[pair] += 1
		counter[string(data.Template[i])] += 1

	}

	counter[string(data.Template[len(data.Template)-1])] += 1

	steps := 10

	for i := 0; i < steps; i++ {
		pairs = work(pairs, counter, data.Rules)

	}

	var maxSl []int
	for _, v := range counter {
		maxSl = append(maxSl, v)
	}
	max := support.MaxInt(maxSl)

	var minSl []int
	for _, v := range counter {
		minSl = append(minSl, v)
	}
	min := support.MinInt(minSl)

	fmt.Println(max - min)

	return nil
}

func work(pairs map[string]int, counter map[string]int, rules map[string]string) map[string]int {
	update := map[string]int{}
	for pair := range pairs {
		if element, ok := rules[pair]; ok {
			new1 := pair[:1] + element
			new2 := element + pair[1:]
			update[new1] += pairs[pair]
			update[new2] += pairs[pair]
			counter[element] += pairs[pair]
		} else {
			update[pair] = pairs[pair]
		}
	}

	return update
}

type Data struct {
	Template string
	Rules    map[string]string
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
