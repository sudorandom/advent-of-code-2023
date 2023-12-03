package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./3/1/input")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
	}

	var total int64
	var inPart bool
	var numberStr string
	renderPart := func() {
		defer func() {
			numberStr = ""
		}()
		if !inPart {
			return
		}
		inPart = false
		num, err := strconv.ParseInt(numberStr, 10, 64)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		fmt.Println("number", num)
		total += num
	}
	matrix := strings.Split(string(content), "\n")
	for i, line := range matrix {
		for j, c := range line {
			if isNumber(c) {
				numberStr = numberStr + string(c)
				if isTouchingSymbol(matrix, i, j) {
					inPart = true
				}
			} else {
				renderPart()
			}
		}
	}
	renderPart()

	fmt.Println("total:", total)
}

func isSymbol(c rune) bool {
	return !(isPeriod(c) || isNumber(c))
}

func isPeriod(c rune) bool {
	return c == '.'
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func isTouchingSymbol(matrix []string, i, j int) bool {
	c := matrix[i][j]
	if !isNumber(rune(c)) {
		return false
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
		if !isSymbol(other) {
			continue
		}
		return true
	}

	return false
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
