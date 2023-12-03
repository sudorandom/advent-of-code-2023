package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("./1/input2")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
	}

	total := 0
	for _, line := range strings.Split(string(content), "\n") {
		calibrationValue := (firstDigit(line) * 10) + lastDigit(line)
		fmt.Println(line, calibrationValue)
		total += calibrationValue
	}

	fmt.Println("total:", total)
}

func firstDigit(line string) int {
	for i, c := range line {
		if d, ok := charToDigit(c); ok {
			return d
		}

		if d, ok := digitPrefix(line[i:]); ok {
			return d
		}
	}

	return 0
}

func lastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if d, ok := charToDigit(rune(line[i])); ok {
			return d
		}

		if d, ok := digitPrefix(line[i:]); ok {
			return d
		}
	}

	return 0
}

func digitPrefix(s string) (int, bool) {
	switch {
	case strings.HasPrefix(s, "zero"):
		return 0, true
	case strings.HasPrefix(s, "one"):
		return 1, true
	case strings.HasPrefix(s, "two"):
		return 2, true
	case strings.HasPrefix(s, "three"):
		return 3, true
	case strings.HasPrefix(s, "four"):
		return 4, true
	case strings.HasPrefix(s, "five"):
		return 5, true
	case strings.HasPrefix(s, "six"):
		return 6, true
	case strings.HasPrefix(s, "seven"):
		return 7, true
	case strings.HasPrefix(s, "eight"):
		return 8, true
	case strings.HasPrefix(s, "nine"):
		return 9, true
	}
	return 0, false
}

func charToDigit(c rune) (int, bool) {
	switch c {
	case '0':
		return 0, true
	case '1':
		return 1, true
	case '2':
		return 2, true
	case '3':
		return 3, true
	case '4':
		return 4, true
	case '5':
		return 5, true
	case '6':
		return 6, true
	case '7':
		return 7, true
	case '8':
		return 8, true
	case '9':
		return 9, true
	}
	return 0, false
}
