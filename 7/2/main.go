package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input
var input string

type HandClass int

func (c HandClass) String() string {
	switch c {
	case FiveOfAKind:
		return "FiveOfAKind"
	case FourOfAKind:
		return "FourOfAKind"
	case FullHouse:
		return "FullHouse"
	case ThreeOfKind:
		return "ThreeOfKind"
	case TwoPair:
		return "TwoPair"
	case OnePair:
		return "OnePair"
	case HighCard:
		return "HighCard"
	}
	return "???"
}

const (
	FiveOfAKind HandClass = iota
	FourOfAKind
	FullHouse
	ThreeOfKind
	TwoPair
	OnePair
	HighCard
)

type Hand struct {
	Cards string
	Bid   int
	Class HandClass
}

func HandFromLine(s string) Hand {
	parts := strings.Fields(s)
	hand := parts[0]

	bid, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		log.Fatalf("invalid bid: %s", err)
	}
	// Sort by class and then by card values
	return Hand{
		Cards: hand,
		Bid:   int(bid),
		Class: HandClassFromHand(hand),
	}
}

func CardToValue(r rune) int {
	switch r {
	case 'J':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	log.Fatalf("invalid card?: %s", string(r))
	return 0
}
func HandClassFromHand(s string) HandClass {
	freqs := make(map[rune]int)
	for _, char := range s {
		freqs[char] = freqs[char] + 1
	}

	var maxChar rune
	var maxCount int
	for char, count := range freqs {
		if char == 'J' {
			continue
		}
		if count > maxCount {
			maxCount = count
			maxChar = char
		}
	}

	freqs[maxChar] += freqs['J']
	delete(freqs, 'J')

	counts := make(map[int]struct{})
	for _, count := range freqs {
		counts[count] = struct{}{}
	}

	// Five of a kind
	if len(freqs) == 1 {
		return FiveOfAKind
	}

	if len(freqs) == 2 {
		// Four of a kind
		if _, ok := counts[4]; ok {
			return FourOfAKind
		}

		// Full House
		return FullHouse
	}

	// Three of a kind
	if _, ok := counts[3]; ok {
		return ThreeOfKind
	}

	// Pairs
	pairCount := 0
	for _, count := range freqs {
		if count == 2 {
			pairCount++
		}
	}
	switch pairCount {
	case 2:
		return TwoPair
	case 1:
		return OnePair
	}

	// This hand sucks
	return HighCard
}

func main() {
	lines := strings.Split(input, "\n")
	hands := make([]Hand, 0, len(lines))
	for _, line := range lines {
		hands = append(hands, HandFromLine(line))
	}

	for _, hand := range hands {
		fmt.Printf("hand: %s; class=%s\n", hand.Cards, hand.Class)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Class < hands[j].Class {
			return true
		} else if hands[i].Class > hands[j].Class {
			return false
		}

		for k := 0; k < 5; k++ {
			a := CardToValue(rune(hands[i].Cards[k]))
			b := CardToValue(rune(hands[j].Cards[k]))
			if a == b {
				continue
			}

			if a < b {
				return false
			} else {
				return true
			}
		}

		return false
	})

	total := 0
	for i, hand := range hands {
		total += (len(hands) - i) * hand.Bid
	}

	fmt.Println(total)
}
