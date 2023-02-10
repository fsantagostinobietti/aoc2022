package utils

import (
	"fmt"
	"hash/fnv"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
	"gonum.org/v1/gonum/stat/combin"
)

func SortAscending[T constraints.Ordered](vv []T) {
	sort.Slice(vv, func(i, j int) bool {
		return vv[i] < vv[j]
	})
}

func SortDescending[T constraints.Ordered](vv []T) {
	sort.Slice(vv, func(i, j int) bool {
		return vv[i] > vv[j]
	})
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

func Sqrt[T constraints.Integer | constraints.Float](v T) T {
	return T(math.Sqrt(float64(v)))
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

func Hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
func BitCount(i uint32) int {
	i = i - ((i >> 1) & 0x55555555)                // add pairs of bits
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333) // quads
	i = (i + (i >> 4)) & 0x0F0F0F0F                // groups of 8
	return int((i * 0x01010101) >> 24)             // horizontal sum of bytes
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

func Factorial(n int) int {
	if n < 0 {
		panic("Requires non negative integer!")
	}
	if n == 0 {
		return 1
	}
	f := 1
	for i := 1; i <= n; i += 1 {
		f *= i
	}
	return f
}

func RemoveInSlice(s []int, i int) []int {
	dst := make([]int, len(s)-1)
	copy(dst, s[:i])
	copy(dst[i:], s[i+1:])
	return dst
}

type Permutation[T any] struct {
	LIST []T
	gen  *combin.PermutationGenerator
}

func NewPermutation[T any](ll []T) *Permutation[T] {
	perm := Permutation[T]{LIST: ll, gen: combin.NewPermutationGenerator(len(ll), len(ll))}
	return &perm
}
func (perm *Permutation[T]) Permutation() []T {
	idxs := perm.gen.Permutation(nil)
	ll := make([]T, len(idxs))
	for i, idx := range idxs {
		ll[i] = perm.LIST[idx]
	}
	return ll
}
func (perm *Permutation[T]) HasNext() bool {
	return perm.gen.Next()
}

type Product[T any] struct {
	LISTS [][]T
	gen   *combin.CartesianGenerator
}

func NewProduct(ll ...any) *Product[any] {
	LL := make([][]any, 0)
	for _, lst := range ll {
		// sort of cast of any into []any
		v := reflect.ValueOf(lst)
		rr := make([]any, v.Len())
		for i := 0; i < v.Len(); i++ {
			rr[i] = v.Index(i).Interface()
		}
		LL = append(LL, rr)
	}
	lens := []int{}
	for _, lst := range LL {
		lens = append(lens, len(lst))
	}
	prod := Product[interface{}]{LISTS: LL, gen: combin.NewCartesianGenerator(lens)}
	return &prod
}
func (prod *Product[T]) Product() []T {
	idxs := prod.gen.Product(nil)
	ll := make([]T, len(idxs))
	for i, idx := range idxs {
		ll[i] = prod.LISTS[i][idx]
	}
	return ll
}
func (prod *Product[T]) HasNext() bool {
	return prod.gen.Next()
}

type Combination[T interface{}] struct {
	LIST []T
	gen  *combin.CombinationGenerator
}

func NewCombination[T interface{}](r int, ll []T) *Combination[T] {
	comb := Combination[T]{LIST: ll, gen: combin.NewCombinationGenerator(len(ll), r)}
	return &comb
}
func (comb *Combination[T]) Combination() []T {
	idxs := comb.gen.Combination(nil)
	ll := make([]T, len(idxs))
	for i, idx := range idxs {
		ll[i] = comb.LIST[idx]
	}
	return ll
}
func (comb *Combination[T]) HasNext() bool {
	return comb.gen.Next()
}
