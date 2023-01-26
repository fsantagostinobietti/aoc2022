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

// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
func PopCount(i uint32) uint32 {
	i = i - ((i >> 1) & 0x55555555)                // add pairs of bits
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333) // quads
	i = (i + (i >> 4)) & 0x0F0F0F0F                // groups of 8
	return (i * 0x01010101) >> 24                  // horizontal sum of bytes
}

func IsMarker(s string) bool {
	var chars uint32
	for i := 0; i < len(s); i += 1 {
		chars |= 1 << (uint32(s[i] - "a"[0]))
	}
	return PopCount(chars) == uint32(len(s))
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day6!")

	line := strings.Split(input1, "\n")[0]
	var num1 int
	for i := 4; i <= len(line); i += 1 {
		if IsMarker(line[i-4 : i]) {
			num1 = i
			break
		}
	}

	fmt.Print("puzzle1 solution: [", num1, "]")
	Assert(num1 == 1142, "  <- WRONG!")

	var num2 int
	for i := 14; i <= len(line); i += 1 {
		if IsMarker(line[i-14 : i]) {
			num2 = i
			break
		}
	}

	fmt.Print("puzzle2 solution: [", num2, "]")
	Assert(num2 == 2803, "  <- WRONG!")
}
