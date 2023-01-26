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

func getTopN(n int, in string) []int {
	// top 'n' total values (in ascending order)
	// NB. one more position used to test new possible value
	topN := make([]int, n+1)

	for _, helfStr := range strings.Split(in, "\n\n") {
		tot := 0
		for _, calStr := range strings.Split(helfStr, "\n") {
			tot += Atoi(calStr)
		}
		topN[n] = tot // override smallest with new one
		SortDescending(topN)
	}
	return topN[:n]
}

func main() {
	fmt.Println("AoC 2022, day1!")

	const N = 3
	topN := getTopN(N, input1)

	fmt.Print("puzzle1 solution: [", topN[0], "]")
	Assert(topN[0] == 65912, "  <- WRONG!")

	fmt.Print("puzzle2 solution: [", topN[0]+topN[1]+topN[2], "]")
	Assert(topN[0]+topN[1]+topN[2] == 195625, "  <- WRONG!")
}
