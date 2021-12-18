package main

import (
	"fmt"
	"strings"

	"github.com/bryan/adventofcode2021/internal/support"
	"github.com/davecgh/go-spew/spew"
)

func main() {

}

type Pair interface {
	SetLeft(Pair)
	Left() Pair
	SetRight(Pair)
	Right() Pair
	Level() int
}

type Number struct {
	Value int
	level int
}

var _ Pair = &Number{}

func NewNumber(value, level int) *Number {
	return &Number{
		Value: value,
		level: level,
	}
}

func (n Number) SetLeft(pair Pair) {}

func (n Number) Left() Pair {
	return nil
}

func (n Number) SetRight(pair Pair) {}

func (n Number) Right() Pair {
	return nil
}

func (n Number) Level() int {
	return n.level
}

type ComplexPair struct {
	left  Pair
	right Pair
	level int
}

func (c *ComplexPair) SetLeft(pair Pair) {
	c.left = pair
}

func (c *ComplexPair) Left() Pair {
	return c.left
}

func (c *ComplexPair) SetRight(pair Pair) {
	c.right = pair
}

func (c *ComplexPair) Right() Pair {
	return c.right
}

func (c *ComplexPair) Level() int {
	return c.level
}

var _ Pair = &ComplexPair{}

func CreatePair(input string) (Pair, error) {
	l := &lexer{
		input: input,
		pos:   0,
		pair:  nil,
	}

	return lexerStart(l, 0)
}

type lexer struct {
	input string
	pos   int
	pair  Pair
	err   error
}

func (l *lexer) token() rune {
	return rune(l.input[l.pos])
}

func (l *lexer) inc() {
	l.pos += 1
}

func (l *lexer) unexpectedCharError(expected rune) error {
	return fmt.Errorf(
		"expected '%s', got %s at %d",
		string(expected),
		string(l.token()),
		l.pos,
	)
}

func (l *lexer) slurpNumber() int {
	var sb strings.Builder

	for {
		if t := l.token(); tokenIsNumber(t) {
			sb.WriteString(string(t))
			l.inc()
			break
		}
	}

	return support.ParseInt(sb.String())
}

func lexerStart(l *lexer, level int) (Pair, error) {
	level += 1

	if l.token() != '[' {
		return nil, l.unexpectedCharError('[')
	}

	l.inc()

	left, err := collectNumber(l, level)
	if err != nil {
		return nil, fmt.Errorf("collect left: %w", err)
	}

	if l.token() != ',' {
		spew.Dump(left)
		return nil, l.unexpectedCharError(',')
	}
	l.inc()

	right, err := collectNumber(l, level)
	if err != nil {
		return nil, fmt.Errorf("collect right: %w", err)
	}

	if l.token() != ']' {
		return nil, l.unexpectedCharError(']')
	}
	l.inc()

	pair := &ComplexPair{}
	pair.SetLeft(left)
	pair.SetRight(right)

	return pair, nil
}

func collectNumber(l *lexer, level int) (Pair, error) {
	t := l.token()
	if t == '[' {
		return lexerStart(l, level)
	} else if tokenIsNumber(t) {
		return NewNumber(l.slurpNumber(), level), nil
	}

	return nil, fmt.Errorf("expected [ or digit at %d. Got %s", l.pos, string(l.token()))
}

func tokenIsNumber(token rune) bool {
	return token >= 48 && token <= 57
}