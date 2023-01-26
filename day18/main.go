package main

import (
	. "aoc2022/utils"
	_ "embed"
	"fmt"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

type Cube struct {
	x, y, z int
}

func Contains(cubes []Cube, c Cube) bool {
	for _, c1 := range cubes {
		if c1 == c {
			return true
		}
	}
	return false
}

func CoumputeSurface(droplet []Cube) int {
	surface := 0
	for _, c1 := range droplet {
		freeFaces := 6
		for _, c2 := range droplet {
			if Abs(c1.x-c2.x)+Abs(c1.y-c2.y)+Abs(c1.z-c2.z) == 1 { // adiacent cubes
				freeFaces -= 1
			}
		}
		surface += freeFaces
	}
	return surface
}

func RemoveFromSlice(s []Cube, e Cube) []Cube {
	for idx, e1 := range s {
		if e1 == e {
			s[idx] = s[len(s)-1] // copy last element
			s = s[:len(s)-1]     // truncate slice
		}
	}
	return s
}

func RecursiveRemove(cubes []Cube, c Cube) []Cube {
	if !Contains(cubes, c) {
		return cubes
	}
	// remove cube from list
	cubes = RemoveFromSlice(cubes, c)
	// recurse in all directions
	for _, delta := range [...]Cube{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}} {
		c1 := c
		c1.x += delta.x
		c1.y += delta.y
		c1.z += delta.z
		cubes = RecursiveRemove(cubes, c1)
	}
	return cubes
}

func RemoveExternalCubes(airCubes []Cube) []Cube {
	start := Cube{0, 0, 0}
	airCubes = RecursiveRemove(airCubes, start)
	return airCubes
}

//go:embed input1.txt
var input1 string

var SZ int = 23 // 7 TEST, 23  actual

func main() {
	fmt.Println("AoC 2022, day17!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	// load input
	droplet := make([]Cube, 0)
	for _, line := range Split(input1, "\n") {
		toks := Split(line, ",")
		droplet = append(droplet, Cube{x: Atoi(toks[0]), y: Atoi(toks[1]), z: Atoi(toks[2])})
	}
	//fmt.Println(droplet)

	// count free faces
	surface := CoumputeSurface(droplet)
	fmt.Print("puzzle1 solution: [", surface, "]")
	Assert(surface == 4400, "  <- WRONG!")

	// computes air cubes
	airCubes := make([]Cube, 0)
	for x := 0; x < SZ; x += 1 {
		for y := 0; y < SZ; y += 1 {
			for z := 0; z < SZ; z += 1 {
				c := Cube{x, y, z}
				if !Contains(droplet, c) {
					airCubes = append(airCubes, c)
				}
			}
		}
	}
	//fmt.Println("\n", airCubes)

	// remove all cube external to droplets
	airCubes = RemoveExternalCubes(airCubes)
	//fmt.Println("\n", airCubes)
	surface2 := surface - CoumputeSurface(airCubes)
	fmt.Print("puzzle2 solution: [", surface2, "]")
	Assert(surface2 == 2522, "  <- WRONG!")
}
