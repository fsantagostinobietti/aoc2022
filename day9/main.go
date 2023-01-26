package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

/* utilities */

func SortDescending(ii []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(ii)))
}
func SortAscending(ii []int) {
	sort.Ints(ii)
}

func Atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

var Itoa = strconv.Itoa

var Split = strings.Split

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func Sign(v int) int {
	if v == 0 {
		return 0
	}
	return v / Abs(v)
}

func SplitPair(s string, sep string) (string, string) {
	tt := strings.Split(s, sep)
	return tt[0], tt[1]
}

func Assert(cond bool, errMsg string) {
	if !cond {
		fmt.Print(errMsg)
	}
	fmt.Print("\n")
}

type Position struct {
	x, y int
}

func MoveHead(head Position, direction string) Position {
	// move head
	switch direction {
	case "U":
		head.y += 1
	case "D":
		head.y -= 1
	case "R":
		head.x += 1
	case "L":
		head.x -= 1
	}
	return head
}

func MoveTail(head, tail Position) Position {
	// update tail
	dx, dy := head.x-tail.x, head.y-tail.y
	if Abs(dx) > 1 || Abs(dy) > 1 {
		// move diagonally
		tail.x += Sign(dx)
		tail.y += Sign(dy)
	}
	return tail
}

func MoveKnots(knots []Position, direction string) []Position {
	knots[0] = MoveHead(knots[0], direction)
	for i := 1; i < len(knots); i += 1 {
		knots[i] = MoveTail(knots[i-1], knots[i])
	}
	return knots
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day9!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	var head, tail Position
	visited := make(map[Position]bool)
	visited[tail] = true
	for _, line := range strings.Split(input1, "\n") {
		direction, count := SplitPair(line, " ")
		for c := 0; c < Atoi(count); c += 1 {
			head = MoveHead(head, direction)
			tail = MoveTail(head, tail)
			visited[tail] = true
		}
	}
	num1 := len(visited)
	fmt.Print("puzzle1 solution: [", num1, "]")
	Assert(num1 == 6337, "  <- WRONG!")

	knots := make([]Position, 10) // knots[0]==head, ..., knots[9]==tail
	visited = make(map[Position]bool)
	visited[knots[9]] = true
	for _, line := range strings.Split(input1, "\n") {
		direction, count := SplitPair(line, " ")
		for c := 0; c < Atoi(count); c += 1 {
			knots = MoveKnots(knots, direction)
			visited[knots[9]] = true
		}
	}
	num2 := len(visited)
	fmt.Print("puzzle2 solution: [", num2, "]")
	Assert(num2 == 2455, "  <- WRONG!")

}
