package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

/* utilities */

func SortDescending(ii []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(ii)))
}

func Atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

var Split = strings.Split

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

func InitStacks(state string) []string {
	var stacks []string
	for _, line := range strings.Split(state, "\n") {
		if line != "" {
			if stacks == nil {
				// init stacks
				numStacks := 1 + len(line)/4
				stacks = make([]string, numStacks)
			}
			for i := 1; i < len(line); i += 4 {
				crate := string(line[i])
				if crate != " " {
					stacks[i/4] += string(line[i])
				}
			}
		}
	}
	return stacks
}

// print stack state
func PrintStacks(stacks []string) {
	for i := 0; i < len(stacks); i += 1 {
		fmt.Println(stacks[i])
	}
}

func ApplyMoves1(stacks []string, from, to, num int) {
	for i := 0; i < num; i += 1 {
		ApplyMoves2(stacks, from, to, 1)
	}
}

func ApplyMoves2(stacks []string, from, to, num int) {
	crates := stacks[from][0:num]
	stacks[from] = stacks[from][num:]
	stacks[to] = crates + stacks[to]
}

func GetMsg(stacks []string) string {
	msg := ""
	for i := 0; i < len(stacks); i += 1 {
		msg += string(stacks[i][0])
	}
	return msg
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day5!")

	var msg1, msg2 string
	state, moves := SplitPair(input1, "\n\n")
	stacks1 := InitStacks(state)
	stacks2 := InitStacks(state)
	// apply moves
	var moveExp = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`) // ex. "move 3 from 5 to 2"
	for _, line := range strings.Split(moves, "\n") {
		if line != "" {
			match := moveExp.FindStringSubmatch(line)
			num, from, to := Atoi(match[1]), Atoi(match[2]), Atoi(match[3])
			ApplyMoves1(stacks1, from-1, to-1, num)
			ApplyMoves2(stacks2, from-1, to-1, num)
		}
	}
	//PrintStacks(stacks)
	msg1, msg2 = GetMsg(stacks1), GetMsg(stacks2)

	fmt.Print("puzzle1 solution: [", msg1, "]")
	Assert(msg1 == "MQSHJMWNH", "  <- WRONG!")

	fmt.Print("puzzle2 solution: [", msg2, "]")
	Assert(msg2 == "LLWJRBHVZ", "  <- WRONG!")

}
