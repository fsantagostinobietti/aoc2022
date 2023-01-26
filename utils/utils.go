package utils

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

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

func Min[T constraints.Ordered](vn ...T) (m T) {
	m = vn[0]
	for _, v := range vn {
		if v < m {
			m = v
		}
	}
	return m
}

func Max[T constraints.Ordered](vn ...T) (m T) {
	m = vn[0]
	for _, v := range vn {
		if v > m {
			m = v
		}
	}
	return m
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

func Sqrt(v int) int {
	return int(math.Sqrt(float64(v)))
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

// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
func PopCount(i uint32) uint32 {
	i = i - ((i >> 1) & 0x55555555)                // add pairs of bits
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333) // quads
	i = (i + (i >> 4)) & 0x0F0F0F0F                // groups of 8
	return (i * 0x01010101) >> 24                  // horizontal sum of bytes
}

// Greatest common divisor
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Least common multiple
func Lcm(a, b int) int {
	return Abs(a) * Abs(b) / Gcd(a, b)
}
