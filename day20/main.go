package main

import (
	. "aoc2022/utils"
	_ "embed"
	"fmt"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

type Node struct {
	value int
	prev  int // index pos
	next  int //
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day20!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	MUL := 1
	encrypted, zero := LoadEncrypted(input1, MUL)
	TIMES := 1
	plain := MixArray(encrypted, TIMES)
	//fmt.Println(plain)
	res1 := GetNodeAfter(plain, zero, 1000%len(plain)).value + GetNodeAfter(plain, zero, 2000%len(plain)).value + GetNodeAfter(plain, zero, 3000%len(plain)).value
	fmt.Print("puzzle1 solution: [", res1, "]")
	Assert(res1 == 23321, "  <- WRONG!")

	MUL = 811589153
	encrypted, zero = LoadEncrypted(input1, MUL)
	TIMES = 10
	plain = MixArray(encrypted, TIMES)
	res2 := GetNodeAfter(plain, zero, 1000%len(plain)).value + GetNodeAfter(plain, zero, 2000%len(plain)).value + GetNodeAfter(plain, zero, 3000%len(plain)).value
	fmt.Print("puzzle2 solution: [", res2, "]")
	Assert(res2 == 1428396909280, "  <- WRONG!")
}

func GetNodeAfter(slice []Node, idx int, N int) Node {
	i := idx
	for n := 0; n < N; n += 1 {
		i = slice[i].next
	}
	return slice[i]
}

func MixArray(slice []Node, iters int) []Node {
	for it := 0; it < iters; it += 1 {
		for idx, v := range slice {
			if v.value != 0 {
				Move(slice, idx, v.value)
			}
		}
	}
	return slice
}

func Move(slice []Node, idx, N int) {
	// remove node
	prev, next := slice[idx].prev, slice[idx].next
	slice[prev].next = next
	slice[next].prev = prev

	// move node curosor
	i := slice[idx].next
	sgn := Sign(N)
	N = Abs(N) % (len(slice) - 1) // reduce cycles
	for n := 0; n < N; n += 1 {
		if sgn > 0 {
			i = slice[i].next
		} else { // sgn < 0
			i = slice[i].prev
		}
	}

	// insert node to the left of 'i'
	InsertNode(slice, idx, i)
}

func InsertNode(slice []Node, idx int, i int) {
	prev := slice[i].prev
	slice[idx].next = i
	slice[prev].next = idx
	slice[i].prev = idx
	slice[idx].prev = prev
}

func LoadEncrypted(input1 string, MUL int) ([]Node, int) {
	s := make([]Node, 0)
	var zero int
	for i, v := range Split(input1, "\n") {
		s = append(s, Node{value: MUL * Atoi(v), prev: 0, next: 0})
		if s[i].value == 0 {
			zero = i
		}
		if i > 0 {
			s[i].prev = i - 1
			s[i-1].next = i
		}
	}
	s[0].prev = len(s) - 1
	s[len(s)-1].next = 0
	return s, zero
}
