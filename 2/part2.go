package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./2/input2")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
	}

	var total int64
	for _, line := range strings.Split(string(content), "\n") {
		gameID, scores, err := parseGameState(line)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		power := calculatePower(scores)
		fmt.Println(gameID, power, scores)
		total += power
	}

	fmt.Println("total:", total)
}

func calculatePower(scores [3]int64) int64 {
	return scores[0] * scores[1] * scores[2]
}

// red
// green
// blue

func parseGameState(line string) (int64, [3]int64, error) {
	var gameID int64
	maximums := [3]int64{0, 0, 0}
	parts := strings.Split(line, ":")
	gameHeaderParts := strings.Split(parts[0], " ")
	gameID, err := strconv.ParseInt(gameHeaderParts[1], 10, 64)
	if err != nil {
		return gameID, maximums, err
	}

	for _, game := range strings.Split(parts[1], ";") {
		for _, color := range strings.Split(game, ",") {
			colorParts := strings.Split(strings.TrimSpace(color), " ")
			count, err := strconv.ParseInt(colorParts[0], 10, 64)
			fmt.Println(colorParts, count)
			if err != nil {
				return gameID, maximums, err
			}
			switch colorParts[1] {
			case "red":
				maximums[0] = max(maximums[0], count)
			case "green":
				maximums[1] = max(maximums[1], count)
			case "blue":
				maximums[2] = max(maximums[2], count)
			}
		}
	}
	return gameID, maximums, nil
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
