package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./4/1/input")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
	}

	var total int64
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		winning, recved := parseGameState(line)
		winningMap := map[int64]struct{}{}
		for _, number := range winning {
			winningMap[number] = struct{}{}
		}
		var score int64
		for _, number := range recved {
			if _, ok := winningMap[number]; !ok {
				continue
			}
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
		total += score
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
