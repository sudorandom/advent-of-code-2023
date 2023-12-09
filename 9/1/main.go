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
	var total int
	for _, line := range lines {
		nums := parseIntList(line)
		lastCols := resolveTriangle(nums)
		fmt.Printf("lastCols: %+v\n", lastCols)
		var delta int
		for i := len(lastCols) - 2; i >= 0; i-- {
			lastCol := lastCols[i]
			delta = lastCol + delta
		}
		fmt.Println("row", delta)
		total += delta
	}

	fmt.Println(total)
}

func resolveTriangle(row []int) []int {
	result := []int{row[len(row)-1]}
	for {
		row = resolveTriangleRow(row)
		if len(row) == 0 {
			return result
		}
		result = append(result, row[len(row)-1])
		var needsMore bool
		for _, i := range row {
			if i != 0 {
				needsMore = true
				break
			}
		}
		if !needsMore {
			return result
		}
	}
}

func resolveTriangleRow(row []int) []int {
	next := make([]int, len(row)-1)
	for i := 0; i < len(row)-1; i++ {
		next[i] = row[i+1] - row[i]
	}
	return next
}

func parseIntList(s string) []int {
	intStrs := strings.Fields(s)
	ints := make([]int, len(intStrs))
	for i, s := range intStrs {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalf("Error: %s", s)
		}
		ints[i] = int(num)
	}
	return ints
}
