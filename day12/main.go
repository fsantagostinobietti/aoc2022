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

type Pos struct {
	x, y int
}

func InitAlgo(input string) ([][]string, Pos, Pos) {
	grid := make([][]string, 0)
	var start, end Pos
	for y, line := range strings.Split(input, "\n") {
		row := strings.Split(line, "")
		// search for 'S' and 'E'
		for x, char := range row {
			if char == "S" {
				start = Pos{x, y}
				row[x] = "a"
			} else if char == "E" {
				end = Pos{x, y}
				row[x] = "z"
			}
		}
		grid = append(grid, row)
	}
	return grid, start, end
}

func PrintMatrix(m [][]bool) {
	for _, r := range m {
		for _, c := range r {
			if c {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println("")
	}
}

func TestNext(grid [][]string, visited [][]bool, last Pos, next Pos) bool {
	if next.x < 0 || next.y < 0 || next.y >= len(grid) || next.x >= len(grid[next.y]) {
		return false
	}
	s := grid[last.y][last.x][0]
	e := grid[next.y][next.x][0]
	delta := int(e) - int(s)
	if delta > 1 {
		return false
	}
	return !visited[next.y][next.x]
}

func ContainsEnd(list []Pos, end Pos) bool {
	for _, t := range list {
		if t.y == end.y && t.x == end.x {
			return true
		}
	}
	return false
}

var MOVES = []Pos{{0, -1} /*UP*/, {0, 1} /*DOWN*/, {-1, 0} /*LEFT*/, {1, 0} /*RIGTH*/}

func RunAlgo(grid [][]string, start, end Pos) int {
	visited := MakeVisited(len(grid), len(grid[0]))
	visited[start.y][start.y] = true
	list := []Pos{start}
	for d := 1; true; d += 1 {
		newList := make([]Pos, 0)
		for _, last := range list {
			for _, move := range MOVES {
				next := Pos{last.x + move.x, last.y + move.y}
				if TestNext(grid, visited, last, next) {
					newList = append(newList, next)
					visited[next.y][next.x] = true
				}
			}
		}
		if len(newList) == 0 {
			return -1 // 'E' not reachable
		}
		if ContainsEnd(newList, end) {
			return d
		}
		list = newList
		//fmt.Println(list)
	}
	return -2 // never here
}

func MakeVisited(rows, cols int) [][]bool {
	m := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]bool, cols)
	}
	return m
}

func Reset(visited [][]bool) {
	for r, _ := range visited {
		for c, _ := range visited[r] {
			visited[r][c] = false
		}
	}
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day12!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	grid, start, end := InitAlgo(input1)
	dist1 := RunAlgo(grid, start, end)
	//PrintMatrix(visited)
	fmt.Print("puzzle1 solution: [", dist1, "]")
	Assert(dist1 == 330, "  <- WRONG!")

	shortest := 10000 // big enough
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "a" {
				s := Pos{x, y}
				d := RunAlgo(grid, s, end)
				if d > 0 && d < shortest {
					shortest = d
				}
			}
		}
	}
	fmt.Print("puzzle2 solution: [", shortest, "]")
	Assert(shortest == 321, "  <- WRONG!")

}
