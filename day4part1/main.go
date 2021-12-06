package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
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

	builder := NewGameBuilder(data)

	game, err := builder.Build()
	if err != nil {
		return err
	}

	card, draw, err := game.Play()
	if err != nil {
		return err
	}

	d, _ := strconv.Atoi(draw)
	fmt.Println(d * card.UnmarkedSum())

	return nil
}

type Game struct {
	Draws []string
	Cards []Card
}

func NewGame(draws []string, cardData [][]string) (*Game, error) {
	game := &Game{
		Draws: draws,
	}

	for _, data := range cardData {
		card, err := NewCard(data)
		if err != nil {
			return nil, err
		}

		game.Cards = append(game.Cards, *card)
	}

	return game, nil
}

func (game *Game) Play() (Card, string, error) {
	for _, draw := range game.Draws {
		for i, card := range game.Cards {
			card.Select(draw)
			if card.IsWin() {
				fmt.Println("board win", i)
				return card, draw, nil
			}
		}
	}
	return Card{}, "", errors.New("no win")
}

type GameBuilder struct {
	Draws []string
	Cards [][]string
	Line  int
	Data  []string
}

func NewGameBuilder(data []string) *GameBuilder {
	builder := &GameBuilder{
		Data: data,
	}

	var state builderState = readDraws

	for state != nil {
		state = state(builder)
	}

	return builder
}

func (builder *GameBuilder) Build() (*Game, error) {
	return NewGame(builder.Draws, builder.Cards)
}

func (builder *GameBuilder) Next() (string, bool) {
	if len(builder.Data) <= builder.Line {
		return "", true
	}

	line := builder.Data[builder.Line]
	builder.Line += 1

	return line, false
}

func (builder *GameBuilder) Rewind() {
	builder.Line -= 1
}

type builderState func(builder *GameBuilder) builderState

func readDraws(builder *GameBuilder) builderState {
	row, isEOF := builder.Next()
	if isEOF {
		return nil
	}

	builder.Draws = strings.Split(row, ",")
	return startCard
}

func startCard(builder *GameBuilder) builderState {
	builder.Cards = append(builder.Cards, []string{})

	for {
		// read until we have data
		line, isEOF := builder.Next()
		if isEOF {
			return nil
		}

		if line != "" {
			builder.Rewind()
			return readCard

		}
	}
}

func readCard(builder *GameBuilder) builderState {
	index := len(builder.Cards) - 1

	re := regexp.MustCompile(`\s+`)

	for {
		card := builder.Cards[index]
		line, isEOF := builder.Next()
		if isEOF {
			return nil
		}

		if line == "" {
			return startCard
		}

		split := re.Split(strings.TrimSpace(line), -1)
		for _, value := range split {
			card = append(card, value)
		}

		builder.Cards[index] = card
	}
}

type Card struct {
	Numbers []Number
	Side    int
}

func NewCard(in []string) (*Card, error) {
	card := &Card{}

	for _, x := range in {
		card.Numbers = append(card.Numbers, Number{Value: x})
	}

	l := len(card.Numbers)
	if !isPerfectSquare(l) {
		return nil, fmt.Errorf("not a square (%d)", l)
	}

	card.Side = int(math.Sqrt(float64(l)))

	return card, nil
}

func (c *Card) Select(value string) {
	for i := range c.Numbers {
		if c.Numbers[i].Value == value {
			c.Numbers[i].Selected = true
		}
	}
}

func (c *Card) IsWin() bool {
	if c.checkRows() {
		fmt.Println("row win")
		return true
	}

	if c.checkColumns() {
		fmt.Println("col win")
		return true
	}

	// if c.checkDiagonals() {
	// 	fmt.Println("diag win")
	// 	return true
	// }

	return false
}

func (c *Card) UnmarkedSum() int {
	var list []int

	for _, n := range c.Numbers {
		if !n.Selected {
			i, err := strconv.Atoi(n.Value)
			if err != nil {
				panic(err)
			}

			list = append(list, i)
		}
	}

	sum := 0
	for _, n := range list {
		sum += n
	}

	return sum
}

func (c *Card) checkRows() bool {
	for row := 0; row < c.Side; row++ {
		win := true
		for column := 0; column < c.Side; column++ {
			position := column + (row * c.Side)
			if !c.IsSelected(position) {
				win = false
			}
		}

		if win {
			return true
		}
	}

	return false
}

func (c *Card) checkColumns() bool {
	for column := 0; column < c.Side; column++ {
		win := true
		for position := 0; position < (c.Side * c.Side); position += c.Side {
			if !c.IsSelected(position + column) {
				win = false
			}
		}

		if win {
			return true
		}
	}

	return false
}

func (c *Card) checkDiagonals() bool {
	win := true
	for position := 0; position < (c.Side * c.Side); position += c.Side + 1 {
		if !c.IsSelected(position) {
			win = false
		}
	}

	if win {
		return true
	}

	win = true
	for position := c.Side - 1; position < (c.Side*c.Side)-1; position += c.Side - 1 {
		if !c.IsSelected(position) {
			win = false
		}
	}
	if win {
		return true
	}

	return false
}

func (c *Card) IsSelected(position int) bool {
	return c.Numbers[position].Selected
}

type Number struct {
	Value    string
	Selected bool
}

func isPerfectSquare(n int) bool {
	root := math.Sqrt(float64(n))
	return int(root*root) == n
}
