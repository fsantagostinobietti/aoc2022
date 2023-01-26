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

func InitWall(input1 string) [][]bool {
	wall := make([][]bool, 0) // initially empty
	for _, line := range strings.Split(input1, "\n") {
		points := strings.Split(line, " -> ")
		for i := 0; i < len(points)-1; i += 1 {
			startX, startY := SplitPair(points[i], ",")
			endX, endY := SplitPair(points[i+1], ",")
			wall = DrawFromTo(wall, Atoi(startX), Atoi(startY), Atoi(endX), Atoi(endY))
		}
	}
	return wall
}

func DrawFromTo(wall [][]bool, startX, startY, endX, endY int) [][]bool {
	deltaY := Max(startY, endY) - len(wall)
	if deltaY > 0 {
		for ; deltaY >= 0; deltaY -= 1 {
			wall = append(wall, make([]bool, 1000)) // big enough
		}
	}
	if startX == endX {
		for y := Min(startY, endY); y <= Max(startY, endY)-1; y += 1 {
			wall[y][startX] = true
		}
	}
	if startY == endY {
		for x := Min(startX, endX); x <= Max(startX, endX); x += 1 {
			wall[startY][x] = true
		}
	}
	return wall
}

func PrintWall(wall [][]bool) {
	for _, line := range wall {
		for x, v := range line {
			if x >= 500-12 && x <= 500+12 {
				if v {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("")
	}
}

func MoveSandOneStep(wall [][]bool, x, y int) (int, int) {
	if wall[y+1][x] == false {
		return x, y + 1
	}
	if wall[y+1][x-1] == false {
		return x - 1, y + 1
	}
	if wall[y+1][x+1] == false {
		return x + 1, y + 1
	}
	return x, y
}

func RunSand(wall [][]bool) int {
	sandUnitsAtRest := 0
	again := true
	for again {
		// new sand unit
		sandX, sandY := 500, 0
		for sandY < len(wall)-1 {
			//fmt.Println(sandX, sandY)
			x, y := MoveSandOneStep(wall, sandX, sandY)
			if y == sandY { // sand stopped
				wall[y][x] = true
				sandUnitsAtRest += 1
				if x == 500 && y == 0 { // stop sand
					return sandUnitsAtRest
				}
				break
			}
			sandX, sandY = x, y
		}
		again = sandY < len(wall)-1 // true when sand is freefalling
	}
	return sandUnitsAtRest
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day13!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	// init wall from input text
	wall := InitWall(input1)
	//PrintWall(wall)

	sandUnitsAtRest := RunSand(wall)
	//PrintWall(wall)

	fmt.Print("puzzle1 solution: [", sandUnitsAtRest, "]")
	Assert(sandUnitsAtRest == 618, "  <- WRONG!")

	wall = InitWall(input1)
	// add floor
	wall = append(wall, make([]bool, len(wall[0])))
	wall = append(wall, make([]bool, len(wall[0])))
	for i := range wall[len(wall)-1] {
		wall[len(wall)-1][i] = true
	}
	//PrintWall(wall)
	sandUnitsAtRest = RunSand(wall)
	//PrintWall(wall)

	fmt.Print("puzzle2 solution: [", sandUnitsAtRest, "]")
	Assert(sandUnitsAtRest == 26358, "  <- WRONG!")
}
