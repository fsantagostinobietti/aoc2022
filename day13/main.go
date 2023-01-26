package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

func AtoiArray(a []string) []int {
	r := make([]int, len(a))
	for i, item := range a {
		r[i] = Atoi(item)
	}
	return r
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

func IsInt(s string) bool {
	//fmt.Println(">" + s + "<")
	return s[0] != '['
}

func NextItem(s string) (string, string) {
	item, rem := "", ""
	if IsInt(s) { // next is a int
		idx := 0
		for ; idx < len(s) && unicode.IsDigit(rune(s[idx])); idx += 1 {
		}
		item, rem = s[:idx], s[idx:]
		if idx < len(s) {
			rem = s[idx+1:] // skip ','
		}
	} else { // next is a list
		open := 1 // open brackets
		idx := 1
		for ; open > 0; idx += 1 {
			if s[idx] == '[' {
				open += 1
			} else if s[idx] == ']' {
				open -= 1
			}
		}
		item, rem = s[:idx], s[idx:]
		if idx < len(s) {
			rem = s[idx+1:] // skip ','
		}
	}
	//fmt.Println(item, rem)
	return item, rem
}

func AreInOrder(left, right string) int {
	if len(left) == 0 && len(right) == 0 {
		return 0
	}
	if len(left) == 0 {
		return 1
	}
	if len(right) == 0 {
		return -1
	}
	itemL, remL := NextItem(left)
	itemR, remR := NextItem(right)
	if IsInt(itemL) && IsInt(itemR) {
		if Atoi(itemL) < Atoi(itemR) {
			return 1
		}
		if Atoi(itemL) > Atoi(itemR) {
			return -1
		}
		return AreInOrder(remL, remR)
	} else { // al least one list
		// remove external brackets to lists
		if !IsInt(itemL) {
			itemL = itemL[1 : len(itemL)-1]
		}
		if !IsInt(itemR) {
			itemR = itemR[1 : len(itemR)-1]
		}
		test := AreInOrder(itemL, itemR)
		if test != 0 {
			return test
		}
		return AreInOrder(remL, remR)
	}
}

func Order(s []string) []string {
	sort.Slice(s, func(i, j int) bool {
		left := s[i][1 : len(s[i])-1]
		right := s[j][1 : len(s[j])-1]
		return AreInOrder(left, right) > 0
	})
	return s
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day13!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	sum1 := 0
	for i, pair := range strings.Split(input1, "\n\n") {
		left, right := SplitPair(pair, "\n")
		// since they are lists, remove external brackets
		left = left[1 : len(left)-1]
		right = right[1 : len(right)-1]
		if AreInOrder(left, right) > 0 {
			sum1 += (i + 1)
		}
	}

	fmt.Print("puzzle1 solution: [", sum1, "]")
	Assert(sum1 == 6395, "  <- WRONG!")

	input1 = strings.Replace(input1, "\n\n", "\n", -1)
	input1 += "\n[[2]]\n[[6]]"
	ordered := Order(strings.Split(input1, "\n"))
	key := 1
	for i, packet := range ordered {
		if packet == "[[2]]" || packet == "[[6]]" {
			key *= (i + 1)
		}
	}
	fmt.Print("puzzle2 solution: [", key, "]")
	Assert(key == 24921, "  <- WRONG!")

}
