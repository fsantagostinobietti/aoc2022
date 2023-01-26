package main

import (
	. "aoc2022/utils"
	_ "embed"
	"fmt"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

const (
	N  = iota // North
	NE        // North-East
	E         // East
	SE        // South-East
	S         // South
	SW        // South-West
	W         // West
	NW        // North-West
)

var DELTAS = [8][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

// elf position
type Elf struct {
	r, c int
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day23!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	elfs := LoadElfs(input1)
	Run(elfs, 10)
	minR, minC, maxR, maxC := MinRectangle(elfs)
	//fmt.Printf("(%d,%d) (%d,%d)\n", minR, minC, maxR, maxC)
	empty := (maxR-minR+1)*(maxC-minC+1) - len(elfs)

	fmt.Print("puzzle1 solution: [", empty, "]")
	Assert(empty == 4025, "  <- WRONG!")

	elfs = LoadElfs(input1)
	n := Run(elfs, MaxInt)
	fmt.Print("puzzle2 solution: [", n, "]")
	Assert(n == 935, "  <- WRONG!")
}

func MinRectangle(elfs map[Elf]bool) (int, int, int, int) {
	minR, minC := MaxInt, MaxInt
	maxR, maxC := MinInt, MinInt
	for e := range elfs {
		minR, maxR = Min(minR, e.r), Max(maxR, e.r)
		minC, maxC = Min(minC, e.c), Max(maxC, e.c)
	}
	return minR, minC, maxR, maxC
}

func Run(elfs map[Elf]bool, maxRounds int) int {
	RULES := [][3]int{{N, NE, NW}, {S, SE, SW}, {W, NW, SW}, {E, NE, SE}}
	ruleIdx := 0
	for n := 0; n < maxRounds; n += 1 {
		hitCount := make(map[Elf]int)
		// compute proposed positions for every elf
		proposed := make([][2]Elf, 0)
		for e := range elfs {
			nexts, num := Neighbours(elfs, e)
			if num == 0 { // no elfs
				// stay in the same position
			} else {
				// propose next position
				e1 := NewPosition(e, nexts, RULES, ruleIdx)
				proposed = append(proposed, [2]Elf{e, e1}) // e -> e1
				// record position hit
				hitCount[e1] = hitCount[e1] + 1
			}
		}
		// move only elfs with non-colliding position
		nMoved := 0
		for i := range proposed {
			e, e1 := proposed[i][0], proposed[i][1]
			if hitCount[e1] == 1 { // no collision
				// move elf
				delete(elfs, e)
				elfs[e1] = true
				nMoved += 1
			}
		}
		if nMoved == 0 { // no more rounds
			return n + 1
		}
		// update rules
		ruleIdx = (ruleIdx + 1) % 4
	}
	return maxRounds
}

func NewPosition(e Elf, nexts []*Elf, RULES [][3]int, ruleIdx int) Elf {
	for i := 0; i < 4; i += 1 { // four rules
		valid := true
		for _, dir := range RULES[ruleIdx] {
			if nexts[dir] != nil {
				valid = false
				break
			}
		}
		if valid {
			moveDir := RULES[ruleIdx][0]
			dr, dc := DELTAS[moveDir][0], DELTAS[moveDir][1]
			return Elf{e.r + dr, e.c + dc}
		}
		ruleIdx = (ruleIdx + 1) % 4
	}
	return e // stay in same position
}

func Neighbours(elfs map[Elf]bool, e Elf) ([]*Elf, int) {
	nexts := make([]*Elf, 0, 8)
	num := 0
	for dir := 0; dir < 8; dir += 1 {
		dr, dc := DELTAS[dir][0], DELTAS[dir][1]
		e1 := Elf{e.r + dr, e.c + dc}
		var n *Elf = nil
		if elfs[e1] { // found neighbour
			n = &e1
			num += 1
		}
		nexts = append(nexts, n)
	}
	return nexts, num
}

func LoadElfs(input string) map[Elf]bool {
	elfs := make(map[Elf]bool)
	for r, line := range Split(input, "\n") {
		for c := 0; c < len(line); c += 1 {
			if line[c] == '#' {
				elfs[Elf{r, c}] = true
			}
		}
	}
	return elfs
}
