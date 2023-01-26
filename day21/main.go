package main

import (
	. "aoc2022/utils"
	_ "embed"
	"fmt"
	"strings"
	//mapset "github.com/deckarep/golang-set/v2"
)

type Node struct {
	n int // > 0 leaf node; = 0 operation node; -1 unknown node
	// otherwise
	operation string // char as '+', '-', '*', '/'
	op1, op2  *Node
}

//go:embed input1.txt
var input1 string

func main() {
	fmt.Println("AoC 2022, day21!")

	// remove useless newline
	input1 = strings.TrimSuffix(input1, "\n")
	root, nodeMap := LoadTree(input1)
	res1 := ComputeResult(root)
	fmt.Print("puzzle1 solution: [", res1, "]")
	Assert(res1 == 168502451381566, "  <- WRONG!")

	humn := nodeMap["humn"]
	humn.n = -1 // set 'humn' as unknown value
	v1 := ComputeResult(root.op1)
	v2 := ComputeResult(root.op2)
	var tot int = v1
	var node *Node = root.op2
	if v1 < 0 { // unknown value in op1 branch
		node = root.op1
		tot = v2
	}
	res2 := GetUnknownValue(node, tot)
	fmt.Print("puzzle2 solution: [", res2, "]")
	Assert(res2 == 3343167719435, "  <- WRONG!")
}

func GetUnknownValue(node *Node, tot int) int {
	for node.n != -1 {
		if node.n == 0 { // operation node
			v1 := ComputeResult(node.op1)
			v2 := ComputeResult(node.op2)
			if v2 == -1 { // unknown value in op2 branch
				switch node.operation {
				case "+":
					tot = tot - v1
				case "-":
					tot = v1 - tot
				case "*":
					tot = tot / v1
				case "/":
					tot = v1 / tot
				}
				node = node.op2 // follow op2 branch
			} else { // unknown value in op1 branch
				switch node.operation {
				case "+":
					tot = tot - v2
				case "-":
					tot = tot + v2
				case "*":
					tot = tot / v2
				case "/":
					tot = tot * v2
				}
				node = node.op1 // follow op1 branch
			}
		}
	}
	return tot
}

func ComputeResult(root *Node) int {
	if root.n != 0 {
		return root.n
	}
	n1 := ComputeResult(root.op1)
	if n1 < 0 {
		return -1
	}
	n2 := ComputeResult(root.op2)
	if n2 < 0 {
		return -1
	}
	switch root.operation {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	}
	return 0 // never here
}

func LoadTree(input1 string) (*Node, map[string]*Node) {
	nodeMap := make(map[string]*Node)

	for _, line := range Split(input1, "\n") {
		monkey, action := SplitPair(line, ": ")
		// get node reference if already present
		node := GetNode(nodeMap, monkey)

		toks := Split(action, " ")
		if len(toks) == 1 { // leaf node
			node.n = Atoi(action)
		} else {
			node.op1 = GetNode(nodeMap, toks[0])
			node.operation = toks[1]
			node.op2 = GetNode(nodeMap, toks[2])
		}
	}
	return nodeMap["root"], nodeMap
}

func GetNode(nodeMap map[string]*Node, key string) *Node {
	node, ok := nodeMap[key]
	if !ok {
		// add new empty node
		node = &Node{}
		nodeMap[key] = node
	}
	return node
}
