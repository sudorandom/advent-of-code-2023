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
	fmt.Println(directions)
	fmt.Println(network)
	fmt.Println("answer", traverseMap(network, directions))
}

func traverseMap(network map[string]Node, directions string) int {
	var steps int
	current := network["AAA"]
	for {
		for _, direction := range directions {
			var next string
			switch direction {
			case 'L':
				next = current.Left
			case 'R':
				next = current.Right
			}
			steps++
			if next == "ZZZ" {
				return steps
			}
			current = network[next]
		}
	}
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
