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

type Monkey struct {
	items     []int
	op        string // '+' | '*'
	opValue   int    // int | 0 means 'old'
	testValue int
	testTrue  int
	testFalse int
}

func InitMonkeys(input1 string) []Monkey {
	monkeys := make([]Monkey, 0)
	// init monkeys
	for _, monkeyStruct := range strings.Split(input1, "\n\n") {
		var m Monkey
		for id, line := range strings.Split(monkeyStruct, "\n") {
			//fmt.Println(line)
			switch id {
			case 1:
				itemList := line[len("  Starting items: "):]
				m.items = AtoiArray(strings.Split(itemList, ", "))
			case 2:
				toks := strings.Split(line, " ")
				opStr := toks[len(toks)-2]
				opValueStr := toks[len(toks)-1]
				m.op, m.opValue = opStr, 0
				if opValueStr != "old" {
					m.opValue = Atoi(opValueStr)
				}
			case 3:
				toks := strings.Split(line, " ")
				m.testValue = Atoi(toks[len(toks)-1])
			case 4:
				toks := strings.Split(line, " ")
				m.testTrue = Atoi(toks[len(toks)-1])
			case 5:
				toks := strings.Split(line, " ")
				m.testFalse = Atoi(toks[len(toks)-1])
			}
		}
		monkeys = append(monkeys, m)
	}
	return monkeys
}

func Execute(rounds int, isPuzzle1 bool, monkeys []Monkey, stats []int) {
	mcm := 1 // MCM of test value (prime numbers)
	for _, m := range monkeys {
		mcm *= m.testValue
	}
	for r := 0; r < rounds; r += 1 {
		for m := 0; m < len(monkeys); m += 1 {
			for len(monkeys[m].items) > 0 {
				it := monkeys[m].items[0]
				monkeys[m].items = monkeys[m].items[1:]
				// execute operation
				value := monkeys[m].opValue
				if value == 0 {
					value = it
				}
				switch monkeys[m].op {
				case "+":
					it += value
				case "*":
					it *= value
				}
				// divide by 3 (only puzzle 1)
				if isPuzzle1 {
					it = it / 3
				}
				// test value
				it %= mcm // reduce item value
				test := (it % monkeys[m].testValue) == 0
				if test {
					monkeys[monkeys[m].testTrue].items = append(monkeys[monkeys[m].testTrue].items, it)
				} else {
					monkeys[monkeys[m].testFalse].items = append(monkeys[monkeys[m].testFalse].items, it)
				}
				stats[m] += 1
			}
		}
	}
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day11!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	monkeys := InitMonkeys(input1)
	// executes 20 rounds
	stats := make([]int, len(monkeys))
	Execute(20, true, monkeys, stats)

	SortDescending(stats)
	fmt.Print("puzzle1 solution: [", stats[0]*stats[1], "]")
	Assert(stats[0]*stats[1] == 62491, "  <- WRONG!")

	monkeys = InitMonkeys(input1)
	// executes 10000 rounds
	stats = make([]int, len(monkeys))
	Execute(10000, false, monkeys, stats)
	SortDescending(stats)

	fmt.Print("puzzle2 solution: [", stats[0]*stats[1], "]")
	Assert(stats[0]*stats[1] == 17408399184, "  <- WRONG!")

}
