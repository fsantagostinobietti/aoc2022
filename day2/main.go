package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/* utilities */

func SortDescending(ii []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(ii)))
}

func Atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func Assert(cond bool, errMsg string) {
	if !cond {
		fmt.Print(errMsg)
	}
	fmt.Print("\n")
}

//go:embed input1.txt
var input1 string

// compute match score
func MatchScore(move1 string, move2 string) int {
	v1 := strings.Index("ABC", move1)
	v2 := strings.Index("XYZ", move2)

	// add my move score
	score := v2 + 1

	// add match score
	d := ((v2 - v1) + 3) % 3   // 0 draw, 1 win, 2 lose
	score += 3 * ((d + 1) % 3) // 3 draw, 6 win, 0 lose

	return score
}

func MatchScore2(move1 string, strategy string) int {
	v1 := strings.Index("ABC", move1)
	s := strings.Index("XYZ", strategy) // 0 lose, 1 draw, 2 win

	// compute my move index
	v2 := (v1 + (s - 1) + 3) % 3
	// add move score
	score := v2 + 1

	// add result score
	score += 3 * s // 0 lose, 3 draw, 6 win

	return score
}

func main() {
	fmt.Println("AoC 2022, day2!")

	scoreTot1 := 0
	for _, matchStr := range strings.Split(input1, "\n") {
		if matchStr != "" {
			xx := strings.Split(matchStr, " ")
			scoreTot1 += MatchScore(xx[0], xx[1])
		}
	}

	fmt.Print("puzzle1 solution: [", scoreTot1, "]")
	Assert(scoreTot1 == 9241, "  <- WRONG!")

	scoreTot2 := 0
	for _, matchStr := range strings.Split(input1, "\n") {
		if matchStr != "" {
			moves := strings.Split(matchStr, " ")
			scoreTot2 += MatchScore2(moves[0], moves[1])
		}
	}

	fmt.Print("puzzle2 solution: [", scoreTot2, "]")
	Assert(scoreTot2 == 14610, "  <- WRONG!")
}
