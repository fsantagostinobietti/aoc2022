package main

import (
	_ "embed"
	"fmt"
	"regexp"
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

func SumFileSizes(s string) int {
	sum := 0
	for _, file := range strings.Split(s, "\n") {
		if file != "" && !strings.HasPrefix(file, "dir ") {
			size, _ := SplitPair(file, " ")
			sum += Atoi(size)
		}
	}
	return sum
}

var startFolderExp = regexp.MustCompile(`\$ cd ([\w/]+)\n\$ ls\n`)
var endFolderExp = regexp.MustCompile(`\$ cd ([\w\.]+)\n`)

/*
	 leaf deirectory example:
			$ cd blrnnv
			$ ls
			169869 mjjj.wnq
			$ cd ..
*/
func CollapseLeafDir(in string) (string, int) {
	for _, m0 := range startFolderExp.FindAllStringSubmatchIndex(in, -1) { // find start of a directory
		head := in[:m0[0]]
		dirName := in[m0[2]:m0[3]]
		in1 := in[m0[1]:]
		m1 := endFolderExp.FindStringSubmatchIndex(in1) // find '$ cd xxx'
		files, tail := "", ""
		if m1 == nil {
			files, tail = in1, ""
		} else {
			dir := in1[m1[2]:m1[3]]
			if dir == ".." {
				files, tail = in1[:m1[0]], in1[m1[1]:]
			}
		}
		if files != "" {
			// found leaf directory
			sum := SumFileSizes(files)
			collapsedFile := Itoa(sum) + " " + dirName + "\n"
			return head + collapsedFile + tail, sum
		}
	}
	return "", 0
}

//go:embed input1.txt
var input string

func main() {
	fmt.Println("AoC 2022, day7!")

	var sizes []int // dir sizes
	for input != "" {
		var leafDirSize int
		input, leafDirSize = CollapseLeafDir(input)
		sizes = append(sizes, leafDirSize)
	}
	SortAscending(sizes)

	sum1 := 0
	for i := 0; i < len(sizes); i += 1 {
		if sizes[i] > 100000 {
			break
		}
		sum1 += sizes[i]
	}
	fmt.Print("puzzle1 solution: [", sum1, "]")
	Assert(sum1 == 1297159, "  <- WRONG!")

	totUsed := sizes[len(sizes)-1]
	minToDelete := 30000000 - (70000000 - totUsed)
	sz2 := 0
	for i := 0; i < len(sizes); i += 1 {
		if sizes[i] >= minToDelete {
			sz2 = sizes[i]
			break
		}
	}
	fmt.Print("puzzle2 solution: [", sz2, "]")
	Assert(sz2 == 3866390, "  <- WRONG!")
}
