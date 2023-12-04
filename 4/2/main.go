package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./4/2/input")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
	}

	extraCards := map[int]int{}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		winning, received := parseGameState(line)
		winningMap := map[int64]struct{}{}
		for _, number := range winning {
			winningMap[number] = struct{}{}
		}
		var matches int
		for _, number := range received {
			if _, ok := winningMap[number]; !ok {
				continue
			}
			matches++
		}

		extraCount := extraCards[i] + 1
		for j := 1; j <= matches; j++ {
			extraCards[i+j] += extraCount
		}
	}
	fmt.Println(extraCards)

	total := len(lines)
	for _, count := range extraCards {
		total += count
	}

	fmt.Println("total:", total)
}

func parseGameState(line string) ([]int64, []int64) {
	gameParts := strings.Split(line, ":")
	winhave := strings.Split(gameParts[1], "|")

	winningNumbers := []int64{}
	for _, numStr := range strings.Fields(winhave[0]) {
		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			log.Fatalf("err parsing num: %s", err)
		}
		winningNumbers = append(winningNumbers, num)
	}

	receivedNumbers := []int64{}
	for _, numStr := range strings.Fields(winhave[1]) {
		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			log.Fatalf("err parsing num: %s", err)
		}
		receivedNumbers = append(receivedNumbers, num)
	}
	return winningNumbers, receivedNumbers
}
