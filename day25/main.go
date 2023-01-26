package main

import (
	. "aoc2022/utils"
	_ "embed"
	"fmt"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

var SNAFU2DEC = map[byte]int{'0': 0, '1': 1, '2': 2, '-': -1, '=': -2}
var DEC2SNAFU = map[int]string{0: "0", 1: "1", 2: "2", -1: "-", -2: "="}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day25!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")
	sum := "0"
	for _, s := range Split(input1, "\n") {
		sum = SnafuSum(sum, s)
	}
	fmt.Print("puzzle1 solution: [", sum, "]")
	Assert(sum == "2-2=12=1-=-1=000=222", "  <- WRONG!")
}

func SnafuSum(s1, s2 string) string {
	sLen := Max(len(s1), len(s2)) + 1
	s1 = fmt.Sprintf("%0*s", sLen, s1)
	s2 = fmt.Sprintf("%0*s", sLen, s2)
	sum := ""
	carry := 0
	for i := sLen - 1; i >= 0; i -= 1 {
		s := carry + SNAFU2DEC[s1[i]] + SNAFU2DEC[s2[i]]
		digit := ((s+2)+5)%5 - 2
		sum = DEC2SNAFU[digit] + sum
		carry = (s - digit) / 5
	}
	return strings.TrimLeft(sum, "0")
}
