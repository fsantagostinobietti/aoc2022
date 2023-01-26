package main

import (
	. "aoc2022/utils"
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
	//mapset "github.com/deckarep/golang-set/v2"
)

const (
	N = 0 // North
	E = 1 // East
	S = 2 // South
	W = 3 // West
)

type Bot struct {
	r, c   int
	facing int // N="^", E=">", S="v", N="W"
}

type Face struct {
	next [4]*Face // neighbours in this order: N, E, S, W
	r, c int      // upper left corner position in board
}

//go:embed input1.txt
var input1 string

// global
var board []string
var faces []*Face // used in part2
var faceSz int    //

func main() {
	fmt.Println("AoC 2022, day22!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")

	var actions []string
	board, actions = LoadFile(input1)

	bot1 := InitBot(board)
	bot1 = bot1.Run(actions, WrapAround1)
	res1 := 1000*bot1.r + 4*bot1.c + slices.Index([]int{E, S, W, N}, bot1.facing)
	fmt.Print("puzzle1 solution: [", res1, "]")
	Assert(res1 == 1484, "  <- WRONG!")

	faceSz = FaceSize(board)
	faces = FoldBoardOnCube(board)
	bot2 := InitBot(board)
	bot2 = bot2.Run(actions, WrapAround2)
	res2 := 1000*bot2.r + 4*bot2.c + slices.Index([]int{E, S, W, N}, bot2.facing)
	fmt.Print("puzzle2 solution: [", res2, "]")
	Assert(res2 == 142228, "  <- WRONG!")
}

func CurrentFace(bot Bot) *Face {
	r := 1 + ((bot.r-1)/faceSz)*faceSz // row/col indexes start from 1
	c := 1 + ((bot.c-1)/faceSz)*faceSz //
	for _, f := range faces {
		if f.r == r && f.c == c {
			return f
		}
	}
	return nil // never here
}

func WrapAround2(bot Bot) Bot {
	var bot1 Bot
	face := CurrentFace(bot)

	// compute new bot position
	dr, dc := bot.r-face.r, bot.c-face.c
	dir := bot.facing
	nextFace := face.next[dir]
	dir1 := ToDirection(nextFace, face)
	mapping := [][]int{ // dir, dir1, flag swap dr1<->dc1, flag complement dr1, flag complement dc1
		{N, S, 0, 1, 0}, {S, N, 0, 1, 0}, {E, W, 0, 0, 1}, {W, E, 0, 0, 1}, // dir/dir1 opposite
		{N, N, 0, 0, 1}, {S, S, 0, 0, 1}, {E, E, 0, 1, 0}, {W, W, 0, 1, 0}, // dir/dir1 same
		{E, N, 1, 1, 1}, {N, E, 1, 1, 1}, {W, N, 1, 0, 0}, {N, W, 1, 0, 0}, // dir or dir1 is North
		{W, S, 1, 1, 1}, {S, W, 1, 1, 1}, {E, S, 1, 0, 0}, {S, E, 1, 0, 0}, // dir or dir1 is South
	}
	dr1, dc1 := dr, dc
	for _, m := range mapping {
		if m[0] == dir && m[1] == dir1 {
			if m[2] == 1 { // swap dr1/dc1
				dr1, dc1 = dc1, dr1
			}
			if m[3] == 1 { // complement dr1
				dr1 = (faceSz - 1) - dr1
			}
			if m[4] == 1 { // complement dc1
				dc1 = (faceSz - 1) - dc1
			}
		}
	}
	bot1.r, bot1.c = nextFace.r+dr1, nextFace.c+dc1

	// compute new bot facing
	bot1.facing = bot.facing
	dd := dir - dir1
	if dd == 0 { // dir==dir1
		bot1 = bot1.Rotate(+2)
	} else if Abs(dd) == 2 {
		// nope i.e. bot1 = RotateBot(bot1, 0)
	} else {
		bot1 = bot1.Rotate(dd)
	}

	if board[bot1.r][bot1.c] == '#' { // hit wall
		return bot // nope
	}
	return bot1
}

func (bot Bot) Rotate(dd int) Bot {
	if dd < 0 { // counter-clockwise
		dd = (dd + 4) % 4 // reverse rotation
	}
	// clockwise rotation
	bot.facing = (bot.facing + dd) % 4
	return bot
}

func FoldBoardOnCube(board []string) []*Face {
	cube := BuildInitialCube()

	bot := InitBot(board)
	visited := make([]*Face, 0, 6)
	visited = WalkUnfoldedFaces(cube, bot.r, bot.c, visited)
	return visited
}

func WalkUnfoldedFaces(face *Face, r, c int, visited []*Face) []*Face {
	face.r, face.c = r, c // init upper left corner
	visited = append(visited, face)
	for dir, delta := range [][2]int{{0, 0}, {0, faceSz}, {faceSz, 0}, {0, -faceSz}} {
		r1, c1 := r+delta[0], c+delta[1]
		if c1 >= 0 && board[r1][c1] != ' ' { // face found
			if !isVisited(visited, [2]int{r1, c1}) {
				AdjustNextCubeFace(face, dir)
				visited = WalkUnfoldedFaces(face.next[dir], r1, c1, visited)
			}
		}
	}
	return visited
}

func AdjustNextCubeFace(face *Face, direction int) {
	n := face.next[direction]
	dir := (direction + 2) % 4 // opposite to 'direction' (ex. N->S, W->E, ...)
	for ToDirection(n, face) != dir {
		// rotate face clockwise
		n.next[N], n.next[E], n.next[S], n.next[W] = n.next[E], n.next[S], n.next[W], n.next[N]
	}
}

func ToDirection(from *Face, to *Face) int {
	for dir, f := range from.next {
		if f == to {
			return dir
		}
	}
	return -1 // never here
}

func isVisited(visited []*Face, v [2]int) bool {
	for _, tst := range visited {
		if tst.r == v[0] && tst.c == v[1] {
			return true
		}
	}
	return false
}

func BuildInitialCube() *Face {
	up, down, front, back, left, right := Face{}, Face{}, Face{}, Face{}, Face{}, Face{}

	front.next = [4]*Face{&up, &right, &down, &left}
	down.next = [4]*Face{&front, &right, &back, &left}
	back.next = [4]*Face{&down, &right, &up, &left}
	up.next = [4]*Face{&back, &right, &front, &left}
	right.next = [4]*Face{&up, &back, &down, &front}
	left.next = [4]*Face{&up, &front, &down, &back}

	return &up
}

func FaceSize(board []string) int {
	tot := 0
	for _, line := range board {
		for _, v := range line {
			if v != ' ' {
				tot += 1
			}
		}
	}
	return Sqrt(tot / 6)
}

func (bot Bot) Run(actions []string, wrapAround func(Bot) Bot) Bot {
	for _, a := range actions {
		if a != "L" && a != "R" { // move
			moves := Atoi(a)
			bot = bot.Move(moves, wrapAround)
		} else { // rotate
			rot := 1 // "R"
			if a == "L" {
				rot = -1
			}
			bot = bot.Rotate(rot)
		}
	}
	return bot
}

func (bot Bot) Move(moves int, wrapAround func(Bot) Bot) Bot {
	dr, dc := DeltaMove(bot)

	for ; moves > 0; moves -= 1 {
		nextR := bot.r + dr
		nextC := bot.c + dc
		if board[nextR][nextC] == '#' { // hit wall
			break
		}
		if board[nextR][nextC] == ' ' { // wrap around
			bot = wrapAround(bot)
			dr, dc = DeltaMove(bot)
		} else {
			// update bot position
			bot.r = nextR
			bot.c = nextC
		}
	}

	return bot
}

func WrapAround1(bot Bot) Bot {
	dr, dc := DeltaMove(bot)
	dr, dc = -dr, -dc

	r := bot.r
	c := bot.c
	for board[r+dr][c+dc] != ' ' {
		r += dr
		c += dc
	}
	if board[r][c] != '#' {
		bot.r = r
		bot.c = c
	}
	return bot
}

func DeltaMove(bot Bot) (int, int) {
	DELTA := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dr := DELTA[bot.facing][0]
	dc := DELTA[bot.facing][1]
	return dr, dc
}

func InitBot(board []string) Bot {
	r, c := 1, 1
	for i, char := range Split(board[r], "") {
		if char != " " {
			c = i
			break
		}
	}
	return Bot{r, c, E}
}

func LoadFile(input string) ([]string, []string) {
	s1, s2 := SplitPair(input, "\n\n")
	board := Split(s1, "\n")
	// max lines lenght
	max := 0
	for _, line := range board {
		if len(line) > max {
			max = len(line)
		}
	}
	// surround with empty char border
	max += 2
	board = append([]string{""}, board...)
	board = append(board, "")
	for i, line := range board {
		board[i] = " " + line + fmt.Sprintf("%*s", (max-1-len(line)), "")
	}

	steps := regexp.MustCompile("[R|L]").Split(s2, -1)
	rots := regexp.MustCompile("[R|L]").FindAllString(s2, -1)
	actions := make([]string, 0, 2*len(steps))
	for i := 0; i < len(rots); i += 1 {
		actions = append(actions, steps[i])
		actions = append(actions, rots[i])
	}
	actions = append(actions, steps[len(steps)-1])

	return board, actions
}
