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

func IsVisible(mm []string, row, col int) int {
	v := true
	for c := 0; c < col; c += 1 {
		v = v && (mm[row][c] < mm[row][col])
	}
	if v {
		return 1
	}

	v = true
	for c := col + 1; c < len(mm[row]); c += 1 {
		v = v && (mm[row][c] < mm[row][col])
	}
	if v {
		return 1
	}

	v = true
	for r := 0; r < row; r += 1 {
		v = v && (mm[r][col] < mm[row][col])
	}
	if v {
		return 1
	}

	v = true
	for r := row + 1; r < len(mm); r += 1 {
		v = v && (mm[r][col] < mm[row][col])
	}
	if v {
		return 1
	}

	return 0
}

func ScenicScore(mm []string, row, col int) int {
	score := 1

	v := 0
	c := col - 1
	for ; (c >= 0) && (mm[row][c] < mm[row][col]); c -= 1 {
		v += 1
	}
	if c >= 0 {
		v += 1
	}
	//fmt.Print(v, " ")
	score *= v

	v = 0
	c = col + 1
	for ; (c < len(mm[row])) && (mm[row][c] < mm[row][col]); c += 1 {
		v += 1
	}
	if c < len(mm[row]) {
		v += 1
	}
	//fmt.Print(v, " ")
	score *= v

	v = 0
	r := row - 1
	for ; r >= 0 && (mm[r][col] < mm[row][col]); r -= 1 {
		v += 1
	}
	if r >= 0 {
		v += 1
	}
	//fmt.Print(v, " ")
	score *= v

	v = 0
	r = row + 1
	for ; r < len(mm) && (mm[r][col] < mm[row][col]); r += 1 {
		v += 1
	}
	if r < len(mm) {
		v += 1
	}
	//fmt.Println(v, " ")
	score *= v

	return score
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day8!")

	// init matrix
	mm := []string{}
	var cols int
	rows := 0
	for _, line := range strings.Split(input1, "\n") {
		if line != "" {
			mm = append(mm, line)
			cols = len(line)
			rows += 1
		}
	}

	// test trees visibility
	num1 := 0
	for r := 0; r < rows; r += 1 {
		for c := 0; c < cols; c += 1 {
			num1 += IsVisible(mm, r, c)
		}
	}
	fmt.Print("puzzle1 solution: [", num1, "]")
	Assert(num1 == 1845, "  <- WRONG!")

	num2 := 0
	for r := 0; r < rows; r += 1 {
		for c := 0; c < cols; c += 1 {
			n := ScenicScore(mm, r, c)
			if n > num2 {
				num2 = n
			}
		}
	}
	fmt.Print("puzzle2 solution: [", num2, "]")
	Assert(num2 == 230112, "  <- WRONG!")

	//fmt.Println(ScenicScore(mm, 3, 2))
}
