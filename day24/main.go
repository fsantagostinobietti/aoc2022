package main

import (
	. "aoc2022/utils"
	_ "embed"
	"fmt"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

const (
	N = iota // North
	E        // East
	S        // South
	W        // West
	NOOP
)

var DELTAS = [5][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}, {0, 0}}

// position
type Pos struct {
	r, c int
}

type Blizard struct {
	pos Pos
	dir int // N, E, S, W
}

// global
var blizardsLast []Blizard

// map[Pos]int
var rows, cols int
var start, end Pos

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day24!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	Init(input1)
	//fmt.Println("rows:", rows, "cols:", cols, "start:", start, "end:", end)

	min1 := BestTime(start, end, 0)
	fmt.Print("puzzle1 solution: [", min1, "]")
	Assert(min1 == 299, "  <- WRONG!")

	min2 := BestTime(end, start, min1)
	min2 = BestTime(start, end, min2)
	fmt.Print("puzzle2 solution: [", min2, "]")
	Assert(min2 == 899, "  <- WRONG!")
}

func BestTime(initial Pos, final Pos, t int) int {
	currPos := make(map[Pos]bool)
	currPos[initial] = true
	for {
		//fmt.Println("t:", t, currPos)
		t += 1
		nextPos := make(map[Pos]bool)
		blizardsSet := NextBlizardsSet()
		for pos := range currPos {
			//fmt.Println(blizardsSet)
			for _, dd := range DELTAS {
				pos1 := Pos{pos.r + dd[0], pos.c + dd[1]}
				if pos1 == final { // found solution
					return t
				}
				if isInBoard(pos1) && isEmpty(pos1, blizardsSet) {
					nextPos[pos1] = true
				}
			}
		}
		currPos = nextPos
	}
}

func isInBoard(pos Pos) bool {
	if pos == start || pos == end {
		return true
	}
	if pos.r < 0 || pos.r >= rows || pos.c < 0 || pos.c >= cols {
		return false
	}
	return true
}

func isEmpty(pos Pos, set map[Pos]bool) bool {
	_, isPresent := set[pos]
	return !isPresent
}

func NextBlizardsSet() map[Pos]bool {
	var blizardsSet map[Pos]bool
	blizardsLast, blizardsSet = NextBlizards(blizardsLast)
	return blizardsSet
}

func PositionSet(blizards []Blizard) map[Pos]bool {
	set := make(map[Pos]bool)
	for _, b := range blizards {
		set[b.pos] = true
	}
	return set
}

func NextBlizards(bb []Blizard) ([]Blizard, map[Pos]bool) {
	bb1 := make([]Blizard, 0)
	set1 := make(map[Pos]bool)
	for _, b := range bb {
		d := DELTAS[b.dir]
		p1 := Pos{b.pos.r + d[0], b.pos.c + d[1]}
		if p1.r < 0 || p1.r == rows {
			// wrap vertically
			p1.r = (p1.r + rows) % rows
		}
		if p1.c < 0 || p1.c == cols {
			// wrap orizontally
			p1.c = (p1.c + cols) % cols
		}
		b.pos = p1
		bb1 = append(bb1, b)
		set1[b.pos] = true
	}
	return bb1, set1
}

func Init(input string) {
	blizards := make([]Blizard, 0)
	lines := Split(input, "\n")
	for r, line := range lines {
		if r == 0 {
			c := strings.Index(line, ".") - 1
			start = Pos{r - 1, c}
		} else if r == len(lines)-1 {
			c := strings.Index(line, ".") - 1
			end = Pos{r - 1, c}
		} else {
			for c, char := range line {
				if char != '#' && char != '.' { // blizard found
					pos := Pos{r - 1, c - 1}
					dir := strings.Index("^>v<", string(char))
					blizards = append(blizards, Blizard{pos, dir})
				}
			}
		}
	}
	blizardsLast = blizards

	cols = len(lines[0]) - 2 // remove walls
	rows = len(lines) - 2    // ignore first and last row
}
