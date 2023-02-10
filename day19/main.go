package main

import (
	. "aoc2022/utils"
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

// indexes and size definitions
const (
	ORE      int = 0
	CLAY     int = 1
	OBSIDIAN int = 2
	GEODE    int = 3

	N_BOTS      int = 4
	N_RESOURCES int = 3
)

func CollectResources(botResources [N_RESOURCES]int, nResources [N_RESOURCES]int, nBots [N_BOTS]int) (int, [N_RESOURCES]int) {
	for i := 0; i < N_RESOURCES; i += 1 {
		if botResources[i] > 0 && nBots[i] == 0 { // no collecting bots available
			return math.MaxInt, nResources
		}
	}
	mm := 0
	for nResources[ORE] < botResources[ORE] || nResources[CLAY] < botResources[CLAY] || nResources[OBSIDIAN] < botResources[OBSIDIAN] {
		nResources[ORE] += nBots[ORE]
		nResources[CLAY] += nBots[CLAY]
		nResources[OBSIDIAN] += nBots[OBSIDIAN]
		mm += 1
	}
	// collected during bot building
	nResources[ORE] += nBots[ORE]
	nResources[CLAY] += nBots[CLAY]
	nResources[OBSIDIAN] += nBots[OBSIDIAN]
	mm += 1
	return mm, nResources
}

func UpperboundMaxGeodes(baseGeodes int, nBots [4]int, mm int) int {
	upper := baseGeodes
	for mm > 0 {
		upper += nBots[GEODE]
		// build new geode bot
		nBots[GEODE] += 1
		mm -= 1
	}
	return upper
}

// recursive function
func _MaxGeodes(blueprint [N_BOTS][N_RESOURCES]int, nBots [N_BOTS]int, nResources [N_RESOURCES]int, mm int,
	baseGeodes int, maxGeodes int) int {

	for botIdx := 0; botIdx < N_BOTS; botIdx += 1 {

		if nBots[botIdx] >= maxBots[botIdx] { // enough bots of this type
			maxGeodes = Max(maxGeodes, baseGeodes+nBots[GEODE]*mm)
			continue
		}

		_nBots := nBots
		_mm := mm
		// minutes needed to collect resources to build bot
		minutes, _nResources := CollectResources(blueprint[botIdx], nResources, _nBots)
		if minutes > _mm { // no time to build bot
			maxGeodes = Max(maxGeodes, baseGeodes+_nBots[GEODE]*_mm)
			continue
		}

		// build bot
		// update cracked geodes and remaing time
		geodes := _nBots[3] * minutes // geodes cracked during resource collection and build
		_mm -= minutes
		//consume resources
		for i := 0; i < N_RESOURCES; i += 1 {
			_nResources[i] -= blueprint[botIdx][i]
		}
		// build bot
		_nBots[botIdx] += 1
		if UpperboundMaxGeodes(baseGeodes+geodes, _nBots, _mm) > maxGeodes { // possible increase in maxGeodes
			// deep search
			maxGeodes = Max(maxGeodes, _MaxGeodes(blueprint, _nBots, _nResources, _mm, baseGeodes+geodes, maxGeodes))
		}
	}
	return maxGeodes
}

func MaxGeodes(blueprint [N_BOTS][N_RESOURCES]int, minutes int) int {
	maxBots = MaxBots(blueprint) // global var
	// init state
	nBots := [N_BOTS]int{1, 0, 0, 0}
	nResources := [N_RESOURCES]int{0, 0, 0}
	baseGeodes := 0
	maxGeodes := 0
	return _MaxGeodes(blueprint, nBots, nResources, minutes, baseGeodes, maxGeodes)
}

func MaxBots(blueprint [N_BOTS][N_RESOURCES]int) [N_BOTS]int {
	mBots := [N_BOTS]int{0, 0, 0, math.MaxInt}
	for r := 0; r < N_RESOURCES; r += 1 {
		for b := 0; b < N_BOTS; b += 1 {
			if blueprint[b][r] > mBots[r] {
				mBots[r] = blueprint[b][r]
			}
		}
	}
	return mBots
}

func LoadBlueprints(input string) [][N_BOTS][N_RESOURCES]int {
	bps := make([][N_BOTS][N_RESOURCES]int, 0)
	var moveExp = regexp.MustCompile(`Blueprint \d+: Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	for _, line := range strings.Split(input, "\n") {
		match := moveExp.FindStringSubmatch(line)
		var bp [N_BOTS][N_RESOURCES]int
		bp[ORE][ORE] = Atoi(match[1])
		bp[CLAY][ORE] = Atoi(match[2])
		bp[OBSIDIAN][ORE], bp[OBSIDIAN][CLAY] = Atoi(match[3]), Atoi(match[4])
		bp[GEODE][ORE], bp[GEODE][OBSIDIAN] = Atoi(match[5]), Atoi(match[6])

		bps = append(bps, bp)
	}
	return bps
}

var maxBots [N_BOTS]int

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day19!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	blueprints := LoadBlueprints(input1)

	totQuality := 0
	for i, bp := range blueprints {
		totQuality += (i + 1) * MaxGeodes(bp, 24)
	}
	fmt.Print("puzzle1 solution: [", totQuality, "]")
	Assert(totQuality == 1365, "  <- WRONG!")

	res2 := 1
	for i := 0; i < 3; i += 1 {
		res2 *= MaxGeodes(blueprints[i], 32)
	}
	fmt.Print("puzzle2 solution: [", res2, "]")
	Assert(res2 == 4864, "  <- WRONG!")
}
