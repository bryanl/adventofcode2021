package main

import (
	"errors"
	"fmt"
	"log"
	"time"
	"unicode"

	"github.com/bryan/adventofcode2021/internal/support"
)

type Value struct {
	v     int
	depth int
}

type FlatTree []Value

func (v Value) String() string {
	if v.depth >= 0 {
		return fmt.Sprintf("%d(%d)", v.v, v.depth)
	}
	return fmt.Sprintf("%d", v.v)
}

func Parse(s string) FlatTree {
	res := FlatTree{}
	depth := 0
	i := 0
	for i < len(s) {
		switch {
		case s[i] == '[':
			depth += 1
			i += 1
		case s[i] == ']':
			depth -= 1
			i += 1
		case s[i] == ',':
			i += 1
		case unicode.IsDigit(rune(s[i])):
			v := int(s[i] - '0')
			j := i + 1
			for j < len(s) && unicode.IsDigit(rune(s[j])) {
				v = v*10 + int(s[j]-'0')
				j += 1
			}
			i = j
			res = append(res, Value{v, depth})
		default:
			log.Fatalf("unexpected char %c", s[i])
		}
	}
	return res
}

func add(l1 FlatTree, l2 FlatTree) FlatTree {
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}

	res := FlatTree{}
	res = append(res, l1...)
	res = append(res, l2...)
	for i := range res {
		res[i].depth += 1
	}
	return res
}

func removeIndex(s FlatTree, i int) FlatTree {
	// fmt.Printf("removeIndex %d %v\n", i, s)
	res := make(FlatTree, i, len(s)-1)
	copy(res, s[:i])
	res = append(res, s[i+1:]...)
	// fmt.Printf("removeIndex res %v\n", res)
	return res
}

func replaceIndex(s FlatTree, i int, a Value, b Value) FlatTree {
	res := make(FlatTree, i, len(s)+1)
	copy(res, s[:i])
	res = append(res, a)
	res = append(res, b)
	return append(res, s[i+1:]...)
}

// func explode(l FlatTree) (FlatTree, bool) {
// 	for i := 0; i < len(l)-1; i++ {
// 		if l[i].depth >= 5 && l[i].depth == l[i+1].depth {
// 			if i > 0 {
// 				l[i-1].v += l[i].v
// 			}
// 			if i < len(l)-2 {
// 				l[i+2].v += l[i+1].v
// 			}
// 			l[i].v = 0
// 			l[i].depth -= 1
// 			res := removeIndex(l, i+1)
// 			// fmt.Println("explode", res)
// 			return res, true
// 		}
// 	}
// 	return l, false
// }

func split(l FlatTree) (FlatTree, bool) {
	for i := 0; i < len(l); i++ {
		if l[i].v >= 10 {
			a := l[i].v / 2
			b := l[i].v - a
			newDepth := l[i].depth + 1
			res := replaceIndex(l, i, Value{a, newDepth}, Value{b, newDepth})
			return res, true
		}
	}
	return l, false
}

func explodeStar(l FlatTree) (FlatTree, bool) {
	res := l
	reduced := false
	i := 0
	for i < len(res)-1 {
		if res[i].depth >= 5 && res[i].depth == res[i+1].depth {
			left := res[i].v
			right := res[i+1].v
			res = removeIndex(res, i+1)
			res[i].v = 0
			res[i].depth -= 1
			if i > 0 {
				res[i-1].v += left
			}
			if i < len(res)-1 {
				res[i+1].v += right
			}
			reduced = true
		} else {
			i++
		}
	}
	return res, reduced
}

func normalize(l FlatTree) FlatTree {
	reduced := true
	for reduced {
		l, _ = explodeStar(l)
		l, reduced = split(l)
	}
	return l
}

// func normalize(l FlatTree) FlatTree {
// 	res := l
// 	reduced := true
// 	for reduced {
// 		res, reduced = explodeStar(res)
// 		if !reduced {
// 			res, reduced = split(res)
// 		}
// 	}
// 	return res
// }

func (s *Stack) PushMagnitude(v Value) {
	if s.IsEmpty() {
		s.Push(v)
		return
	}

	top, _ := s.Peek()
	if v.depth == top.depth {
		if _, err := s.Pop(); err != nil {
			panic(err)
		}
		s.PushMagnitude(Value{3*top.v + 2*v.v, v.depth - 1})
	} else {
		s.Push(v)
	}
}

func Magnitude(l FlatTree) int {
	stack := BuildStack()
	for _, v := range l {
		stack.PushMagnitude(v)

	}
	top, _ := stack.Pop()
	return top.v
}

func Part1(lines []string) int {
	// e, _ := explodeStar(Parse("[[[[[9,8],1],2],3],4]"))
	// fmt.Printf("e=%v\n", e)
	// return 0
	exp := FlatTree{}
	for _, l := range lines {
		exp = add(exp, Parse(l))
		exp = normalize(exp)
	}
	return Magnitude(exp)
}

func Part2(lines []string) int {
	var values []FlatTree

	for _, l := range lines {
		values = append(values, normalize(Parse(l)))
	}

	max := 0
	for i, a := range values {
		for j, b := range values {
			if i != j {
				m := Magnitude(normalize(add(a, b)))
				if m > max {
					max = m
				}
				m = Magnitude(normalize(add(b, a)))
				if m > max {
					max = m
				}
			}
		}
	}
	return max

}

func main() {
	support.SetupInput(run)

}

func run(rows []string) error {
	fmt.Println("--2021 day 18 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(rows))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(rows))
	fmt.Println(time.Since(start))

	return nil
}

type Stack FlatTree

func BuildStack() Stack {
	return make([]Value, 0)
}

func (s *Stack) Push(c Value) {
	*s = append(*s, c)
}

func (s *Stack) Pop() (Value, error) {
	l := len(*s)
	if l == 0 {
		return Value{}, errors.New("stack is empty")
	}
	top := (*s)[l-1]
	*s = (*s)[:l-1]
	return top, nil
}

func (s *Stack) Peek() (Value, error) {
	if s.IsEmpty() {
		return Value{}, errors.New("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", *s)
}
