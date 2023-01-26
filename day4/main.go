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

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day4!")

	numContained, numOverlapped := 0, 0
	for _, pair := range strings.Split(input1, "\n") {
		if pair != "" {
			r1, r2 := SplitPair(pair, ",")
			//r1MinStr, r1MaxStr := SplitPair(r1, "-")
			//r2MinStr, r2MaxStr := SplitPair(r2, "-")
			r1Min, r1Max := Atoi(Split(r1, "-")[0]), Atoi(Split(r1, "-")[1])
			r2Min, r2Max := Atoi(Split(r2, "-")[0]), Atoi(Split(r2, "-")[1])

			// (r1Min <= r2Min && r1Max >= r2Max) || (r2Min <= r1Min && r2Max >= r1Max)
			isContained := (r1Min-r2Min)*(r1Max-r2Max) <= 0
			if isContained {
				numContained += 1
			}
			// r1Max >= r2Min && r2Max >= r1Min
			isOverlapped := (r1Max-r2Min)*(r2Max-r1Min) >= 0
			if isOverlapped {
				numOverlapped += 1
			}
		}
	}

	fmt.Print("puzzle1 solution: [", numContained, "]")
	Assert(numContained == 602, "  <- WRONG!")

	fmt.Print("puzzle2 solution: [", numOverlapped, "]")
	Assert(numOverlapped == 891, "  <- WRONG!")

}
