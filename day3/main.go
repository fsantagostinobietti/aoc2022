package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
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

func Priority(char byte) int {
	if char >= "a"[0] {
		return int(1 + (char - "a"[0]))
	}
	return int(27 + (char - "A"[0]))
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day3!")

	sumPri1 := 0
	for _, rucksack := range strings.Split(input1, "\n") {
		if rucksack != "" {
			// empty type array
			arr := make([]bool, 1+52)

			i := 0
			for ; i < len(rucksack)/2; i += 1 {
				p := Priority(rucksack[i])
				arr[p] = true
			}
			for ; i < len(rucksack); i += 1 {
				p := Priority(rucksack[i])
				if arr[p] {
					sumPri1 += p
					break
				}
			}
		}
	}

	fmt.Print("puzzle1 solution: [", sumPri1, "]")
	Assert(sumPri1 == 8233, "  <- WRONG!")

	sumPri2 := 0
	elf := 0 // index in a group
	var arr []byte
	for _, rucksack := range strings.Split(input1, "\n") {
		if rucksack != "" {
			if elf == 0 {
				arr = make([]byte, 1+52) // init type array
			}

			for i := 0; i < len(rucksack); i += 1 {
				p := Priority(rucksack[i])
				arr[p] |= 1 << elf
				if elf == 2 && arr[p] == 0b111 {
					sumPri2 += p
					break // found
				}
			}

			elf += 1
			elf = elf % 3
		}
	}

	fmt.Print("puzzle2 solution: [", sumPri2, "]")
	Assert(sumPri2 == 2821, "  <- WRONG!")

	var s mapset.Set[string]
	s1 := mapset.NewSet(strings.Split("abacddef", "")...)
	s2 := mapset.NewSet(strings.Split("aacddfghh", "")...)
	s = s1.Intersect(s2)
	fmt.Print(s)
}
