package main

import (
	. "aoc2022/utils"
	_ "embed"
	"fmt"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

type Rock struct {
	data []uint8
}

type State struct {
	rockIdx   int
	jetIdx    int
	topLayers []uint8

	height int // not part of state
}

var rocks = []Rock{
	{[]uint8{0b00111100}}, // rock 0
	{[]uint8{0b00010000, 0b00111000, 0b00010000}},
	{[]uint8{0b00111000, 0b00001000, 0b00001000}},
	{[]uint8{0b00100000, 0b00100000, 0b00100000, 0b00100000}},
	{[]uint8{0b00110000, 0b00110000}},
}

func Shift(rock Rock, jet byte) Rock {
	var mask uint8 = 0b11111110 // chamber width

	var shifted = Rock{make([]uint8, len(rock.data))}
	copy(shifted.data, rock.data) // copy rock

	for i := 0; i < len(rock.data); i += 1 {
		if jet == '>' {
			shifted.data[i] >>= 1
			shifted.data[i] &= mask
		} else { // jet == '<'
			shifted.data[i] <<= 1
			shifted.data[i] &= mask
		}
	}
	return shifted
}

func CountBits(v uint8) int {
	return int(PopCount(uint32(v)))
}

func JetStreamAction(chamber []uint8, jet byte, rock Rock, rockH int) Rock {
	shiftedRock := Shift(rock, jet)
	// test if rock hits walls or other rocks
	hit := false
	for i := 0; i < len(rock.data); i += 1 {
		if CountBits(chamber[rockH+i])+CountBits(rock.data[i]) != CountBits(chamber[rockH+i]|shiftedRock.data[i]) {
			hit = true
			break
		}
	}
	if hit {
		return rock // no shift
	}
	return shiftedRock
}

func TopHeight(chamber []uint8) int {
	top := 0
	for ; top < len(chamber) && chamber[top] != 0x0; top += 1 {
	}
	return top
}

func TopLayers(chamber []uint8) []uint8 {
	tops := make([]int, 7)
	for i := 0; chamber[i] != 0x0; i += 1 {
		var m uint8 = 0b10000000
		for col := 0; col < 7; col += 1 {
			if chamber[i]&(m>>col) != 0x0 {
				tops[col] = i
			}
		}
	}
	min := Min(tops...)
	max := Max(tops...)
	layers := make([]uint8, max-min+1)
	copy(layers, chamber[min:])
	return layers
}

func StepDown(chamber []uint8, rock Rock, rockH int) (int, bool) {
	belowLine := rockH - 1
	if belowLine < 0 {
		return rockH, false // reached floor
	}
	for i := 0; i < len(rock.data); i += 1 {
		if chamber[belowLine+i]&rock.data[i] != 0 {
			return rockH, false // hit a rock
		}
	}
	return belowLine, true
}

func Merge(chamber []uint8, rock Rock, rockH int) []uint8 {
	for i := 0; i < len(rock.data); i += 1 {
		chamber[rockH+i] |= rock.data[i]
	}
	return chamber
}

func EqualsState(s1, s2 State) bool {
	if s1.rockIdx != s2.rockIdx || s1.jetIdx != s2.jetIdx || len(s1.topLayers) != len(s2.topLayers) {
		return false
	}
	for i := 0; i < len(s1.topLayers); i += 1 {
		if s1.topLayers[i] != s2.topLayers[i] {
			return false
		}
	}
	return true
}

func LookForCycle(state []State, actual State, n int) (bool, int, int) {
	foundCycle := false
	var cycleLen, cycleHeight int
	for i := 0; i < n; i += 1 {
		if EqualsState(actual, state[i]) {
			foundCycle = true
			cycleLen = n - i                              // number of rocks
			cycleHeight = actual.height - state[i].height // number of layers
			//fmt.Println("cycle found at:", n, "cycle len:", cycleLen)
		}
	}
	return foundCycle, cycleLen, cycleHeight
}

func SimulateFallingRocks(chamber []uint8, states []State, N int) int {
	j := 0 // jet stream index
	foundCycle := false
	var numCycles, cycleHeight int
	for n := 0; n < N; n += 1 { // rocks counter

		if !foundCycle {
			// compute state
			actual := State{
				rockIdx: n % 5, jetIdx: j, topLayers: TopLayers(chamber), height: TopHeight(chamber),
			}
			// update state history
			states[n] = actual

			// look for cycle
			var cycleLen int
			foundCycle, cycleLen, cycleHeight = LookForCycle(states, actual, n)
			if foundCycle {
				// update 'N' removing cycles
				numCycles = (N - n) / cycleLen
				N = n + (N-n)%cycleLen
				//fmt.Println("numCycles:", numCycles, "new 'N':", N)
			}
		}

		// new rock
		rock := rocks[n%5]
		rockH := TopHeight(chamber) + 3

		falling := true
		for falling { // simulate falling of single rock
			// shift by jet stream
			rock = JetStreamAction(chamber, jetStream[j], rock, rockH)
			j = (j + 1) % len(jetStream)

			// fall 1 step down
			rockH, falling = StepDown(chamber, rock, rockH)
		}
		chamber = Merge(chamber, rock, rockH)
	}
	return TopHeight(chamber) + numCycles*cycleHeight
}

func PrintChamber(chamber []uint8) {
	top := TopHeight(chamber)
	for i := top - 1; i >= 0; i -= 1 {
		fmt.Println(fmt.Sprintf("%08b", chamber[i])[:7])
	}
}

// global vars
var jetStream string

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day17!")

	// remove useless newline
	jetStream = strings.TrimSuffix(input1, "\n")

	chamber := make([]uint8, 10000) // empty chamber
	state := make([]State, 5000)    // empty chamber state

	top1 := SimulateFallingRocks(chamber, state, 2022)
	//PrintChamber(chamber)
	fmt.Print("puzzle1 solution: [", top1, "]")
	Assert(top1 == 3153, "  <- WRONG!")

	// resest chamber and its state
	chamber = make([]uint8, 10000)
	state = make([]State, 5000)

	top2 := SimulateFallingRocks(chamber, state, 1000000000000)
	fmt.Print("puzzle2 solution: [", top2, "]")
	Assert(top2 == 1553665689155, "  <- WRONG!")
}
