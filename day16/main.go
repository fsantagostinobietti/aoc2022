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

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
	tt := Split(s, sep)
	return tt[0], tt[1]
}

func Assert(cond bool, errMsg string) {
	if !cond {
		fmt.Print(errMsg)
	}
	fmt.Print("\n")
}

type Valve struct {
	name string
	rate int
	next []string
}

func LoadValves(input string) []*Valve {
	valves := make([]*Valve, 0)
	for _, line := range Split(input, "\n") {
		name := Split(line, " ")[1]
		rate := Atoi(Split(line[len("Valve AA has flow rate="):], ";")[0])
		next := make([]string, 0)
		for _, tok := range Split(line, ",") {
			tok1 := Split(tok, " ")
			next = append(next, tok1[len(tok1)-1])
		}
		valves = append(valves, &Valve{name, rate, next})
	}
	return valves
}

func Weight(v1, v2 *Valve) int {
	if v1 == v2 {
		return 0
	}
	for _, v := range v1.next {
		if v == v2.name {
			return 1
		}
	}
	return 100000 // no direct arc from v1 to v2
}

func FullGraphDistances(valves []*Valve) map[*Valve]map[*Valve]int {
	// use Floyd-Warshall algo
	dist := make(map[*Valve]map[*Valve]int)
	// init dist
	for _, v1 := range valves {
		for _, v2 := range valves {
			if _, ok := dist[v1]; !ok {
				dist[v1] = make(map[*Valve]int)
			}
			dist[v1][v2] = Weight(v1, v2)
		}
	}
	// run algo
	for _, v := range valves {
		for _, v1 := range valves {
			for _, v2 := range valves {
				if dist[v1][v2] > dist[v1][v]+dist[v][v2] {
					dist[v1][v2] = dist[v1][v] + dist[v][v2]
				}
			}
		}
	}
	return dist
}

func FilterWorthValves(valves []*Valve) []*Valve {
	worths := make([]*Valve, 0)
	for _, v := range valves {
		if v.rate > 0 {
			worths = append(worths, v)
		}
	}
	return worths
}

func MaxPressure(start *Valve, set uint16, mm int) int {
	// check in cache
	key := start.name + "-" + Itoa(int(set)) + "-" + Itoa(mm)
	if v, ok := cache[key]; ok {
		return v
	}

	maxTot := 0
	for i := range worths {
		if (1<<i)&set > 0 { // worths[i] present in set
			valve := worths[i]
			d := distMatrix[start][valve]
			if d+1 > mm {
				continue
			}
			// remove worths[i]
			set1 := set ^ (1 << i)
			maxP := MaxPressure(valve, set1, mm-d-1)
			tot := valve.rate*(mm-d-1) + maxP
			if tot > maxTot {
				maxTot = tot
			}
		}
	}
	// update cache
	cache[key] = maxTot

	return maxTot
}

// global vars
var distMatrix map[*Valve]map[*Valve]int
var worths []*Valve
var cache map[string]int = make(map[string]int)

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day16!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	valves := LoadValves(input1)
	//fmt.Println(valves)
	distMatrix = FullGraphDistances(valves)
	//fmt.Println(distMatrix)
	worths = FilterWorthValves(valves)
	//fmt.Println(worths)
	if len(worths) > 15 {
		panic("required max 15 worth valve!")
	}

	// init cache
	//cache = make(map[string]int)

	valveAA := valves[0]
	var set uint16 = 1<<(len(worths)+1) - 1 // all bits set to 1
	maxP := MaxPressure(valveAA, set, 30)
	fmt.Print("puzzle1 solution: [", maxP, "]")
	Assert(maxP == 1641, "  <- WRONG!")

	maxTot := 0
	var allOnes uint16 = 1<<(len(worths)+1) - 1
	var s1 uint16 = 0
	for ; s1 < allOnes/2+1; s1 += 1 {
		m1 := MaxPressure(valveAA, s1, 30-4)
		var s2 uint16 = allOnes ^ s1
		m2 := MaxPressure(valveAA, s2, 30-4)
		if m1+m2 > maxTot {
			maxTot = m1 + m2
		}
	}

	fmt.Print("puzzle2 solution: [", maxTot, "]")
	Assert(maxTot == 2261, "  <- WRONG!")
}
