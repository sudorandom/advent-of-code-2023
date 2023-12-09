package main

import (
	"fmt"
	"regexp"
	"strings"

	_ "embed"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")

	directions := lines[0]
	network := parseNetwork(lines[2:])

	startingPositions := []string{}
	for key := range network {
		if key[len(key)-1] == 'A' {
			startingPositions = append(startingPositions, key)
		}
	}
	fmt.Println(startingPositions)
	fmt.Println(traverseMap(network, startingPositions, directions))
}

func traverseMap(network map[string]Node, positions []string, directions string) int {
	steps := map[int]int{}
	var totalSteps int
	for {
		for _, direction := range directions {
			totalSteps++
			if len(steps) == len(positions) {
				counts := []int{}
				for _, count := range steps {
					counts = append(counts, count)
				}
				return LCM(counts...)
			}
			for i, position := range positions {
				if _, ok := steps[i]; ok {
					continue
				}
				next, isEnding := stepOne(network, position, direction)
				if isEnding {
					steps[i] = totalSteps
				}
				positions[i] = next
			}

		}
	}
}

func stepOne(network map[string]Node, current string, direction rune) (string, bool) {
	var next string
	switch direction {
	case 'L':
		next = network[current].Left
	case 'R':
		next = network[current].Right
	}
	return next, next[len(next)-1] == 'Z'
}

type Node struct {
	Left  string
	Right string
}

func parseNetwork(lines []string) map[string]Node {
	r := regexp.MustCompile(`(?P<Node>\w+) = \((?P<Left>\w+), (?P<Right>\w+)\)`)

	network := map[string]Node{}
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		network[match[1]] = Node{
			Left:  match[2],
			Right: match[3],
		}
	}
	return network
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	if len(integers) < 2 {
		panic("LCM requires two arguments")
	}
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)

	for _, integer := range integers[2:] {
		result = LCM(result, integer)
	}

	return result
}
