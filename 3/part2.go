package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./3/input2")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
	}

	inGears := map[Location]struct{}{}
	var numberStr string
	assemblyCounts := map[Location]int{}
	assemblyRatios := map[Location]int64{}
	renderPart := func() {
		defer func() {
			numberStr = ""
			inGears = map[Location]struct{}{}
		}()
		if len(inGears) == 0 {
			return
		}
		num, err := strconv.ParseInt(numberStr, 10, 64)
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		for gear := range inGears {
			assemblyCounts[gear] += 1
			if _, ok := assemblyRatios[gear]; ok {
				assemblyRatios[gear] *= num
			} else {
				assemblyRatios[gear] = num
			}
		}
	}
	matrix := strings.Split(string(content), "\n")
	for i, line := range matrix {
		for j, c := range line {
			if isNumber(c) {
				numberStr = numberStr + string(c)
				assembly := gearAssembly(matrix, i, j)
				if assembly != nil {
					inGears[*assembly] = struct{}{}
				}
			} else {
				renderPart()
			}
		}
	}
	renderPart()

	var total int64
	for assembly, count := range assemblyCounts {
		if count != 2 {
			continue
		}
		total += assemblyRatios[assembly]
	}

	fmt.Println("total:", total)
}

func isGearIndicator(c rune) bool {
	return c == '*'
}

func isPeriod(c rune) bool {
	return c == '.'
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

type Location struct {
	x, y int
}

func gearAssembly(matrix []string, i, j int) *Location {
	c := matrix[i][j]
	if !isNumber(rune(c)) {
		return nil
	}

	for _, coords := range [][2]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	} {
		other, ok := get(matrix, coords[0], coords[1])
		if !ok {
			continue
		}
		if !isGearIndicator(other) {
			continue
		}
		return &Location{x: coords[1], y: coords[0]}
	}

	return nil
}

func get(matrix []string, i, j int) (rune, bool) {
	if i < 0 || len(matrix)-1 < i {
		return ' ', false
	}
	line := matrix[i]
	if j < 0 || len(line)-1 < j {
		return ' ', false
	}

	return rune(line[j]), true
}
