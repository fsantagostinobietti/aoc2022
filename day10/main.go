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

func ExecuteCommand(values []int, cycle int, cmd string) int {
	if cmd == "noop" {
		cycle += 1
		values[cycle] = values[cycle-1]
		return cycle
	}
	// cmd == addx <n>
	_, n := SplitPair(cmd, " ")
	cycle += 1
	values[cycle] = values[cycle-1]
	cycle += 1
	values[cycle] = values[cycle-1] + Atoi(n)
	return cycle
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day10!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	values := make([]int, 400) // big enough
	cycle := 1
	values[cycle] = 1 // initial value at start of cycle 1
	for _, cmd := range strings.Split(input1, "\n") {
		cycle = ExecuteCommand(values, cycle, cmd)
	}

	sum1 := 0
	for c := 20; c <= 220; c += 40 {
		sum1 += c * values[c]
	}
	fmt.Print("puzzle1 solution: [", sum1, "]")
	Assert(sum1 == 14360, "  <- WRONG!")

	fmt.Println("puzzle2 solution:")
	for c := 1; c <= 240; c += 1 {
		x := (c - 1) % 40
		xSprite := values[c]
		if x >= xSprite-1 && x <= xSprite+1 {
			fmt.Print("█")
		} else {
			fmt.Print(" ")
		}

		if c%40 == 0 {
			fmt.Println("")
		}
	}

	//fmt.Print("puzzle2 solution: [", num2, "]")
	//Assert(num2 == "BGKAEREZ", "  <- WRONG!")

}
