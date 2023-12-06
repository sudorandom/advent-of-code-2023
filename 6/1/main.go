package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	times := parseInts(lines[0])
	distances := parseInts(lines[1])

	total := 1
	for i, raceTime := range times {
		distanceGoal := distances[i]
		winConditions := 0
		for i := 1; i < raceTime; i++ {
			distance := raceTime - i
			speed := i
			if speed*distance > distanceGoal {
				winConditions++
			}
		}
		total *= winConditions
	}

	fmt.Println(total)
}

func parseInts(s string) []int {
	strs := strings.Fields(strings.Split(s, ":")[1])
	ints := make([]int, len(strs))
	for i, str := range strs {
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		ints[i] = int(v)
	}
	return ints
}
