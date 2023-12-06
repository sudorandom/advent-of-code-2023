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
	time := parseInt(lines[0])
	distanceGoal := parseInt(lines[1])

	winConditions := 0
	for i := 1; i < time; i++ {
		distance := time - i
		speed := i
		if speed*distance > distanceGoal {
			winConditions++
		}
	}

	fmt.Println(winConditions)
}

func parseInt(s string) int {
	str := strings.ReplaceAll(strings.Split(s, ":")[1], " ", "")
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	return int(v)
}
