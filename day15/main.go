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
	tt := strings.Split(s, sep)
	return tt[0], tt[1]
}

func Assert(cond bool, errMsg string) {
	if !cond {
		fmt.Print(errMsg)
	}
	fmt.Print("\n")
}

type Pos struct {
	x, y int
}

func LoadInput(input string) ([]Pos, []Pos) {
	sensors := make([]Pos, 0)
	beacons := make([]Pos, 0)
	for _, line := range strings.Split(input, "\n") {
		s, b := SplitPair(line[len("Sensor at "):], ": closest beacon is at ")
		sx, sy := SplitPair(s[len("x="):], ", y=")
		sensors = append(sensors, Pos{Atoi(sx), Atoi(sy)})
		bx, by := SplitPair(b[len("x="):], ", y=")
		beacons = append(beacons, Pos{Atoi(bx), Atoi(by)})
	}
	return sensors, beacons
}

func Distance(s Pos, b Pos) int {
	// Manhattan distance
	return Abs(b.x-s.x) + Abs(b.y-s.y)
}

func RangeXY(sensors []Pos, beacons []Pos) (Pos, Pos) {
	min := Pos{MaxInt, MaxInt}
	max := Pos{MinInt, MinInt}
	for i := range sensors {
		D := Distance(sensors[i], beacons[i])
		min.x = Min(min.x, sensors[i].x-D)
		min.y = Min(min.y, sensors[i].y-D)
		max.x = Max(max.x, sensors[i].x+D)
		max.y = Max(max.y, sensors[i].y+D)
	}
	return min, max
}

func Contains(beacons []Pos, p Pos) bool {
	for _, b := range beacons {
		if p == b {
			return true
		}
	}
	return false
}

func SearchInRow(sensors []Pos, beacons []Pos, y int, min, max Pos) int {
	totNonBeacon := 0
	for x := min.x; x <= max.x; x += 1 {
		p := Pos{x, y}
		if Contains(beacons, p) {
			continue
		}
		for i := range sensors {
			D := Distance(sensors[i], beacons[i])
			if Distance(sensors[i], p) <= D {
				totNonBeacon += 1
				break
			}
		}
	}
	return totNonBeacon
}

type Range struct {
	low, hi int
}

func SearchBeacon(sensors []Pos, beacons []Pos, min, max Pos) int {
	freq := -1
	for y := min.y; y <= max.y; y += 1 {
		// init beacon range
		bRange := []Range{{min.x, max.x}}
		// narrow beacon range using sensors
		for i := range sensors {
			D := Distance(sensors[i], beacons[i])
			dx := D - Abs(sensors[i].y-y) + 1
			if dx > 0 {
				bRange = IntersectRanges(bRange, Range{min.x, sensors[i].x - dx}, Range{sensors[i].x + dx, max.x})
				if len(bRange) == 0 {
					break
				}
			}
		}
		if len(bRange) == 1 && bRange[0].low == bRange[0].hi { // found beacon
			x := bRange[0].low
			freq = x*4000000 + y
			break
		}
	}
	return freq
}

func IntersectRanges(bRange []Range, left, right Range) []Range {
	res := make([]Range, 0)
	for _, r := range bRange {
		i := Intersect(r, left)
		if i != nil {
			res = append(res, *i)
		}
	}
	for _, r := range bRange {
		i := Intersect(r, right)
		if i != nil {
			res = append(res, *i)
		}
	}
	return res
}

func Intersect(r1, r2 Range) *Range {
	low := Max(r1.low, r2.low)
	hi := Min(r1.hi, r2.hi)
	if hi < low {
		return nil
	}
	return &Range{low, hi}
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day15!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	sensors, beacons := LoadInput(input1)
	min, max := RangeXY(sensors, beacons)

	// count non-beacon positions
	Y := 2000000 // 10 for TEST, 2000000 actual input
	tot := SearchInRow(sensors, beacons, Y, min, max)

	fmt.Print("puzzle1 solution: [", tot, "]")
	Assert(tot == 5073496, "  <- WRONG!")

	X, Y := 4000000, 4000000 // 20,20 for TEST, 4000000,4000000 for actual input
	freq := SearchBeacon(sensors, beacons, Pos{0, 0}, Pos{X, Y})
	fmt.Print("puzzle2 solution: [", freq, "]")
	Assert(freq == 13081194638237, "  <- WRONG!")
}
